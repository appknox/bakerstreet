import os
import subprocess
from grpc_tools import protoc
from hatchling.builders.hooks.plugin.interface import BuildHookInterface

class CustomBuildHook(BuildHookInterface):
    def initialize(self, version, build_data):
        compile_proto()

def compile_proto():
    bakerstreet = "bakerstreet"
    proto_file = f"{bakerstreet}/{bakerstreet}.proto"
    bridge_proto_file = f"{bakerstreet}/{bakerstreet}_bridge.proto"

    # Remove existing generated Python files
    py_files = [f for f in os.listdir(bakerstreet) if f.endswith('.py') and f.startswith('bakerstreet')]
    for file in py_files:
        os.remove(os.path.join(bakerstreet, file))

    # Compile Go protobuf and gRPC files (best-effort — skipped if protoc-gen-go is not in PATH)
    for pf in [proto_file, bridge_proto_file]:
        go_compile_cmd = [
            "protoc",
            "-I", ".",
            "--go_out=.",
            "--go_opt=paths=source_relative",
            "--go-grpc_out=.",
            "--go-grpc_opt=paths=source_relative",
            pf
        ]
        result = subprocess.run(go_compile_cmd, capture_output=True)
        if result.returncode != 0:
            print(f"Warning: Go protoc compilation skipped for {pf} (protoc-gen-go not found)")
            break

    # Compile Python protobuf and gRPC files
    for pf in [proto_file, bridge_proto_file]:
        protoc.main(
            ("-I  .", "--python_out=.", "--grpc_python_out=.", "{:s}".format(pf))
        )