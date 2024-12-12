# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
"""Client and server classes corresponding to protobuf-defined services."""
import grpc
import warnings

from bakerstreet import bakerstreet_pb2 as bakerstreet_dot_bakerstreet__pb2

GRPC_GENERATED_VERSION = '1.68.1'
GRPC_VERSION = grpc.__version__
_version_not_supported = False

try:
    from grpc._utilities import first_version_is_lower
    _version_not_supported = first_version_is_lower(GRPC_VERSION, GRPC_GENERATED_VERSION)
except ImportError:
    _version_not_supported = True

if _version_not_supported:
    raise RuntimeError(
        f'The grpc package installed is at version {GRPC_VERSION},'
        + f' but the generated code in bakerstreet/bakerstreet_pb2_grpc.py depends on'
        + f' grpcio>={GRPC_GENERATED_VERSION}.'
        + f' Please upgrade your grpc module to grpcio>={GRPC_GENERATED_VERSION}'
        + f' or downgrade your generated code using grpcio-tools<={GRPC_VERSION}.'
    )


class MoriartyStub(object):
    """Missing associated documentation comment in .proto file."""

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.Echo = channel.unary_unary(
                '/com.appknox.bakerstreet.Moriarty/Echo',
                request_serializer=bakerstreet_dot_bakerstreet__pb2.Message.SerializeToString,
                response_deserializer=bakerstreet_dot_bakerstreet__pb2.Message.FromString,
                _registered_method=True)
        self.LaunchApp = channel.unary_unary(
                '/com.appknox.bakerstreet.Moriarty/LaunchApp',
                request_serializer=bakerstreet_dot_bakerstreet__pb2.App.SerializeToString,
                response_deserializer=bakerstreet_dot_bakerstreet__pb2.Message.FromString,
                _registered_method=True)
        self.ClearProxy = channel.unary_unary(
                '/com.appknox.bakerstreet.Moriarty/ClearProxy',
                request_serializer=bakerstreet_dot_bakerstreet__pb2.Empty.SerializeToString,
                response_deserializer=bakerstreet_dot_bakerstreet__pb2.Message.FromString,
                _registered_method=True)
        self.HealthCheck = channel.unary_unary(
                '/com.appknox.bakerstreet.Moriarty/HealthCheck',
                request_serializer=bakerstreet_dot_bakerstreet__pb2.Empty.SerializeToString,
                response_deserializer=bakerstreet_dot_bakerstreet__pb2.Message.FromString,
                _registered_method=True)
        self.RemovePackage = channel.unary_unary(
                '/com.appknox.bakerstreet.Moriarty/RemovePackage',
                request_serializer=bakerstreet_dot_bakerstreet__pb2.App.SerializeToString,
                response_deserializer=bakerstreet_dot_bakerstreet__pb2.Message.FromString,
                _registered_method=True)
        self.InstallPackage = channel.unary_unary(
                '/com.appknox.bakerstreet.Moriarty/InstallPackage',
                request_serializer=bakerstreet_dot_bakerstreet__pb2.InstallReq.SerializeToString,
                response_deserializer=bakerstreet_dot_bakerstreet__pb2.Message.FromString,
                _registered_method=True)
        self.ConfigureProxy = channel.unary_unary(
                '/com.appknox.bakerstreet.Moriarty/ConfigureProxy',
                request_serializer=bakerstreet_dot_bakerstreet__pb2.ConfigProxyReq.SerializeToString,
                response_deserializer=bakerstreet_dot_bakerstreet__pb2.Message.FromString,
                _registered_method=True)
        self.ConfigureGadget = channel.unary_unary(
                '/com.appknox.bakerstreet.Moriarty/ConfigureGadget',
                request_serializer=bakerstreet_dot_bakerstreet__pb2.App.SerializeToString,
                response_deserializer=bakerstreet_dot_bakerstreet__pb2.Message.FromString,
                _registered_method=True)
        self.ConfigureAutopilot = channel.unary_unary(
                '/com.appknox.bakerstreet.Moriarty/ConfigureAutopilot',
                request_serializer=bakerstreet_dot_bakerstreet__pb2.AutoPilotConfig.SerializeToString,
                response_deserializer=bakerstreet_dot_bakerstreet__pb2.Message.FromString,
                _registered_method=True)
        self.Clean = channel.unary_unary(
                '/com.appknox.bakerstreet.Moriarty/Clean',
                request_serializer=bakerstreet_dot_bakerstreet__pb2.CleanOptions.SerializeToString,
                response_deserializer=bakerstreet_dot_bakerstreet__pb2.Message.FromString,
                _registered_method=True)
        self.Info = channel.unary_unary(
                '/com.appknox.bakerstreet.Moriarty/Info',
                request_serializer=bakerstreet_dot_bakerstreet__pb2.Empty.SerializeToString,
                response_deserializer=bakerstreet_dot_bakerstreet__pb2.Device.FromString,
                _registered_method=True)
        self.InfoV2 = channel.unary_unary(
                '/com.appknox.bakerstreet.Moriarty/InfoV2',
                request_serializer=bakerstreet_dot_bakerstreet__pb2.Empty.SerializeToString,
                response_deserializer=bakerstreet_dot_bakerstreet__pb2.DeviceV2.FromString,
                _registered_method=True)
        self.ListPackages = channel.unary_unary(
                '/com.appknox.bakerstreet.Moriarty/ListPackages',
                request_serializer=bakerstreet_dot_bakerstreet__pb2.Empty.SerializeToString,
                response_deserializer=bakerstreet_dot_bakerstreet__pb2.Apps.FromString,
                _registered_method=True)


class MoriartyServicer(object):
    """Missing associated documentation comment in .proto file."""

    def Echo(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def LaunchApp(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def ClearProxy(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def HealthCheck(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def RemovePackage(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def InstallPackage(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def ConfigureProxy(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def ConfigureGadget(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def ConfigureAutopilot(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def Clean(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def Info(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def InfoV2(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def ListPackages(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')


def add_MoriartyServicer_to_server(servicer, server):
    rpc_method_handlers = {
            'Echo': grpc.unary_unary_rpc_method_handler(
                    servicer.Echo,
                    request_deserializer=bakerstreet_dot_bakerstreet__pb2.Message.FromString,
                    response_serializer=bakerstreet_dot_bakerstreet__pb2.Message.SerializeToString,
            ),
            'LaunchApp': grpc.unary_unary_rpc_method_handler(
                    servicer.LaunchApp,
                    request_deserializer=bakerstreet_dot_bakerstreet__pb2.App.FromString,
                    response_serializer=bakerstreet_dot_bakerstreet__pb2.Message.SerializeToString,
            ),
            'ClearProxy': grpc.unary_unary_rpc_method_handler(
                    servicer.ClearProxy,
                    request_deserializer=bakerstreet_dot_bakerstreet__pb2.Empty.FromString,
                    response_serializer=bakerstreet_dot_bakerstreet__pb2.Message.SerializeToString,
            ),
            'HealthCheck': grpc.unary_unary_rpc_method_handler(
                    servicer.HealthCheck,
                    request_deserializer=bakerstreet_dot_bakerstreet__pb2.Empty.FromString,
                    response_serializer=bakerstreet_dot_bakerstreet__pb2.Message.SerializeToString,
            ),
            'RemovePackage': grpc.unary_unary_rpc_method_handler(
                    servicer.RemovePackage,
                    request_deserializer=bakerstreet_dot_bakerstreet__pb2.App.FromString,
                    response_serializer=bakerstreet_dot_bakerstreet__pb2.Message.SerializeToString,
            ),
            'InstallPackage': grpc.unary_unary_rpc_method_handler(
                    servicer.InstallPackage,
                    request_deserializer=bakerstreet_dot_bakerstreet__pb2.InstallReq.FromString,
                    response_serializer=bakerstreet_dot_bakerstreet__pb2.Message.SerializeToString,
            ),
            'ConfigureProxy': grpc.unary_unary_rpc_method_handler(
                    servicer.ConfigureProxy,
                    request_deserializer=bakerstreet_dot_bakerstreet__pb2.ConfigProxyReq.FromString,
                    response_serializer=bakerstreet_dot_bakerstreet__pb2.Message.SerializeToString,
            ),
            'ConfigureGadget': grpc.unary_unary_rpc_method_handler(
                    servicer.ConfigureGadget,
                    request_deserializer=bakerstreet_dot_bakerstreet__pb2.App.FromString,
                    response_serializer=bakerstreet_dot_bakerstreet__pb2.Message.SerializeToString,
            ),
            'ConfigureAutopilot': grpc.unary_unary_rpc_method_handler(
                    servicer.ConfigureAutopilot,
                    request_deserializer=bakerstreet_dot_bakerstreet__pb2.AutoPilotConfig.FromString,
                    response_serializer=bakerstreet_dot_bakerstreet__pb2.Message.SerializeToString,
            ),
            'Clean': grpc.unary_unary_rpc_method_handler(
                    servicer.Clean,
                    request_deserializer=bakerstreet_dot_bakerstreet__pb2.CleanOptions.FromString,
                    response_serializer=bakerstreet_dot_bakerstreet__pb2.Message.SerializeToString,
            ),
            'Info': grpc.unary_unary_rpc_method_handler(
                    servicer.Info,
                    request_deserializer=bakerstreet_dot_bakerstreet__pb2.Empty.FromString,
                    response_serializer=bakerstreet_dot_bakerstreet__pb2.Device.SerializeToString,
            ),
            'InfoV2': grpc.unary_unary_rpc_method_handler(
                    servicer.InfoV2,
                    request_deserializer=bakerstreet_dot_bakerstreet__pb2.Empty.FromString,
                    response_serializer=bakerstreet_dot_bakerstreet__pb2.DeviceV2.SerializeToString,
            ),
            'ListPackages': grpc.unary_unary_rpc_method_handler(
                    servicer.ListPackages,
                    request_deserializer=bakerstreet_dot_bakerstreet__pb2.Empty.FromString,
                    response_serializer=bakerstreet_dot_bakerstreet__pb2.Apps.SerializeToString,
            ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
            'com.appknox.bakerstreet.Moriarty', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))
    server.add_registered_method_handlers('com.appknox.bakerstreet.Moriarty', rpc_method_handlers)


 # This class is part of an EXPERIMENTAL API.
class Moriarty(object):
    """Missing associated documentation comment in .proto file."""

    @staticmethod
    def Echo(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(
            request,
            target,
            '/com.appknox.bakerstreet.Moriarty/Echo',
            bakerstreet_dot_bakerstreet__pb2.Message.SerializeToString,
            bakerstreet_dot_bakerstreet__pb2.Message.FromString,
            options,
            channel_credentials,
            insecure,
            call_credentials,
            compression,
            wait_for_ready,
            timeout,
            metadata,
            _registered_method=True)

    @staticmethod
    def LaunchApp(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(
            request,
            target,
            '/com.appknox.bakerstreet.Moriarty/LaunchApp',
            bakerstreet_dot_bakerstreet__pb2.App.SerializeToString,
            bakerstreet_dot_bakerstreet__pb2.Message.FromString,
            options,
            channel_credentials,
            insecure,
            call_credentials,
            compression,
            wait_for_ready,
            timeout,
            metadata,
            _registered_method=True)

    @staticmethod
    def ClearProxy(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(
            request,
            target,
            '/com.appknox.bakerstreet.Moriarty/ClearProxy',
            bakerstreet_dot_bakerstreet__pb2.Empty.SerializeToString,
            bakerstreet_dot_bakerstreet__pb2.Message.FromString,
            options,
            channel_credentials,
            insecure,
            call_credentials,
            compression,
            wait_for_ready,
            timeout,
            metadata,
            _registered_method=True)

    @staticmethod
    def HealthCheck(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(
            request,
            target,
            '/com.appknox.bakerstreet.Moriarty/HealthCheck',
            bakerstreet_dot_bakerstreet__pb2.Empty.SerializeToString,
            bakerstreet_dot_bakerstreet__pb2.Message.FromString,
            options,
            channel_credentials,
            insecure,
            call_credentials,
            compression,
            wait_for_ready,
            timeout,
            metadata,
            _registered_method=True)

    @staticmethod
    def RemovePackage(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(
            request,
            target,
            '/com.appknox.bakerstreet.Moriarty/RemovePackage',
            bakerstreet_dot_bakerstreet__pb2.App.SerializeToString,
            bakerstreet_dot_bakerstreet__pb2.Message.FromString,
            options,
            channel_credentials,
            insecure,
            call_credentials,
            compression,
            wait_for_ready,
            timeout,
            metadata,
            _registered_method=True)

    @staticmethod
    def InstallPackage(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(
            request,
            target,
            '/com.appknox.bakerstreet.Moriarty/InstallPackage',
            bakerstreet_dot_bakerstreet__pb2.InstallReq.SerializeToString,
            bakerstreet_dot_bakerstreet__pb2.Message.FromString,
            options,
            channel_credentials,
            insecure,
            call_credentials,
            compression,
            wait_for_ready,
            timeout,
            metadata,
            _registered_method=True)

    @staticmethod
    def ConfigureProxy(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(
            request,
            target,
            '/com.appknox.bakerstreet.Moriarty/ConfigureProxy',
            bakerstreet_dot_bakerstreet__pb2.ConfigProxyReq.SerializeToString,
            bakerstreet_dot_bakerstreet__pb2.Message.FromString,
            options,
            channel_credentials,
            insecure,
            call_credentials,
            compression,
            wait_for_ready,
            timeout,
            metadata,
            _registered_method=True)

    @staticmethod
    def ConfigureGadget(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(
            request,
            target,
            '/com.appknox.bakerstreet.Moriarty/ConfigureGadget',
            bakerstreet_dot_bakerstreet__pb2.App.SerializeToString,
            bakerstreet_dot_bakerstreet__pb2.Message.FromString,
            options,
            channel_credentials,
            insecure,
            call_credentials,
            compression,
            wait_for_ready,
            timeout,
            metadata,
            _registered_method=True)

    @staticmethod
    def ConfigureAutopilot(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(
            request,
            target,
            '/com.appknox.bakerstreet.Moriarty/ConfigureAutopilot',
            bakerstreet_dot_bakerstreet__pb2.AutoPilotConfig.SerializeToString,
            bakerstreet_dot_bakerstreet__pb2.Message.FromString,
            options,
            channel_credentials,
            insecure,
            call_credentials,
            compression,
            wait_for_ready,
            timeout,
            metadata,
            _registered_method=True)

    @staticmethod
    def Clean(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(
            request,
            target,
            '/com.appknox.bakerstreet.Moriarty/Clean',
            bakerstreet_dot_bakerstreet__pb2.CleanOptions.SerializeToString,
            bakerstreet_dot_bakerstreet__pb2.Message.FromString,
            options,
            channel_credentials,
            insecure,
            call_credentials,
            compression,
            wait_for_ready,
            timeout,
            metadata,
            _registered_method=True)

    @staticmethod
    def Info(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(
            request,
            target,
            '/com.appknox.bakerstreet.Moriarty/Info',
            bakerstreet_dot_bakerstreet__pb2.Empty.SerializeToString,
            bakerstreet_dot_bakerstreet__pb2.Device.FromString,
            options,
            channel_credentials,
            insecure,
            call_credentials,
            compression,
            wait_for_ready,
            timeout,
            metadata,
            _registered_method=True)

    @staticmethod
    def InfoV2(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(
            request,
            target,
            '/com.appknox.bakerstreet.Moriarty/InfoV2',
            bakerstreet_dot_bakerstreet__pb2.Empty.SerializeToString,
            bakerstreet_dot_bakerstreet__pb2.DeviceV2.FromString,
            options,
            channel_credentials,
            insecure,
            call_credentials,
            compression,
            wait_for_ready,
            timeout,
            metadata,
            _registered_method=True)

    @staticmethod
    def ListPackages(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(
            request,
            target,
            '/com.appknox.bakerstreet.Moriarty/ListPackages',
            bakerstreet_dot_bakerstreet__pb2.Empty.SerializeToString,
            bakerstreet_dot_bakerstreet__pb2.Apps.FromString,
            options,
            channel_credentials,
            insecure,
            call_credentials,
            compression,
            wait_for_ready,
            timeout,
            metadata,
            _registered_method=True)
