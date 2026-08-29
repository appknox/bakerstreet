# Bridge Plan — Device-Initiated Connections (Relay Mode)

> This is a **separate, additive architecture**. The existing direct-connection mode (`bakerstreet.proto` + `Moriarty` service) remains completely untouched. Both modes coexist permanently.

---

## Problem

The current architecture requires the platform to directly reach device agents over the network. Devices run gRPC servers and the platform connects to them as a client. This forces both to be on the **same network**.

## Solution

Add a second communication path where **devices connect outbound** to a cloud relay server. The relay holds the connection open as a bidirectional stream and routes platform commands through it. No inbound ports needed on the device side.

---

## Both Modes Side by Side

```
MODE A — Direct (existing, unchanged, uses bakerstreet.proto)

┌──────────────┐    gRPC (Moriarty)   ┌──────────────┐
│   Platform   │ ───────────────────► │    Device    │
│   (client)   │ ◄─── response ────── │   (server)   │
└──────────────┘                      └──────────────┘
  Same network required. Device must be reachable.


MODE B — Relay (new, uses bakerstreet_bridge.proto)

┌──────────────┐   Dispatch API    ┌─────────────────┐    bidi stream     ┌──────────────┐
│   Platform   │ ────────────────► │  Cloud Relay    │ ◄────────────────  │    Device    │
│   (client)   │                   │  (public IP)    │ ─────────────────► │   (client)   │
│              │ ◄── response ──── │                 │                    │              │
└──────────────┘                   └─────────────────┘                    └──────────────┘
  Devices connect outbound. Works through NAT/firewalls.
```

The platform decides which mode per device based on a config flag (e.g. `connection_mode: "direct" | "relay"`).

---

## File Separation

All new definitions live in a **separate proto file** that imports shared messages from the existing one. The original `bakerstreet.proto` is **never modified**.

```
bakerstreet/
├── bakerstreet.proto                ← UNTOUCHED
├── bakerstreet_bridge.proto         ← NEW FILE (imports bakerstreet.proto)
│
├── bakerstreet.pb.go                ← Existing (unchanged)
├── bakerstreet_grpc.pb.go           ← Existing (unchanged)
├── bakerstreet_pb2.py               ← Existing (unchanged)
├── bakerstreet_pb2_grpc.py          ← Existing (unchanged)
│
├── bakerstreet_bridge.pb.go         ← NEW generated
├── bakerstreet_bridge_grpc.pb.go    ← NEW generated
├── bakerstreet_bridge_pb2.py        ← NEW generated
└── bakerstreet_bridge_pb2_grpc.py   ← NEW generated
```

Consumers of direct mode keep importing `bakerstreet` — zero impact.
Consumers of relay mode additionally import `bakerstreet_bridge`.

---

## Proto Design — `bakerstreet_bridge.proto`

```proto
syntax = "proto3";
package com.appknox.bakerstreet;

option go_package = "github.com/appknox/bakerstreet";

// Import all existing messages from the original proto.
// Message, App, Apps, Device, DeviceV2, InstallReq,
// ConfigProxyReq, AutoPilotConfig, CleanOptions, Empty
// are all available — nothing is redefined.
import "bakerstreet/bakerstreet.proto";


// ═══════════════════════════════════════════════════════════
//  COMMAND ENVELOPE — relay pushes commands to device
// ═══════════════════════════════════════════════════════════

enum CommandType {
  COMMAND_UNKNOWN          = 0;
  COMMAND_ECHO             = 1;
  COMMAND_LAUNCH_APP       = 2;
  COMMAND_CLEAR_PROXY      = 3;
  COMMAND_HEALTH_CHECK     = 4;
  COMMAND_REMOVE_PACKAGE   = 5;
  COMMAND_INSTALL_PACKAGE  = 6;
  COMMAND_CONFIGURE_PROXY  = 7;
  COMMAND_CONFIGURE_GADGET = 8;
  COMMAND_START_AUTO_PILOT = 9;
  COMMAND_CLEAN            = 10;
  COMMAND_INFO             = 11;
  COMMAND_INFO_V2          = 12;
  COMMAND_LIST_PACKAGES    = 13;
}

// Sent from cloud relay → device through the bidi stream.
message Command {
  string request_id = 1;          // UUID for request/response correlation
  CommandType type = 2;           // Which RPC to invoke

  oneof payload {
    Message echo                     = 10;
    App launch_app                   = 11;
    Empty clear_proxy                = 12;
    Empty health_check               = 13;
    App remove_package               = 14;
    InstallReq install_package       = 15;
    ConfigProxyReq configure_proxy   = 16;
    App configure_gadget             = 17;
    AutoPilotConfig start_auto_pilot = 18;
    CleanOptions clean               = 19;
    Empty info                       = 20;
    Empty info_v2                    = 21;
    Empty list_packages              = 22;
  }
}


// ═══════════════════════════════════════════════════════════
//  COMMAND RESPONSE — device sends results back to relay
// ═══════════════════════════════════════════════════════════

message CommandResponse {
  string request_id = 1;          // Must match the Command's request_id
  bool success = 2;
  string error = 3;               // Error message if success=false

  oneof result {
    Message message    = 10;      // Echo, LaunchApp, ClearProxy, HealthCheck,
                                  // RemovePackage, InstallPackage, ConfigureProxy,
                                  // ConfigureGadget, StartAutoPilot, Clean
    Device device      = 11;      // Info
    DeviceV2 device_v2 = 12;      // InfoV2
    Apps apps          = 13;      // ListPackages
  }
}


// ═══════════════════════════════════════════════════════════
//  DEVICE REGISTRATION — handshake when stream opens
// ═══════════════════════════════════════════════════════════

message DeviceRegistration {
  string device_id = 1;           // Unique device UUID
  string custom_identifier = 2;   // Human-readable name / alias
  int32 platform = 3;             // Platform enum (matches Device.Platform)
  string agent_version = 4;       // Version of the device agent
  string auth_token = 5;          // Pre-shared secret or JWT
}

message RegistrationAck {
  bool accepted = 1;
  string message = 2;             // Reason if rejected
}


// ═══════════════════════════════════════════════════════════
//  BIDI STREAM WRAPPERS
// ═══════════════════════════════════════════════════════════

// Everything the device sends over the bidi stream.
message DeviceMessage {
  oneof payload {
    DeviceRegistration registration = 1;  // First message after connecting
    CommandResponse response = 2;         // All subsequent messages
  }
}

// Everything the cloud relay sends over the bidi stream.
message ServerMessage {
  oneof payload {
    RegistrationAck ack = 1;       // Reply to registration
    Command command = 2;           // Command to execute on device
  }
}


// ═══════════════════════════════════════════════════════════
//  SERVICE: MoriartyBridge — device-facing (bidi stream)
// ═══════════════════════════════════════════════════════════

service MoriartyBridge {
  // Device opens this on startup and keeps it alive.
  // Relay pushes Command messages, device pushes CommandResponse messages.
  rpc Connect(stream DeviceMessage) returns (stream ServerMessage);
}


// ═══════════════════════════════════════════════════════════
//  SERVICE: MoriartyDispatch — platform-facing (unary RPCs)
// ═══════════════════════════════════════════════════════════

// Mirrors all 13 original Moriarty RPCs but each request
// carries a device_id so the relay knows which stream to route through.

message TargetedMessage         { string device_id = 1; Message data = 2; }
message TargetedApp             { string device_id = 1; App data = 2; }
message TargetedEmpty           { string device_id = 1; }
message TargetedInstallReq      { string device_id = 1; InstallReq data = 2; }
message TargetedConfigProxyReq  { string device_id = 1; ConfigProxyReq data = 2; }
message TargetedAutoPilotConfig { string device_id = 1; AutoPilotConfig data = 2; }
message TargetedCleanOptions    { string device_id = 1; CleanOptions data = 2; }

service MoriartyDispatch {
  rpc Echo(TargetedMessage) returns (Message);
  rpc LaunchApp(TargetedApp) returns (Message);
  rpc ClearProxy(TargetedEmpty) returns (Message);
  rpc HealthCheck(TargetedEmpty) returns (Message);
  rpc RemovePackage(TargetedApp) returns (Message);
  rpc InstallPackage(TargetedInstallReq) returns (Message);
  rpc ConfigureProxy(TargetedConfigProxyReq) returns (Message);
  rpc ConfigureGadget(TargetedApp) returns (Message);
  rpc StartAutoPilot(TargetedAutoPilotConfig) returns (Message);
  rpc Clean(TargetedCleanOptions) returns (Message);
  rpc Info(TargetedEmpty) returns (Device);
  rpc InfoV2(TargetedEmpty) returns (DeviceV2);
  rpc ListPackages(TargetedEmpty) returns (Apps);

  // Additional: discover all currently connected relay devices
  rpc ListDevices(Empty) returns (DeviceList);
}

message DeviceList {
  repeated ConnectedDevice devices = 1;
}

message ConnectedDevice {
  string device_id = 1;
  string custom_identifier = 2;
  int32 platform = 3;
  string agent_version = 4;
  int64 connected_at = 5;         // Unix timestamp
}
```

---

## How It Works — Step by Step

### 1. Device Connects to Relay

```
Device                          Cloud Relay
  │                                 │
  │──── Connect() bidi stream ────►│
  │                                 │
  │──── DeviceRegistration ───────►│  Validates auth_token
  │                                 │  Stores device_id → stream
  │                                 │
  │◄─── RegistrationAck ──────────│
  │                                 │
  │     (stream stays open)         │
```

### 2. Platform Dispatches a Command

```
Platform                       Cloud Relay                        Device
  │                                 │                                │
  │── MoriartyDispatch.LaunchApp ─►│                                │
  │   {device_id + App}            │                                │
  │                                 │── Command{LaunchApp} ────────►│
  │                                 │   (via bidi stream)            │
  │                                 │                                │  executes locally
  │                                 │◄── CommandResponse ───────────│
  │                                 │                                │
  │◄── Message (response) ────────│                                │
```

### 3. Keepalive & Reconnection

```
Device                          Cloud Relay
  │                                 │
  │◄──── gRPC keepalive pings ────►│  (HTTP/2 PING frames)
  │                                 │
  │  (network drops)                │
  │         ✕                       │  Relay removes device from registry
  │                                 │
  │──── Connect() (reconnect) ────►│  Device reconnects with backoff
  │──── DeviceRegistration ───────►│
  │◄─── RegistrationAck ──────────│
```

---

## Cloud Relay — Internal Design

The relay is the **only new component** to build. It hosts both `MoriartyBridge` and `MoriartyDispatch` on the same gRPC server.

### Device Registry (in-memory, concurrent-safe)

```
┌──────────────────────────────────────────────────┐
│  device_id  │  stream handle  │  metadata        │
│  ───────────┼─────────────────┼────────────────  │
│  "uuid-1"   │  stream_ref_1   │  DeviceReg       │
│  "uuid-2"   │  stream_ref_2   │  DeviceReg       │
└──────────────────────────────────────────────────┘
```

### Pending Request Map (in-memory, concurrent-safe)

```
┌────────────────────────────────────────────────────┐
│  request_id  │  response channel                   │
│  ────────────┼───────────────────────────────────  │
│  "req-abc"   │  chan CommandResponse                │
│  "req-def"   │  chan CommandResponse                │
└────────────────────────────────────────────────────┘
```

### Relay Pseudocode

```
// Device-facing: runs per device connection
func (s *Relay) Connect(stream) {
    msg = stream.Recv()                          // expect DeviceRegistration
    if !validate(msg.registration):
        stream.Send(RegistrationAck{accepted: false})
        return

    stream.Send(RegistrationAck{accepted: true})
    registry.Add(msg.registration.device_id, stream)
    defer registry.Remove(msg.registration.device_id)

    for {
        msg = stream.Recv()
        pendingRequests.Resolve(msg.response.request_id, msg.response)
    }
}

// Platform-facing: e.g. LaunchApp
func (s *Relay) LaunchApp(ctx, req TargetedApp) (Message, error) {
    stream = registry.Get(req.device_id)
    if stream == nil:
        return error(UNAVAILABLE, "device not connected")

    requestID = uuid()
    ch = pendingRequests.Create(requestID)
    defer pendingRequests.Remove(requestID)

    stream.Send(ServerMessage{command: Command{
        request_id: requestID,
        type: COMMAND_LAUNCH_APP,
        payload: req.data,
    }})

    select {
    case resp = <-ch:
        if !resp.success: return error(resp.error)
        return resp.result.message, nil
    case <-ctx.Done():
        return error(DEADLINE_EXCEEDED)
    }
}
```

---

## Platform Routing Logic

The platform picks the right client per device:

```
func executeOnDevice(device, rpcName, request):
    if device.connection_mode == "relay":
        client = MoriartyDispatchClient(relay_address)
        client.LaunchApp(TargetedApp{device_id: device.id, data: request})

    else:  // "direct" — existing behavior, unchanged
        client = MoriartyClient(device.ip + ":" + device.port)
        client.LaunchApp(request)
```

Both paths coexist permanently. No migration deadline.

---

## Resilience

### Keepalive Configuration

| Parameter              | Device (client) | Relay (server) |
|------------------------|-----------------|----------------|
| Keepalive interval     | 30s             | —              |
| Keepalive timeout      | 10s             | —              |
| Max idle time          | —               | 60s            |
| Min client ping interval | —             | 20s            |

### Device Reconnection

```
on disconnect:
  wait = 1s
  loop:
    try Connect(relay_address)
    if success → send DeviceRegistration → break
    wait = min(wait * 2, 60s)
    sleep(wait + jitter)
```

### Relay Timeout Behavior

| Scenario                          | Relay returns to platform |
|-----------------------------------|--------------------------|
| Device responds in time           | The actual response       |
| Device doesn't respond (30s)      | `DEADLINE_EXCEEDED`       |
| Device not connected              | `UNAVAILABLE`             |
| Device sends `success=false`      | `INTERNAL` with error     |

---

## Authentication & Security

| Path              | Mechanism |
|-------------------|-----------|
| Device → Relay    | TLS + `auth_token` in `DeviceRegistration` (or mTLS) |
| Platform → Relay  | Existing Appknox auth (API keys / JWT / service mTLS) |

---

## Impact on Existing Code

| Component | Change required? |
|-----------|-----------------|
| `bakerstreet.proto` | **None** |
| Generated `bakerstreet.*` stubs | **None** |
| Existing `Moriarty` service consumers | **None** |
| Direct-mode device agents | **None** |
| Direct-mode platform code | **None** |
| Build scripts (`hatch_build.py`, `bakerstreet.sh`) | **Minor update** — also compile `bakerstreet_bridge.proto` |

---

## Implementation Steps

1. **Create `bakerstreet/bakerstreet_bridge.proto`** — the new proto file as shown above.
2. **Update build scripts** — add compilation of `bakerstreet_bridge.proto` to `hatch_build.py` and `scripts/bakerstreet.sh`.
3. **Generate & verify** — run build, confirm existing stubs unchanged, new stubs created.
4. **Build cloud relay server** — new standalone service implementing both `MoriartyBridge` and `MoriartyDispatch`.
5. **Add relay client to device agent** — new module alongside existing `MoriartyServer`, both can run simultaneously.
6. **Add routing to platform** — thin routing layer that picks direct or relay per device.

---

## Open Questions

1. **Relay scaling** — single relay works for moderate counts. For hundreds of streams, consider horizontal scaling with shared registry (Redis) or consistent-hash routing.
2. **Concurrent commands** — can multiple commands be in-flight to one device, or queued? Parallel is faster but requires device-side concurrency.
3. **Reconnect behavior** — drop pending commands on reconnect (simpler) or replay them?
4. **Observability** — relay should expose metrics: connected count, latency percentiles, error rates, reconnection frequency.
5. **Message size** — gRPC default 4MB is fine for current payloads. Configure if future RPCs carry larger data.
