# Bridge Mode Integration Guide

> **Baker Street** version: **2.3.0+** on branch `feature/cyod-pase-1`.
>
> This document is the implementation spec for adding relay/bridge mode to:
> - **Moriarty** — Python/Django platform service
> - **Mary Agent** — Go device agent
>
> Use this as context when implementing changes in either repository.

---

## What Changed in Baker Street

A new proto file `bakerstreet/bakerstreet_bridge.proto` has been added **alongside** the existing `bakerstreet/bakerstreet.proto`. The original proto is completely untouched.

This adds a second communication path — **relay mode** — where devices connect outbound to Moriarty instead of Moriarty connecting to devices. Both modes coexist permanently.

**Nothing about the existing direct mode changes. All current code continues to work as-is.**

---

## Architecture: Both Modes Side by Side

```
MODE A — Direct (existing, unchanged)

  Uses: bakerstreet.proto → Moriarty service

  ┌──────────────┐    gRPC (MoriartyStub)    ┌──────────────┐
  │   Moriarty   │ ────────────────────────►  │  Mary Agent  │
  │   (client)   │ ◄──── response ─────────  │   (server)   │
  └──────────────┘                            └──────────────┘

  Moriarty connects to Mary's IP:port.
  Both must be on the same network.


MODE B — Relay (new, additive)

  Uses: bakerstreet_bridge.proto → MoriartyBridge + MoriartyDispatch services

  ┌──────────────┐                  ┌──────────────────┐                  ┌──────────────┐
  │   Moriarty   │  dispatch API    │  Relay Server    │   bidi stream    │  Mary Agent  │
  │   Platform   │ ───────────────► │  (inside         │ ◄──────────────  │   (client)   │
  │   Code       │                  │   Moriarty)      │ ──────────────►  │              │
  │              │ ◄── response ─── │                  │                  │              │
  └──────────────┘                  └──────────────────┘                  └──────────────┘

  Mary connects outbound to relay. Works through NAT/firewalls.
  Moriarty dispatches commands by device_id, relay routes them to the right stream.
```

Each device is configured to use one mode or the other (e.g., `connection_mode: "direct" | "relay"`). Mary Agent runs in the configured mode at startup. Moriarty routes commands accordingly.

---

## New Proto: `bakerstreet_bridge.proto`

This file imports all existing message types from `bakerstreet.proto` (Message, App, Device, DeviceV2, etc.) and defines:

### Services

| Service | Purpose | Who implements | Who calls |
|---------|---------|----------------|-----------|
| `MoriartyBridge` | Device-facing bidi stream | Relay (Moriarty) | Mary Agent |
| `MoriartyDispatch` | Platform-facing unary RPCs | Relay (Moriarty) | Moriarty platform code |

### Key Message Types

```
Command                  — Relay → Mary: "execute this RPC"
  ├── request_id         — UUID for correlation
  ├── type               — CommandType enum (COMMAND_ECHO, COMMAND_LAUNCH_APP, etc.)
  └── payload            — oneof with the actual request data (Message, App, InstallReq, etc.)

CommandResponse          — Mary → Relay: "here's the result"
  ├── request_id         — matches the Command's request_id
  ├── success            — true/false
  ├── error              — error message if success=false
  └── result             — oneof with the response data (Message, Device, DeviceV2, Apps)

DeviceRegistration       — Mary → Relay: "I am device X" (first message on connect)
  ├── device_id
  ├── custom_identifier
  ├── platform
  ├── agent_version
  └── auth_token

RegistrationAck          — Relay → Mary: "accepted" or "rejected"

DeviceMessage            — Wrapper for everything Mary sends (registration or response)
ServerMessage            — Wrapper for everything relay sends (ack or command)
```

### Targeted Request Wrappers (for MoriartyDispatch)

Every original Moriarty RPC is mirrored with a `device_id` field for routing:

| Original (direct) | Targeted (relay) |
|---|---|
| `Echo(Message)` | `Echo(TargetedMessage{device_id, data})` |
| `LaunchApp(App)` | `LaunchApp(TargetedApp{device_id, data})` |
| `ClearProxy(Empty)` | `ClearProxy(TargetedEmpty{device_id})` |
| `HealthCheck(Empty)` | `HealthCheck(TargetedEmpty{device_id})` |
| `RemovePackage(App)` | `RemovePackage(TargetedApp{device_id, data})` |
| `InstallPackage(InstallReq)` | `InstallPackage(TargetedInstallReq{device_id, data})` |
| `ConfigureProxy(ConfigProxyReq)` | `ConfigureProxy(TargetedConfigProxyReq{device_id, data})` |
| `ConfigureGadget(App)` | `ConfigureGadget(TargetedApp{device_id, data})` |
| `StartAutoPilot(AutoPilotConfig)` | `StartAutoPilot(TargetedAutoPilotConfig{device_id, data})` |
| `Clean(CleanOptions)` | `Clean(TargetedCleanOptions{device_id, data})` |
| `Info(Empty)` | `Info(TargetedEmpty{device_id})` |
| `InfoV2(Empty)` | `InfoV2(TargetedEmpty{device_id})` |
| `ListPackages(Empty)` | `ListPackages(TargetedEmpty{device_id})` |
| *(new)* | `ListDevices(Empty) → DeviceList` |

Response types are **identical** to direct mode. Only request types differ (they carry `device_id`).

---

## How the Relay Works — Step by Step

### 1. Mary Agent Connects

```
Mary Agent                         Relay (Moriarty)
  │                                     │
  │──── Connect() bidi stream ────────► │
  │                                     │
  │──── DeviceMessage{registration} ──► │   Validates auth_token
  │      device_id: "uuid-123"         │   Stores device_id → stream mapping
  │      auth_token: "secret"          │
  │                                     │
  │◄─── ServerMessage{ack} ──────────── │   RegistrationAck{accepted: true}
  │                                     │
  │     (stream stays open forever)     │
```

### 2. Platform Dispatches a Command

```
Platform Code              Relay                              Mary Agent
  │                           │                                   │
  │── LaunchApp ────────────► │                                   │
  │   device_id: "uuid-123"  │                                   │
  │   data: App{Name: "..."}│                                   │
  │                           │   Looks up "uuid-123" in registry │
  │                           │   Generates request_id: "req-abc" │
  │                           │   Creates pending response channel │
  │                           │                                   │
  │                           │── ServerMessage{command} ────────►│
  │                           │   Command{                        │
  │                           │     request_id: "req-abc"         │
  │                           │     type: COMMAND_LAUNCH_APP      │
  │                           │     payload: App{Name: "..."}     │
  │                           │   }                               │
  │                           │                                   │  Mary executes
  │                           │                                   │  LaunchApp locally
  │                           │                                   │
  │                           │◄── DeviceMessage{response} ──────│
  │                           │   CommandResponse{                │
  │                           │     request_id: "req-abc"         │
  │                           │     success: true                 │
  │                           │     result: Message{Data: "ok"}   │
  │                           │   }                               │
  │                           │                                   │
  │                           │   Matches "req-abc" to pending ch │
  │◄── Message{Data: "ok"} ──│                                   │
```

### 3. Disconnect & Reconnect

```
Mary Agent                         Relay
  │                                   │
  │◄──── gRPC keepalive pings ──────► │
  │                                   │
  │  (network drops)                  │
  │         ✕                         │   Relay detects stream closed
  │                                   │   Removes "uuid-123" from registry
  │                                   │   Any pending commands → DEADLINE_EXCEEDED
  │                                   │
  │──── Connect() (reconnect) ──────► │   Mary reconnects with backoff
  │──── DeviceRegistration ─────────► │
  │◄─── RegistrationAck ────────────  │
```

---

## Relay Server Internals (inside Moriarty)

The relay runs as a gRPC server within Moriarty, implementing two services:

### Device Registry

An in-memory, thread-safe dictionary mapping `device_id` to the device's open stream and metadata.

```python
# Conceptual structure
device_registry = {
    "uuid-123": {
        "stream": <bidi_stream_context>,       # The open Connect() stream
        "registration": DeviceRegistration(...), # Metadata from handshake
        "connected_at": 1707700000,             # Unix timestamp
        "send_lock": threading.Lock(),          # Serializes writes to this stream
    },
    "uuid-456": { ... },
}
```

- **Add**: when Mary sends `DeviceRegistration` and it's accepted
- **Remove**: when the stream closes (disconnect, error, EOF)
- **Get**: when a dispatch RPC needs to route to a device
- **List**: for `ListDevices` RPC

### Pending Request Map

An in-memory, thread-safe dictionary mapping `request_id` to a response mechanism (e.g., `threading.Event` + result slot, or `asyncio.Future`).

```python
# Conceptual structure
pending_requests = {
    "req-abc": Future(),  # Dispatch goroutine is blocked waiting on this
    "req-def": Future(),
}
```

- **Create**: dispatch RPC generates a UUID `request_id`, stores a Future
- **Resolve**: when `CommandResponse` arrives from a device, look up `request_id`, set the result on the Future
- **Timeout**: if the Future isn't resolved within 30s, return `DEADLINE_EXCEEDED` to the caller
- **Cleanup**: always remove the entry after resolve or timeout

### MoriartyBridge.Connect (device-facing)

```python
class MoriartyBridgeServicer:
    def Connect(self, request_iterator, context):
        # Phase 1: Registration
        first_msg = next(request_iterator)
        reg = first_msg.registration
        if not validate(reg):
            yield ServerMessage(ack=RegistrationAck(accepted=False, message="..."))
            return

        yield ServerMessage(ack=RegistrationAck(accepted=True))
        registry.add(reg.device_id, stream=context, registration=reg)

        try:
            # Phase 2: Receive responses
            for msg in request_iterator:
                resp = msg.response
                pending_requests.resolve(resp.request_id, resp)
        finally:
            registry.remove(reg.device_id)
```

> **Note**: The bidi stream in Python gRPC requires careful handling. The servicer yields `ServerMessage` objects (for acks and commands) while iterating over incoming `DeviceMessage` objects. Sending commands mid-stream requires a concurrent mechanism (e.g., a queue the servicer drains, or the `context.send()` pattern with `grpc.ServicerContext`).

### MoriartyDispatch RPCs (platform-facing)

Each of the 13 dispatch RPCs follows the same pattern:

```python
class MoriartyDispatchServicer:
    def LaunchApp(self, request, context):
        device_id = request.device_id
        conn = registry.get(device_id)
        if conn is None:
            context.abort(grpc.StatusCode.UNAVAILABLE, f"device {device_id} not connected")

        request_id = str(uuid4())
        future = pending_requests.create(request_id)

        command = Command(
            request_id=request_id,
            type=CommandType.COMMAND_LAUNCH_APP,
            launch_app=request.data,
        )
        conn.send(ServerMessage(command=command))

        try:
            response = future.result(timeout=30)  # blocks until Mary responds
        except TimeoutError:
            context.abort(grpc.StatusCode.DEADLINE_EXCEEDED, "device did not respond")
        finally:
            pending_requests.remove(request_id)

        if not response.success:
            context.abort(grpc.StatusCode.INTERNAL, response.error)

        return response.message  # The Message proto
```

### Error Codes

| Condition | gRPC Status |
|-----------|-------------|
| Device not in registry | `UNAVAILABLE` |
| Failed to send to device stream | `UNAVAILABLE` |
| Device didn't respond in 30s | `DEADLINE_EXCEEDED` |
| Device responded with `success=false` | `INTERNAL` (with device's error message) |
| Missing `device_id` in request | `INVALID_ARGUMENT` |
| Invalid auth token on registration | `UNAUTHENTICATED` |
| Caller cancelled the request | `CANCELLED` |

---

## What Mary Agent (Go) Needs to Do

Mary Agent currently implements `MoriartyServer` (the gRPC server that Moriarty connects to directly). **This stays exactly as-is.**

Additionally, Mary needs a new **bridge client** that connects outbound to the relay.

### New Import

```go
import (
    pb "github.com/appknox/bakerstreet/v2/bakerstreet"
)
```

The bridge types (`MoriartyBridgeClient`, `DeviceMessage`, `ServerMessage`, `Command`, `CommandResponse`, etc.) are all generated in `bakerstreet_bridge.pb.go` and `bakerstreet_bridge_grpc.pb.go`.

### Connection & Registration

```go
func connectToRelay(relayAddr string, deviceInfo DeviceInfo, authToken string) {
    for {
        err := runBridgeSession(relayAddr, deviceInfo, authToken)
        if err != nil {
            log.Printf("bridge session ended: %v, reconnecting...", err)
        }
        backoffSleep() // exponential backoff: 1s → 2s → 4s → ... → 60s max
    }
}

func runBridgeSession(relayAddr string, deviceInfo DeviceInfo, authToken string) error {
    conn, err := grpc.Dial(relayAddr,
        grpc.WithTransportCredentials(credentials.NewTLS(&tls.Config{})),
        grpc.WithKeepaliveParams(keepalive.ClientParameters{
            Time:    30 * time.Second,  // send keepalive ping every 30s
            Timeout: 10 * time.Second,  // wait 10s for pong
        }),
    )
    if err != nil {
        return err
    }
    defer conn.Close()

    client := pb.NewMoriartyBridgeClient(conn)
    stream, err := client.Connect(context.Background())
    if err != nil {
        return err
    }

    // Step 1: Send registration
    err = stream.Send(&pb.DeviceMessage{
        Payload: &pb.DeviceMessage_Registration{
            Registration: &pb.DeviceRegistration{
                DeviceId:         deviceInfo.UUID,
                CustomIdentifier: deviceInfo.Name,
                Platform:         deviceInfo.Platform,
                AgentVersion:     deviceInfo.AgentVersion,
                AuthToken:        authToken,
            },
        },
    })
    if err != nil {
        return err
    }

    // Step 2: Receive registration ack
    serverMsg, err := stream.Recv()
    if err != nil {
        return err
    }
    ack := serverMsg.GetAck()
    if ack == nil || !ack.GetAccepted() {
        return fmt.Errorf("registration rejected: %s", ack.GetMessage())
    }

    log.Println("registered with relay, waiting for commands...")

    // Step 3: Command loop
    for {
        serverMsg, err := stream.Recv()
        if err != nil {
            return err // triggers reconnect
        }

        cmd := serverMsg.GetCommand()
        if cmd == nil {
            continue
        }

        // Execute command and send response
        resp := executeCommand(cmd)
        err = stream.Send(&pb.DeviceMessage{
            Payload: &pb.DeviceMessage_Response{
                Response: resp,
            },
        })
        if err != nil {
            return err
        }
    }
}
```

### Command Dispatch

The `executeCommand` function bridges the `Command` envelope to Mary's existing handler logic. **The actual handler implementations don't change at all** — only the transport wrapper is new.

```go
func executeCommand(cmd *pb.Command) *pb.CommandResponse {
    resp := &pb.CommandResponse{
        RequestId: cmd.GetRequestId(),
        Success:   true,
    }

    switch cmd.GetType() {
    case pb.CommandType_COMMAND_ECHO:
        result, err := handleEcho(cmd.GetEcho())           // existing handler
        if err != nil {
            resp.Success = false
            resp.Error = err.Error()
        } else {
            resp.Result = &pb.CommandResponse_Message{Message: result}
        }

    case pb.CommandType_COMMAND_LAUNCH_APP:
        result, err := handleLaunchApp(cmd.GetLaunchApp())  // existing handler
        if err != nil {
            resp.Success = false
            resp.Error = err.Error()
        } else {
            resp.Result = &pb.CommandResponse_Message{Message: result}
        }

    case pb.CommandType_COMMAND_INFO:
        result, err := handleInfo()                         // existing handler
        if err != nil {
            resp.Success = false
            resp.Error = err.Error()
        } else {
            resp.Result = &pb.CommandResponse_Device{Device: result}
        }

    case pb.CommandType_COMMAND_INFO_V2:
        result, err := handleInfoV2()                       // existing handler
        if err != nil {
            resp.Success = false
            resp.Error = err.Error()
        } else {
            resp.Result = &pb.CommandResponse_DeviceV2{DeviceV2: result}
        }

    case pb.CommandType_COMMAND_LIST_PACKAGES:
        result, err := handleListPackages()                 // existing handler
        if err != nil {
            resp.Success = false
            resp.Error = err.Error()
        } else {
            resp.Result = &pb.CommandResponse_Apps{Apps: result}
        }

    // ... same pattern for all other commands:
    // COMMAND_CLEAR_PROXY, COMMAND_HEALTH_CHECK, COMMAND_REMOVE_PACKAGE,
    // COMMAND_INSTALL_PACKAGE, COMMAND_CONFIGURE_PROXY, COMMAND_CONFIGURE_GADGET,
    // COMMAND_START_AUTO_PILOT, COMMAND_CLEAN

    default:
        resp.Success = false
        resp.Error = fmt.Sprintf("unknown command type: %v", cmd.GetType())
    }

    return resp
}
```

**Key point**: The `handleEcho()`, `handleLaunchApp()`, etc. are Mary's **existing handler functions** — the same logic that currently runs inside the `MoriartyServer` implementation. The only new code is the switch-case that unwraps the `Command` envelope and wraps the result into a `CommandResponse`.

### Command Type to Handler & Response Type Mapping

| CommandType | Payload field | Handler returns | Response field |
|---|---|---|---|
| `COMMAND_ECHO` | `cmd.GetEcho()` → `*Message` | `*Message, error` | `CommandResponse_Message` |
| `COMMAND_LAUNCH_APP` | `cmd.GetLaunchApp()` → `*App` | `*Message, error` | `CommandResponse_Message` |
| `COMMAND_CLEAR_PROXY` | `cmd.GetClearProxy()` → `*Empty` | `*Message, error` | `CommandResponse_Message` |
| `COMMAND_HEALTH_CHECK` | `cmd.GetHealthCheck()` → `*Empty` | `*Message, error` | `CommandResponse_Message` |
| `COMMAND_REMOVE_PACKAGE` | `cmd.GetRemovePackage()` → `*App` | `*Message, error` | `CommandResponse_Message` |
| `COMMAND_INSTALL_PACKAGE` | `cmd.GetInstallPackage()` → `*InstallReq` | `*Message, error` | `CommandResponse_Message` |
| `COMMAND_CONFIGURE_PROXY` | `cmd.GetConfigureProxy()` → `*ConfigProxyReq` | `*Message, error` | `CommandResponse_Message` |
| `COMMAND_CONFIGURE_GADGET` | `cmd.GetConfigureGadget()` → `*App` | `*Message, error` | `CommandResponse_Message` |
| `COMMAND_START_AUTO_PILOT` | `cmd.GetStartAutoPilot()` → `*AutoPilotConfig` | `*Message, error` | `CommandResponse_Message` |
| `COMMAND_CLEAN` | `cmd.GetClean()` → `*CleanOptions` | `*Message, error` | `CommandResponse_Message` |
| `COMMAND_INFO` | `cmd.GetInfo()` → `*Empty` | `*Device, error` | `CommandResponse_Device` |
| `COMMAND_INFO_V2` | `cmd.GetInfoV2()` → `*Empty` | `*DeviceV2, error` | `CommandResponse_DeviceV2` |
| `COMMAND_LIST_PACKAGES` | `cmd.GetListPackages()` → `*Empty` | `*Apps, error` | `CommandResponse_Apps` |

### Reconnection Strategy

```go
func backoffSleep() {
    // Start at 1s, double each time, cap at 60s, add jitter
    wait := min(wait * 2, 60 * time.Second)
    jitter := time.Duration(rand.Int63n(int64(wait) / 5))
    time.Sleep(wait + jitter)
}
```

### What Runs When

Mary Agent runs in **one mode or the other**, controlled by configuration — never both simultaneously.

```go
func main() {
    if config.ConnectionMode == "relay" {
        // Relay mode: connect outbound to the relay server
        // Mary acts as a gRPC CLIENT
        connectToRelay(config.RelayAddr, deviceInfo, config.AuthToken)
    } else {
        // Direct mode (existing): listen for incoming connections
        // Mary acts as a gRPC SERVER
        startMoriartyServer(listenPort)
    }
}
```

The mode is determined at startup by a config flag (e.g., `connection_mode: "direct" | "relay"`). This is a per-device decision — some devices may be configured for direct, others for relay, depending on their network environment.

---

## What Moriarty (Python/Django) Needs to Do

Moriarty currently uses `MoriartyStub` to call Mary Agent directly. **This stays exactly as-is for direct-mode devices.**

### New Import

```python
from bakerstreet import bakerstreet_bridge_pb2 as bridge_pb2
from bakerstreet import bakerstreet_bridge_pb2_grpc as bridge_grpc
```

### 1. Implement the Relay Server

The relay server is a gRPC server running inside Moriarty (can run on a separate port or as a separate process alongside Django). It implements two servicers:

- `MoriartyBridgeServicer` — accepts device connections via bidi stream
- `MoriartyDispatchServicer` — accepts platform dispatch calls

See the [Relay Server Internals](#relay-server-internals-inside-moriarty) section above for the implementation pattern.

### 2. Route Commands by Connection Mode

The platform code needs a routing layer that picks direct vs relay per device:

```python
def execute_on_device(device, rpc_name, request):
    if device.connection_mode == "relay":
        # Use MoriartyDispatch — command goes through the relay
        #
        # Option A: call via gRPC stub (if relay runs on a separate port)
        stub = bridge_grpc.MoriartyDispatchStub(relay_channel)
        targeted_request = wrap_targeted(device.device_id, request)
        return getattr(stub, rpc_name)(targeted_request)

        # Option B: call relay internals directly (if relay is in-process)
        return relay.dispatch(device.device_id, rpc_name, request)

    else:
        # Existing direct mode — unchanged
        channel = grpc.insecure_channel(f"{device.ip}:{device.port}")
        stub = bakerstreet_pb2_grpc.MoriartyStub(channel)
        return getattr(stub, rpc_name)(request)
```

### 3. Use ListDevices for Discovery

```python
# Find all devices currently connected via relay
stub = bridge_grpc.MoriartyDispatchStub(relay_channel)
device_list = stub.ListDevices(bakerstreet_pb2.Empty())
for device in device_list.devices:
    print(f"{device.device_id} ({device.custom_identifier}) - connected at {device.connected_at}")
```

---

## What Stays the Same

| Component | Changes? |
|-----------|----------|
| `bakerstreet.proto` | **No** |
| `bakerstreet.pb.go` / `bakerstreet_pb2.py` | **No** |
| `bakerstreet_grpc.pb.go` / `bakerstreet_pb2_grpc.py` | **No** |
| Mary's `MoriartyServer` implementation | **No** |
| Moriarty's `MoriartyStub` usage for direct-mode devices | **No** |
| Mary's existing handler logic (echo, launch, install, etc.) | **No** — reused by both modes |
| All existing message types (Message, App, Device, etc.) | **No** — reused by bridge proto |

---

## Importing the New Stubs

### Go (Mary Agent)

```go
import (
    pb "github.com/appknox/bakerstreet/v2/bakerstreet"
)

// All bridge types are in the same package:
// pb.MoriartyBridgeClient, pb.DeviceMessage, pb.ServerMessage,
// pb.Command, pb.CommandResponse, pb.CommandType_COMMAND_ECHO, etc.
```

### Python (Moriarty)

```python
# Existing (unchanged)
from bakerstreet import bakerstreet_pb2
from bakerstreet import bakerstreet_pb2_grpc

# New (additive)
from bakerstreet import bakerstreet_bridge_pb2
from bakerstreet import bakerstreet_bridge_pb2_grpc
```

Update `bakerstreet` package version to `>= 2.3.0` in your requirements.

---

## Keepalive Configuration

### Mary Agent (gRPC client to relay)

```go
grpc.WithKeepaliveParams(keepalive.ClientParameters{
    Time:    30 * time.Second,  // send ping every 30s
    Timeout: 10 * time.Second,  // wait 10s for pong before considering dead
})
```

### Relay Server (gRPC server in Moriarty)

```python
server = grpc.server(
    futures.ThreadPoolExecutor(max_workers=10),
    options=[
        ('grpc.max_connection_idle_ms', 60000),               # close idle connections after 60s
        ('grpc.keepalive_time_ms', 30000),                     # send keepalive ping every 30s
        ('grpc.keepalive_timeout_ms', 10000),                  # wait 10s for pong
        ('grpc.http2.min_ping_interval_without_data_ms', 20000), # min 20s between client pings
        ('grpc.keepalive_permit_without_calls', True),         # allow pings without active RPCs
    ],
)
```

---

## Summary of New Work

| Team | What to build | Effort |
|------|--------------|--------|
| **Mary Agent** | Bridge client mode: connect to relay, registration handshake, command dispatch loop (switch on CommandType → call existing handlers), reconnection with backoff. Runs in **either** direct or relay mode per config — not both. | Moderate — new transport layer, but reuses all existing handler logic |
| **Moriarty** | Relay server: MoriartyBridgeServicer (accept device streams, manage registry), MoriartyDispatchServicer (accept platform calls, route to device streams), device registry, pending request map, routing layer (direct vs relay per device) | Larger — new gRPC server component within Moriarty |

---

## Implementation Checklist: Mary Agent (Go)

> Copy this section into the Mary Agent repo as context when implementing.

### Prerequisites
- Update `bakerstreet` dependency to `>= v2.3.0` (branch `feature/cyod-pase-1`)
- `go get github.com/appknox/bakerstreet/v2@feature/cyod-pase-1`

### New Config Fields
```
connection_mode  string  // "direct" (default) or "relay"
relay_addr       string  // e.g. "relay.appknox.com:443"
auth_token       string  // pre-shared token for relay authentication
```

### Files to Create/Modify

1. **Config** — add `connection_mode`, `relay_addr`, `auth_token` fields to existing config struct/file
2. **Bridge client** (new file, e.g. `bridge.go` or `relay_client.go`):
   - `connectToRelay(relayAddr, deviceInfo, authToken)` — outer reconnect loop with exponential backoff
   - `runBridgeSession(relayAddr, deviceInfo, authToken) error` — single session lifecycle:
     1. `grpc.Dial()` with TLS + keepalive
     2. `pb.NewMoriartyBridgeClient(conn).Connect(ctx)` — open bidi stream
     3. Send `DeviceMessage{Registration: DeviceRegistration{...}}`
     4. Recv `ServerMessage`, check `GetAck().GetAccepted()`
     5. Loop: `Recv()` → `GetCommand()` → `executeCommand()` → `Send(DeviceMessage{Response: ...})`
   - `executeCommand(cmd *pb.Command) *pb.CommandResponse` — switch on `cmd.GetType()`:
     - Call the **same handler functions** already used by `MoriartyServer`
     - Wrap result in `CommandResponse` with matching `request_id`
     - On error: set `success=false`, `error=err.Error()`
3. **Main / startup** — change to either/or based on `connection_mode`:
   ```go
   if config.ConnectionMode == "relay" {
       connectToRelay(config.RelayAddr, deviceInfo, config.AuthToken)
   } else {
       startMoriartyServer(listenPort)  // existing code, unchanged
   }
   ```

### Rules
- **Do NOT modify any existing handler functions** — they are reused as-is by both modes
- **Do NOT modify the existing `MoriartyServer` implementation** — direct mode stays untouched
- **Do NOT run both modes simultaneously** — it's one or the other per config
- All bridge types (`DeviceMessage`, `ServerMessage`, `Command`, `CommandResponse`, `CommandType_*`) are in the same `pb` package already imported
- The `executeCommand` switch must cover all 13 `CommandType` values (see mapping table above)

---

## Implementation Checklist: Moriarty (Python/Django)

> Copy this section into the Moriarty repo as context when implementing.

### Prerequisites
- Update `bakerstreet` dependency to `>= 2.3.0`
- New imports:
  ```python
  from bakerstreet import bakerstreet_bridge_pb2 as bridge_pb2
  from bakerstreet import bakerstreet_bridge_pb2_grpc as bridge_grpc
  from bakerstreet import bakerstreet_pb2
  ```

### Components to Build

#### 1. Device Registry (thread-safe)
- In-memory `dict[str, DeviceConnection]` guarded by `threading.Lock`
- `DeviceConnection` holds: stream context, `DeviceRegistration` metadata, `connected_at` timestamp, per-device `threading.Lock` for serializing sends
- Methods: `add(device_id, conn)`, `remove(device_id)`, `get(device_id) → DeviceConnection | None`, `list() → list[ConnectedDevice]`

#### 2. Pending Requests (thread-safe)
- In-memory `dict[str, concurrent.futures.Future]` guarded by `threading.Lock`
- Methods: `create(request_id) → Future`, `resolve(request_id, response)`, `remove(request_id)`
- Future timeout: 30 seconds (configurable)

#### 3. MoriartyBridgeServicer (device-facing)
- Implements `bridge_grpc.MoriartyBridgeServicer`
- `Connect(request_iterator, context)`:
  - Recv first message → must be `DeviceRegistration`
  - Validate `auth_token`
  - Send `RegistrationAck`
  - Add to registry
  - Loop recv: incoming `CommandResponse` → resolve pending request by `request_id`
  - On disconnect: remove from registry (in `finally`)
- **Bidi stream challenge**: Python gRPC bidi streams use generator-based pattern. Sending commands mid-stream requires a thread-safe queue that the generator drains.

#### 4. MoriartyDispatchServicer (platform-facing)
- Implements `bridge_grpc.MoriartyDispatchServicer`
- All 13 dispatch RPCs follow the same pattern:
  1. Extract `device_id` from targeted request
  2. `registry.get(device_id)` → abort `UNAVAILABLE` if not found
  3. Generate `request_id` (UUID)
  4. `pending.create(request_id)` → Future
  5. Build `Command(request_id, type, payload)` → send via device stream
  6. `future.result(timeout=30)` → abort `DEADLINE_EXCEEDED` on timeout
  7. Check `response.success` → abort `INTERNAL` if false
  8. Return the typed result from `response`
- `ListDevices(Empty)` → return `DeviceList` from registry (no device routing)

#### 5. gRPC Server Setup
- Run gRPC server on a dedicated port (separate from Django's HTTP port)
- Register both servicers on same server
- Configure keepalive options
- Start as a background thread/process alongside Django

#### 6. Routing Layer
- Modify existing code that calls `MoriartyStub` to check `device.connection_mode`
- If `"relay"`: use `MoriartyDispatchStub` (or call relay internals directly if in-process)
- If `"direct"`: use existing `MoriartyStub` to device IP — **no changes**

#### 7. Device Model
- Add `connection_mode` field: `"direct"` (default) or `"relay"`
- Add `device_id` field (UUID) — used for relay routing (may already exist)

### Rules
- **Do NOT modify any existing direct-mode code paths** — they stay untouched
- **Do NOT modify bakerstreet proto files** — they are already done
- The relay gRPC server runs alongside Django, not inside it (gRPC needs its own port)
- All 13 dispatch methods are near-identical — use a helper function to avoid repetition
- Device registry and pending requests must be thread-safe (multiple gRPC threads access them)
- Always clean up pending requests in a `finally` block to prevent memory leaks
