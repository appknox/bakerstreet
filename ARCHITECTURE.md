# Baker Street - Architecture & Overview

> "The place where most characters meet." — RPC definitions for Appknox services.

## What is Baker Street?

Baker Street is a **Protocol Buffer (protobuf) and gRPC definition library** that serves as the central communication contract between services in the [Appknox](https://www.appknox.com/) ecosystem. It defines the messages and RPC methods that allow Appknox's central services to communicate with remote testing devices and agents.

The project generates client/server stubs for both **Python** and **Go**, making it a shared dependency consumed by services written in either language.

## Tech Stack

| Component        | Technology                  |
| ---------------- | --------------------------- |
| IDL              | Protocol Buffers (proto3)   |
| RPC Framework    | gRPC                        |
| Go version       | 1.22                        |
| Python version   | >= 3.12                     |
| Python packaging | Hatch + UV                  |
| Go module        | `github.com/appknox/bakerstreet/v2` |

## Project Structure

```
bakerstreet/
├── bakerstreet/
│   ├── bakerstreet.proto          # Source of truth — all message & service definitions
│   ├── bakerstreet.pb.go          # Generated Go protobuf code
│   ├── bakerstreet_grpc.pb.go     # Generated Go gRPC stubs
│   ├── bakerstreet_pb2.py         # Generated Python protobuf code
│   ├── bakerstreet_pb2_grpc.py    # Generated Python gRPC stubs
│   └── __init__.py                # Python package marker
├── scripts/
│   ├── bakerstreet.sh             # Proto compilation script
│   └── deploy.sh                  # Publishing script
├── go.mod / go.sum                # Go module & dependency lock
├── pyproject.toml                 # Python package config (Hatch backend)
├── hatch_build.py                 # Custom build hook — compiles .proto during build
├── uv.lock                        # Python dependency lock (UV)
└── README.md
```

### Single source of truth

Everything originates from **`bakerstreet/bakerstreet.proto`**. The `.pb.go`, `_pb2.py`, and `_grpc` files are all **auto-generated** — they should never be edited by hand.

## The Moriarty gRPC Service

The proto file defines a single gRPC service called **Moriarty** with 12 RPC endpoints. These can be grouped into four functional areas:

### Device Information

| RPC            | Request   | Response   | Purpose                                   |
| -------------- | --------- | ---------- | ----------------------------------------- |
| `Info`         | `Empty`   | `Device`   | Get basic device info (UUID, platform, CPU, model) |
| `InfoV2`       | `Empty`   | `DeviceV2` | Get extended device info (capabilities, agent versions, VPN, screen lock, etc.) |
| `HealthCheck`  | `Empty`   | `Message`  | Check if the device agent is alive        |
| `Echo`         | `Message` | `Message`  | Simple echo for connectivity testing      |

### App Lifecycle

| RPC              | Request      | Response | Purpose                        |
| ---------------- | ------------ | -------- | ------------------------------ |
| `InstallPackage` | `InstallReq` | `Message` | Install an app from a URL     |
| `RemovePackage`  | `App`        | `Message` | Uninstall an app by name      |
| `LaunchApp`      | `App`        | `Message` | Launch an installed app        |
| `ListPackages`   | `Empty`      | `Apps`   | List all installed packages    |

### Proxy & Network Configuration

| RPC              | Request          | Response | Purpose                                  |
| ---------------- | ---------------- | -------- | ---------------------------------------- |
| `ConfigureProxy` | `ConfigProxyReq` | `Message` | Set up proxy (IP, port, hosts, chaining) |
| `ClearProxy`     | `Empty`          | `Message` | Remove proxy configuration               |

### Automation

| RPC               | Request           | Response | Purpose                                    |
| ----------------- | ----------------- | -------- | ------------------------------------------ |
| `StartAutoPilot`  | `AutoPilotConfig` | `Message` | Start automated security testing           |
| `ConfigureGadget` | `App`             | `Message` | Configure Frida gadget for instrumentation |
| `Clean`           | `CleanOptions`    | `Message` | Clean up device state after testing        |

## Key Message Types

### `Device` — Basic device metadata
Fields: `Uuid`, `IsTablet`, `Platform`, `PlatformVersion`, `CpuArchitecture`, `Model`, `Name`

### `DeviceV2` — Extended device capabilities
Adds fields for automation capabilities: `DeviceIdentifier`, `CanRunManual`, `CanRunAutomated`, `HasAutomationEngine`, `HasVnc`, `SimNetwork`, `PhoneNumber`, `ScreenLock`, `VpnPackageName`, `PersistentApps`, `ExtraCapabilities`, `AgentVersion`, `FridaVersion`, `FridaScriptVersion`, `AutoPilotVersion`, `CustomIdentifier`

### `AutoPilotConfig` — Automation settings
Fields: `DryRun` (bool), `AutoPilotPreference` (string), `App` (string)

### `ConfigProxyReq` — Proxy configuration
Fields: `Ip`, `Port`, `Hosts`, `ChainHost`, `ChainPort`, `CaptureAll`

## How It Fits Together

```
┌─────────────────────┐         gRPC (Moriarty)        ┌──────────────────────┐
│                     │ ──────────────────────────────► │                      │
│  Appknox Platform   │   InstallPackage, LaunchApp,   │   Device Agent       │
│  (Python / Go)      │   ConfigureProxy, StartAuto..  │   (on test device)   │
│                     │ ◄────────────────────────────── │                      │
│  Uses: MoriartyStub │   Device, DeviceV2, Message    │  Implements: Moriarty│
└─────────────────────┘                                 └──────────────────────┘
```

- The **Appknox platform** acts as the gRPC **client**, calling methods on remote devices.
- **Device agents** act as gRPC **servers**, implementing the Moriarty service interface.
- Baker Street is the **shared contract** — both sides depend on its generated code.

## Build & Development

### Prerequisites
- `protoc` (Protocol Buffer compiler)
- `protoc-gen-go` and `protoc-gen-go-grpc` (Go plugins)
- `uv` (Python package manager)
- `grpcio-tools` (Python gRPC code generation)

### Building
```bash
uv build
```
The custom Hatch build hook (`hatch_build.py`) automatically:
1. Removes old generated files
2. Compiles `bakerstreet.proto` → Go stubs via `protoc`
3. Compiles `bakerstreet.proto` → Python stubs via `grpc_tools.protoc`

### Publishing
```bash
uv publish
```
Publishes the Python package to PyPI as `bakerstreet`.

### Consuming in Go
```go
import bakerstreet "github.com/appknox/bakerstreet/v2/bakerstreet"
```

### Consuming in Python
```python
from bakerstreet import bakerstreet_pb2, bakerstreet_pb2_grpc
```

## Versioning

Current version: **2.3.0**

The project uses semantic versioning. The Go module path includes `/v2` for major version compatibility. Version is defined in `pyproject.toml`.
