# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# NO CHECKED-IN PROTOBUF GENCODE
# source: bakerstreet/bakerstreet.proto
# Protobuf Python Version: 5.28.1
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import runtime_version as _runtime_version
from google.protobuf import symbol_database as _symbol_database
from google.protobuf.internal import builder as _builder
_runtime_version.ValidateProtobufRuntimeVersion(
    _runtime_version.Domain.PUBLIC,
    5,
    28,
    1,
    '',
    'bakerstreet/bakerstreet.proto'
)
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x1d\x62\x61kerstreet/bakerstreet.proto\x12\x17\x63om.appknox.bakerstreet\"\x17\n\x07Message\x12\x0c\n\x04\x44\x61ta\x18\x01 \x01(\t\"1\n\x04\x41pps\x12)\n\x03\x41pp\x18\x01 \x03(\x0b\x32\x1c.com.appknox.bakerstreet.App\"\x13\n\x03\x41pp\x12\x0c\n\x04Name\x18\x01 \x01(\t\"\x89\x01\n\x06\x44\x65vice\x12\x0c\n\x04Uuid\x18\x01 \x01(\t\x12\x10\n\x08IsTablet\x18\x02 \x01(\x08\x12\x10\n\x08Platform\x18\x03 \x01(\x05\x12\x17\n\x0fPlatformVersion\x18\x04 \x01(\t\x12\x17\n\x0f\x43puArchitecture\x18\x05 \x01(\t\x12\r\n\x05Model\x18\x06 \x01(\t\x12\x0c\n\x04Name\x18\x07 \x01(\t\"\xb5\x03\n\x08\x44\x65viceV2\x12\x0c\n\x04Uuid\x18\x01 \x01(\t\x12\x10\n\x08IsTablet\x18\x02 \x01(\x08\x12\x10\n\x08Platform\x18\x03 \x01(\x05\x12\x17\n\x0fPlatformVersion\x18\x04 \x01(\t\x12\x17\n\x0f\x43puArchitecture\x18\x05 \x01(\t\x12\r\n\x05Model\x18\x06 \x01(\t\x12\x0c\n\x04Name\x18\x07 \x01(\t\x12\x18\n\x10\x44\x65viceIdentifier\x18\x08 \x01(\t\x12\x14\n\x0c\x43\x61nRunManual\x18\t \x01(\x08\x12\x17\n\x0f\x43\x61nRunAutomated\x18\n \x01(\x08\x12\x1b\n\x13HasAutomationEngine\x18\x0b \x01(\x08\x12\x0e\n\x06HasVnc\x18\x0c \x01(\x08\x12\x12\n\nSimNetwork\x18\x0e \x01(\t\x12\x13\n\x0bPhoneNumber\x18\x0f \x01(\t\x12\x12\n\nScreenLock\x18\x10 \x01(\t\x12\x16\n\x0eVpnPackageName\x18\x11 \x01(\t\x12\x16\n\x0ePersistentApps\x18\x12 \x01(\t\x12\x19\n\x11\x45xtraCapabilities\x18\x13 \x01(\t\x12\x14\n\x0c\x41gentVersion\x18\x14 \x01(\t\x12\x14\n\x0c\x46ridaVersion\x18\x15 \x01(\t\"-\n\x07\x46inding\x12\r\n\x05Title\x18\x01 \x01(\t\x12\x13\n\x0b\x44\x65scription\x18\x02 \x01(\t\"\x19\n\nInstallReq\x12\x0b\n\x03URL\x18\x01 \x01(\t\"s\n\x0e\x43onfigProxyReq\x12\n\n\x02IP\x18\x01 \x01(\t\x12\x0c\n\x04Port\x18\x02 \x01(\t\x12\r\n\x05Hosts\x18\x03 \x03(\t\x12\x11\n\tChainHost\x18\x04 \x01(\t\x12\x11\n\tChainPort\x18\x05 \x01(\t\x12\x12\n\nCaptureAll\x18\x06 \x01(\x08\".\n\x0f\x41utoPilotConfig\x12\x1b\n\x13\x41utoPilotPreference\x18\x01 \x01(\t\"\x1e\n\x0c\x43leanOptions\x12\x0e\n\x06\x44ryRun\x18\x01 \x01(\x08\"\x07\n\x05\x45mpty2\xb7\x08\n\x08Moriarty\x12J\n\x04\x45\x63ho\x12 .com.appknox.bakerstreet.Message\x1a .com.appknox.bakerstreet.Message\x12K\n\tLaunchApp\x12\x1c.com.appknox.bakerstreet.App\x1a .com.appknox.bakerstreet.Message\x12N\n\nClearProxy\x12\x1e.com.appknox.bakerstreet.Empty\x1a .com.appknox.bakerstreet.Message\x12O\n\x0bHealthCheck\x12\x1e.com.appknox.bakerstreet.Empty\x1a .com.appknox.bakerstreet.Message\x12O\n\rRemovePackage\x12\x1c.com.appknox.bakerstreet.App\x1a .com.appknox.bakerstreet.Message\x12W\n\x0eInstallPackage\x12#.com.appknox.bakerstreet.InstallReq\x1a .com.appknox.bakerstreet.Message\x12[\n\x0e\x43onfigureProxy\x12\'.com.appknox.bakerstreet.ConfigProxyReq\x1a .com.appknox.bakerstreet.Message\x12Q\n\x0f\x43onfigureGadget\x12\x1c.com.appknox.bakerstreet.App\x1a .com.appknox.bakerstreet.Message\x12`\n\x12\x43onfigureAutopilot\x12(.com.appknox.bakerstreet.AutoPilotConfig\x1a .com.appknox.bakerstreet.Message\x12P\n\x05\x43lean\x12%.com.appknox.bakerstreet.CleanOptions\x1a .com.appknox.bakerstreet.Message\x12G\n\x04Info\x12\x1e.com.appknox.bakerstreet.Empty\x1a\x1f.com.appknox.bakerstreet.Device\x12K\n\x06InfoV2\x12\x1e.com.appknox.bakerstreet.Empty\x1a!.com.appknox.bakerstreet.DeviceV2\x12M\n\x0cListPackages\x12\x1e.com.appknox.bakerstreet.Empty\x1a\x1d.com.appknox.bakerstreet.AppsB Z\x1egithub.com/appknox/bakerstreetb\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'bakerstreet.bakerstreet_pb2', _globals)
if not _descriptor._USE_C_DESCRIPTORS:
  _globals['DESCRIPTOR']._loaded_options = None
  _globals['DESCRIPTOR']._serialized_options = b'Z\036github.com/appknox/bakerstreet'
  _globals['_MESSAGE']._serialized_start=58
  _globals['_MESSAGE']._serialized_end=81
  _globals['_APPS']._serialized_start=83
  _globals['_APPS']._serialized_end=132
  _globals['_APP']._serialized_start=134
  _globals['_APP']._serialized_end=153
  _globals['_DEVICE']._serialized_start=156
  _globals['_DEVICE']._serialized_end=293
  _globals['_DEVICEV2']._serialized_start=296
  _globals['_DEVICEV2']._serialized_end=733
  _globals['_FINDING']._serialized_start=735
  _globals['_FINDING']._serialized_end=780
  _globals['_INSTALLREQ']._serialized_start=782
  _globals['_INSTALLREQ']._serialized_end=807
  _globals['_CONFIGPROXYREQ']._serialized_start=809
  _globals['_CONFIGPROXYREQ']._serialized_end=924
  _globals['_AUTOPILOTCONFIG']._serialized_start=926
  _globals['_AUTOPILOTCONFIG']._serialized_end=972
  _globals['_CLEANOPTIONS']._serialized_start=974
  _globals['_CLEANOPTIONS']._serialized_end=1004
  _globals['_EMPTY']._serialized_start=1006
  _globals['_EMPTY']._serialized_end=1013
  _globals['_MORIARTY']._serialized_start=1016
  _globals['_MORIARTY']._serialized_end=2095
# @@protoc_insertion_point(module_scope)
