# Reverse RPC Plan — Device-Initiated Connections

## Problem

The current architecture requires the **Appknox platform to directly reach device agents** over the network. Devices run gRPC servers and the platform connects to them as a client. This forces both to be on the **same network** (or requires port forwarding, VPNs, etc.), which limits where devices can be deployed.

```
CURRENT (broken by NAT / firewalls)

┌──────────────┐    gRPC connect     ┌──────────────┐
│   Platform   │ ──────────────────► │    Device    │
│   (client)   │                     │   (server)   │
│              │ ◄── response ────── │              │
└──────────────┘                     └──────────────┘
      Needs to know device IP.
      Device must be reachable.
```

## Proposed Architecture

Reverse the roles. A **cloud relay server** with a public IP hosts the gRPC server. **Devices initiate outbound connections** to it and hold open a **bidirectional stream**. The platform dispatches commands through the relay, which routes them down the device's open stream.

```
PROPOSED

┌──────────────┐   Dispatch API    ┌─────────────────┐    bidi stream     ┌──────────────┐
│   Platform   │ ────────────────► │  Cloud Relay    │ ◄─────────────────  │    Device    │
│              │                   │  (public IP)    │ ──────────────────► │   (client)   │
│              │ ◄── response ──── │                 │                     │              │
└──────────────┘                   └─────────────────┘                     └──────────────┘
                                    Devices connect out.
                                    No inbound ports needed.
```

**Key properties:**
- Devices only make **outbound connections** — works through NAT, firewalls, corporate networks
- The cloud relay is the **single point of contact** — one public IP, TLS-terminated
- All 13 existing RPCs (Echo, LaunchApp, Info, etc.) continue to work unchanged in semantics
- The platform no longer needs to know device IPs

---

## Proto Design

The approach uses **command envelopes** over a bidirectional stream. Every existing RPC is wrapped into a tagged union so the server can push any command down the stream and the device can send back the typed response.

### New Messages

```proto
syntax = "proto3";
package com.appknox.bakerstreet;

// ─── Existing messages stay unchanged ───
// Message, App, Apps, Device, DeviceV2, Finding,
// InstallReq, ConfigProxyReq, AutoPilotConfig, CleanOptions, Empty
// (all kept exactly as-is)


// ─── New: Command envelope (server → device) ───

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

// Sent from cloud server → device through the bidi stream.
// Wraps every existing RPC into a single message with a correlation ID.
message Command {
  string request_id = 1;          // UUID — correlates response to request
  CommandType type = 2;           // Which RPC to invoke

  // Payload — only one is set, matching the type
  oneof payload {
    Message echo             = 10;
    App launch_app           = 11;
    Empty clear_proxy        = 12;
    Empty health_check       = 13;
    App remove_package       = 14;
    InstallReq install_package  = 15;
    ConfigProxyReq configure_proxy = 16;
    App configure_gadget     = 17;
    AutoPilotConfig start_auto_pilot = 18;
    CleanOptions clean       = 19;
    Empty info               = 20;
    Empty info_v2            = 21;
    Empty list_packages      = 22;
  }
}


// ─── New: Command response envelope (device → server) ───

// Sent from device → cloud server through the bidi stream.
message CommandResponse {
  string request_id = 1;          // Must match the Command's request_id
  bool success = 2;               // True if the command executed without error
  string error = 3;               // Error message if success=false

  // Response payload — only one is set, matching the original command type
  oneof result {
    Message message    = 10;      // For Echo, LaunchApp, ClearProxy, HealthCheck,
                                  //     RemovePackage, InstallPackage, ConfigureProxy,
                                  //     ConfigureGadget, StartAutoPilot, Clean
    Device device      = 11;      // For Info
    DeviceV2 device_v2 = 12;      // For InfoV2
    Apps apps          = 13;      // For ListPackages
  }
}


// ─── New: Device registration ───

// First message a device sends after opening the stream.
// Identifies the device to the relay so commands can be routed to it.
message DeviceRegistration {
  string device_id = 1;           // Unique device UUID
  string custom_identifier = 2;   // Human-readable name / alias
  int32 platform = 3;             // Android=0, iOS=1
  string agent_version = 4;       // Version of the agent software
}

message RegistrationAck {
  bool accepted = 1;
  string message = 2;             // Reason if rejected
}


// ─── New: Stream wrapper messages ───

// Everything the device sends over the bidi stream.
message DeviceMessage {
  oneof payload {
    DeviceRegistration registration = 1;  // First message
    CommandResponse response = 2;         // All subsequent messages
  }
}

// Everything the cloud server sends over the bidi stream.
message ServerMessage {
  oneof payload {
    RegistrationAck ack = 1;       // Response to registration
    Command command = 2;           // Command to execute
  }
}
```

### New Services

```proto
// ─── Device-facing service (devices connect to this) ───

service MoriartyBridge {
  // Device opens this stream on startup and keeps it alive.
  // Server pushes commands, device pushes responses.
  rpc Connect(stream DeviceMessage) returns (stream ServerMessage);
}


// ─── Platform-facing service (Appknox backend calls this) ───

// Keeps the same RPC signatures as the original Moriarty service,
// but every request includes a device_id so the relay knows
// which device stream to route the command through.

message TargetedMessage   { string device_id = 1; Message data = 2; }
message TargetedApp       { string device_id = 1; App data = 2; }
message TargetedEmpty     { string device_id = 1; }
message TargetedInstallReq     { string device_id = 1; InstallReq data = 2; }
message TargetedConfigProxyReq { string device_id = 1; ConfigProxyReq data = 2; }
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

  // New: list all currently connected devices
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

### 1. Device Connects

```
Device                          Cloud Relay
  │                                 │
  │──── Connect() bidi stream ────►│
  │                                 │
  │──── DeviceRegistration ───────►│  Relay records device_id → stream mapping
  │                                 │
  │◄─── RegistrationAck ──────────│
  │                                 │
  │     (stream stays open)         │
```

- Device opens a `MoriartyBridge.Connect()` bidi stream on startup.
- First message is always `DeviceRegistration` with the device's identity.
- Relay validates and stores the mapping: `device_id → open stream`.
- Stream stays alive indefinitely (with keepalives).

### 2. Platform Sends a Command

```
Platform                       Cloud Relay                        Device
  │                                 │                                │
  │── MoriartyDispatch.LaunchApp ─►│                                │
  │   (device_id + App)            │                                │
  │                                 │── Command{LaunchApp} ────────►│
  │                                 │   (via bidi stream)            │
  │                                 │                                │  Device executes
  │                                 │◄── CommandResponse ───────────│  LaunchApp locally
  │                                 │                                │
  │◄── Message (response) ────────│                                │
  │                                 │                                │
```

- Platform calls `MoriartyDispatch.LaunchApp(TargetedApp{device_id: "xxx", data: App{Name: "com.example"}})`.
- Relay looks up the device's open stream by `device_id`.
- Relay wraps the request into a `Command` envelope with a unique `request_id` and pushes it down the stream.
- Device receives the `Command`, executes `LaunchApp` locally, wraps the result in `CommandResponse` with the matching `request_id`.
- Relay receives the `CommandResponse`, matches it to the pending platform request, and returns the response.

### 3. Keepalive & Reconnection

```
Device                          Cloud Relay
  │                                 │
  │◄──── gRPC keepalive pings ────►│  (HTTP/2 PING frames)
  │                                 │
  │  (network drops)                │
  │         ✕                       │  Relay detects stream closed,
  │                                 │  removes device from registry
  │                                 │
  │──── Connect() (reconnect) ────►│  Device reconnects with backoff
  │──── DeviceRegistration ───────►│
  │◄─── RegistrationAck ──────────│
```

---

## Cloud Relay Server — Internal Design

The relay server is the new component that needs to be built. It has two responsibilities:

### Device Registry (in-memory)

```
┌─────────────────────────────────────────────┐
│              Device Registry                │
│                                             │
│  device_id  │  stream_handle  │  metadata   │
│  ───────────┼─────────────────┼───────────  │
│  "uuid-1"   │  stream_ref_1   │  DeviceReg  │
│  "uuid-2"   │  stream_ref_2   │  DeviceReg  │
│  "uuid-3"   │  stream_ref_3   │  DeviceReg  │
└─────────────────────────────────────────────┘
```

- When a device calls `Connect()` and sends `DeviceRegistration`, the relay stores the stream reference keyed by `device_id`.
- When a device disconnects (stream closes/errors), the relay removes it.
- `ListDevices` returns the current registry contents.

### Pending Request Map (in-memory)

```
┌──────────────────────────────────────────────────────┐
│              Pending Requests                        │
│                                                      │
│  request_id  │  device_id  │  response_channel       │
│  ────────────┼─────────────┼──────────────────────── │
│  "req-abc"   │  "uuid-1"   │  chan CommandResponse    │
│  "req-def"   │  "uuid-2"   │  chan CommandResponse    │
└──────────────────────────────────────────────────────┘
```

- When the platform calls a `MoriartyDispatch` RPC, the relay:
  1. Generates a `request_id` (UUID).
  2. Creates a response channel and stores it in the pending map.
  3. Sends the `Command` down the device's bidi stream.
  4. Blocks waiting on the response channel (with timeout).
- When a `CommandResponse` arrives from a device, the relay:
  1. Looks up the `request_id` in the pending map.
  2. Sends the response to the waiting channel.
  3. Removes the entry from the pending map.

---

## What Changes for Each Component

### Device Agent (currently: gRPC server → becomes: gRPC client)

| Aspect | Before | After |
|--------|--------|-------|
| Role | gRPC server (listens on port) | gRPC client (connects outbound) |
| Startup | `grpc.NewServer()` + `Listen()` | `grpc.Dial(relay_address)` + `Connect()` |
| Receives commands via | Incoming RPC calls | `Command` messages on bidi stream |
| Sends responses via | RPC return values | `CommandResponse` messages on bidi stream |
| Network requirement | Must accept inbound connections | Only needs outbound HTTPS (port 443) |
| Reconnection | N/A (was server) | Exponential backoff reconnect loop |

**The actual command execution logic stays the same.** The device still runs `LaunchApp`, `InstallPackage`, etc. locally. Only the transport layer changes — instead of implementing `MoriartyServer`, the device runs a message loop that dispatches `Command` envelopes to the same handler functions.

### Platform (currently: gRPC client → stays: gRPC client, different target)

| Aspect | Before | After |
|--------|--------|-------|
| Connects to | Device IP directly | Cloud relay server |
| Service used | `Moriarty` | `MoriartyDispatch` |
| Must specify | Device IP/port | `device_id` in request |
| Network requirement | Same network as device | Network access to cloud relay |

**Minimal change.** The platform swaps `MoriartyClient` for `MoriartyDispatchClient` and passes `device_id` instead of connecting to device IPs.

### Cloud Relay (new component)

This is the **only new component** that needs to be built:
- Implements `MoriartyBridge` service (device-facing)
- Implements `MoriartyDispatch` service (platform-facing)
- Maintains device registry + pending request map
- Handles stream lifecycle, keepalives, timeouts

---

## Connection Lifecycle & Resilience

### Keepalives

gRPC bidirectional streams run over HTTP/2, which supports native PING frames. Configure both sides:

| Parameter | Device (client) | Relay (server) |
|-----------|-----------------|----------------|
| Keepalive interval | 30s | — |
| Keepalive timeout | 10s | — |
| Max idle time | — | 60s |
| Min client ping interval | — | 20s |

### Reconnection Strategy (device side)

```
on disconnect:
  wait = 1s
  loop:
    try Connect()
    if success:
      send DeviceRegistration
      break
    wait = min(wait * 2, 60s)   // exponential backoff, cap at 60s
    sleep(wait + random(0, wait * 0.2))  // jitter
```

### Request Timeout (relay side)

- Every `MoriartyDispatch` RPC waits for a `CommandResponse` with a **configurable timeout** (default: 30s).
- If the device doesn't respond in time, the relay returns `DEADLINE_EXCEEDED` to the platform.
- If the device's stream is not found (disconnected), the relay returns `UNAVAILABLE`.

---

## Authentication & Security

### Device → Relay

- **TLS**: All connections use TLS (standard for gRPC over public internet).
- **Device tokens**: Each device gets a pre-shared token (or mTLS certificate) included in `DeviceRegistration` or gRPC metadata. The relay validates before accepting the stream.

```proto
message DeviceRegistration {
  string device_id = 1;
  string custom_identifier = 2;
  int32 platform = 3;
  string agent_version = 4;
  string auth_token = 5;           // Pre-shared secret or JWT
}
```

### Platform → Relay

- Standard API authentication (API keys, JWT, or mTLS) on the `MoriartyDispatch` service.
- Can reuse existing Appknox platform auth mechanisms.

---

## Backwards Compatibility & Migration

### Phase 1 — Add new proto definitions

- Add `MoriartyBridge`, `MoriartyDispatch`, and all new messages to `bakerstreet.proto`.
- The existing `Moriarty` service definition stays untouched.
- Regenerate Go and Python stubs.
- This is a **non-breaking change** — existing consumers are unaffected.

### Phase 2 — Build the cloud relay server

- Implement the relay as a standalone service.
- Deploy with a public IP and TLS.
- The relay implements both `MoriartyBridge` and `MoriartyDispatch`.

### Phase 3 — Update device agents

- Add bidi stream client logic alongside the existing gRPC server.
- Devices can run **both modes simultaneously** during migration:
  - Old: listening as gRPC server (for devices still on same network)
  - New: connecting to relay via bidi stream
- Controlled rollout — enable relay mode per-device via config flag.

### Phase 4 — Update platform

- Add `MoriartyDispatchClient` as an alternative to `MoriartyClient`.
- Route commands through the relay for devices that are connected via bidi stream.
- Fall back to direct connection for devices still in old mode.

### Phase 5 — Deprecate direct mode

- Once all devices are migrated to relay mode, remove the old `Moriarty` service and direct connection code.
- Remove the old `Moriarty` service definition from the proto (major version bump to v3).

---

## Proto File Changes Summary

| What | Action |
|------|--------|
| Existing messages (Message, App, Device, etc.) | **No change** |
| Existing `Moriarty` service | **Keep as-is** during migration, remove in v3 |
| `Command` + `CommandType` enum | **Add** — server→device command envelope |
| `CommandResponse` | **Add** — device→server response envelope |
| `DeviceRegistration` + `RegistrationAck` | **Add** — device identity handshake |
| `DeviceMessage` + `ServerMessage` | **Add** — bidi stream wrappers |
| `MoriartyBridge` service | **Add** — device-facing bidi stream endpoint |
| `Targeted*` request wrappers | **Add** — platform requests with device routing |
| `MoriartyDispatch` service | **Add** — platform-facing dispatch API |
| `DeviceList` + `ConnectedDevice` | **Add** — device discovery |

---

## Open Questions

1. **Single relay or multiple?** If many devices connect, a single relay may need horizontal scaling. gRPC streams are sticky to one server, so a load-balanced setup needs sticky sessions or a shared device registry (Redis, etc.).

2. **Stream multiplexing** — Should one device open one stream, or could it open multiple? One stream per device is simpler and recommended.

3. **Command ordering** — Should commands be strictly ordered per device, or can multiple commands be in-flight simultaneously? Parallel in-flight commands are more efficient but require the device to handle concurrency.

4. **Large payloads** — `InstallPackage` passes a URL (small), but if future RPCs need large payloads, consider gRPC message size limits (default 4MB). The current design should be fine since all payloads are small metadata.

5. **Observability** — The relay should expose metrics: connected device count, command latency, error rates, stream reconnection frequency.
