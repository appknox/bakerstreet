# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
import grpc

import bakerstreet_pb2 as bakerstreet__pb2


class AgentStub(object):
  # missing associated documentation comment in .proto file
  pass

  def __init__(self, channel):
    """Constructor.

    Args:
      channel: A grpc.Channel.
    """
    self.Info = channel.unary_unary(
        '/com.appknox.bakerstreet.Agent/Info',
        request_serializer=bakerstreet__pb2.Message.SerializeToString,
        response_deserializer=bakerstreet__pb2.Device.FromString,
        )
    self.Echo = channel.unary_unary(
        '/com.appknox.bakerstreet.Agent/Echo',
        request_serializer=bakerstreet__pb2.Message.SerializeToString,
        response_deserializer=bakerstreet__pb2.Message.FromString,
        )
    self.LaunchApp = channel.unary_unary(
        '/com.appknox.bakerstreet.Agent/LaunchApp',
        request_serializer=bakerstreet__pb2.Message.SerializeToString,
        response_deserializer=bakerstreet__pb2.Message.FromString,
        )
    self.HealthCheck = channel.unary_unary(
        '/com.appknox.bakerstreet.Agent/HealthCheck',
        request_serializer=bakerstreet__pb2.Message.SerializeToString,
        response_deserializer=bakerstreet__pb2.Message.FromString,
        )
    self.RemovePackage = channel.unary_unary(
        '/com.appknox.bakerstreet.Agent/RemovePackage',
        request_serializer=bakerstreet__pb2.Message.SerializeToString,
        response_deserializer=bakerstreet__pb2.Message.FromString,
        )
    self.InstallPackage = channel.unary_unary(
        '/com.appknox.bakerstreet.Agent/InstallPackage',
        request_serializer=bakerstreet__pb2.Message.SerializeToString,
        response_deserializer=bakerstreet__pb2.Message.FromString,
        )
    self.ListPackages = channel.unary_stream(
        '/com.appknox.bakerstreet.Agent/ListPackages',
        request_serializer=bakerstreet__pb2.Message.SerializeToString,
        response_deserializer=bakerstreet__pb2.Package.FromString,
        )


class AgentServicer(object):
  # missing associated documentation comment in .proto file
  pass

  def Info(self, request, context):
    # missing associated documentation comment in .proto file
    pass
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def Echo(self, request, context):
    # missing associated documentation comment in .proto file
    pass
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def LaunchApp(self, request, context):
    # missing associated documentation comment in .proto file
    pass
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def HealthCheck(self, request, context):
    # missing associated documentation comment in .proto file
    pass
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def RemovePackage(self, request, context):
    # missing associated documentation comment in .proto file
    pass
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def InstallPackage(self, request, context):
    # missing associated documentation comment in .proto file
    pass
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def ListPackages(self, request, context):
    # missing associated documentation comment in .proto file
    pass
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')


def add_AgentServicer_to_server(servicer, server):
  rpc_method_handlers = {
      'Info': grpc.unary_unary_rpc_method_handler(
          servicer.Info,
          request_deserializer=bakerstreet__pb2.Message.FromString,
          response_serializer=bakerstreet__pb2.Device.SerializeToString,
      ),
      'Echo': grpc.unary_unary_rpc_method_handler(
          servicer.Echo,
          request_deserializer=bakerstreet__pb2.Message.FromString,
          response_serializer=bakerstreet__pb2.Message.SerializeToString,
      ),
      'LaunchApp': grpc.unary_unary_rpc_method_handler(
          servicer.LaunchApp,
          request_deserializer=bakerstreet__pb2.Message.FromString,
          response_serializer=bakerstreet__pb2.Message.SerializeToString,
      ),
      'HealthCheck': grpc.unary_unary_rpc_method_handler(
          servicer.HealthCheck,
          request_deserializer=bakerstreet__pb2.Message.FromString,
          response_serializer=bakerstreet__pb2.Message.SerializeToString,
      ),
      'RemovePackage': grpc.unary_unary_rpc_method_handler(
          servicer.RemovePackage,
          request_deserializer=bakerstreet__pb2.Message.FromString,
          response_serializer=bakerstreet__pb2.Message.SerializeToString,
      ),
      'InstallPackage': grpc.unary_unary_rpc_method_handler(
          servicer.InstallPackage,
          request_deserializer=bakerstreet__pb2.Message.FromString,
          response_serializer=bakerstreet__pb2.Message.SerializeToString,
      ),
      'ListPackages': grpc.unary_stream_rpc_method_handler(
          servicer.ListPackages,
          request_deserializer=bakerstreet__pb2.Message.FromString,
          response_serializer=bakerstreet__pb2.Package.SerializeToString,
      ),
  }
  generic_handler = grpc.method_handlers_generic_handler(
      'com.appknox.bakerstreet.Agent', rpc_method_handlers)
  server.add_generic_rpc_handlers((generic_handler,))