// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v5.26.1
// source: bakerstreet/bakerstreet.proto

package bakerstreet

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Message struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data string `protobuf:"bytes,1,opt,name=Data,proto3" json:"Data,omitempty"`
}

func (x *Message) Reset() {
	*x = Message{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bakerstreet_bakerstreet_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Message) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Message) ProtoMessage() {}

func (x *Message) ProtoReflect() protoreflect.Message {
	mi := &file_bakerstreet_bakerstreet_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Message.ProtoReflect.Descriptor instead.
func (*Message) Descriptor() ([]byte, []int) {
	return file_bakerstreet_bakerstreet_proto_rawDescGZIP(), []int{0}
}

func (x *Message) GetData() string {
	if x != nil {
		return x.Data
	}
	return ""
}

type Apps struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	App []*App `protobuf:"bytes,1,rep,name=App,proto3" json:"App,omitempty"`
}

func (x *Apps) Reset() {
	*x = Apps{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bakerstreet_bakerstreet_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Apps) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Apps) ProtoMessage() {}

func (x *Apps) ProtoReflect() protoreflect.Message {
	mi := &file_bakerstreet_bakerstreet_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Apps.ProtoReflect.Descriptor instead.
func (*Apps) Descriptor() ([]byte, []int) {
	return file_bakerstreet_bakerstreet_proto_rawDescGZIP(), []int{1}
}

func (x *Apps) GetApp() []*App {
	if x != nil {
		return x.App
	}
	return nil
}

type App struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=Name,proto3" json:"Name,omitempty"`
}

func (x *App) Reset() {
	*x = App{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bakerstreet_bakerstreet_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *App) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*App) ProtoMessage() {}

func (x *App) ProtoReflect() protoreflect.Message {
	mi := &file_bakerstreet_bakerstreet_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use App.ProtoReflect.Descriptor instead.
func (*App) Descriptor() ([]byte, []int) {
	return file_bakerstreet_bakerstreet_proto_rawDescGZIP(), []int{2}
}

func (x *App) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type Device struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uuid            string `protobuf:"bytes,1,opt,name=Uuid,proto3" json:"Uuid,omitempty"`
	IsTablet        bool   `protobuf:"varint,2,opt,name=IsTablet,proto3" json:"IsTablet,omitempty"`
	Platform        int32  `protobuf:"varint,3,opt,name=Platform,proto3" json:"Platform,omitempty"`
	PlatformVersion string `protobuf:"bytes,4,opt,name=PlatformVersion,proto3" json:"PlatformVersion,omitempty"`
	CpuArchitecture string `protobuf:"bytes,5,opt,name=CpuArchitecture,proto3" json:"CpuArchitecture,omitempty"`
	Model           string `protobuf:"bytes,6,opt,name=Model,proto3" json:"Model,omitempty"`
	Name            string `protobuf:"bytes,7,opt,name=Name,proto3" json:"Name,omitempty"`
}

func (x *Device) Reset() {
	*x = Device{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bakerstreet_bakerstreet_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Device) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Device) ProtoMessage() {}

func (x *Device) ProtoReflect() protoreflect.Message {
	mi := &file_bakerstreet_bakerstreet_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Device.ProtoReflect.Descriptor instead.
func (*Device) Descriptor() ([]byte, []int) {
	return file_bakerstreet_bakerstreet_proto_rawDescGZIP(), []int{3}
}

func (x *Device) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

func (x *Device) GetIsTablet() bool {
	if x != nil {
		return x.IsTablet
	}
	return false
}

func (x *Device) GetPlatform() int32 {
	if x != nil {
		return x.Platform
	}
	return 0
}

func (x *Device) GetPlatformVersion() string {
	if x != nil {
		return x.PlatformVersion
	}
	return ""
}

func (x *Device) GetCpuArchitecture() string {
	if x != nil {
		return x.CpuArchitecture
	}
	return ""
}

func (x *Device) GetModel() string {
	if x != nil {
		return x.Model
	}
	return ""
}

func (x *Device) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type DeviceV2 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Platform          int32  `protobuf:"varint,1,opt,name=Platform,proto3" json:"Platform,omitempty"`
	PlatformVersion   string `protobuf:"bytes,2,opt,name=PlatformVersion,proto3" json:"PlatformVersion,omitempty"`
	Name              string `protobuf:"bytes,3,opt,name=Name,proto3" json:"Name,omitempty"`
	Address           string `protobuf:"bytes,4,opt,name=Address,proto3" json:"Address,omitempty"`
	SerialNumber      string `protobuf:"bytes,5,opt,name=SerialNumber,proto3" json:"SerialNumber,omitempty"`
	Model             string `protobuf:"bytes,6,opt,name=Model,proto3" json:"Model,omitempty"`
	CpuArchitecture   string `protobuf:"bytes,7,opt,name=CpuArchitecture,proto3" json:"CpuArchitecture,omitempty"`
	IsTablet          bool   `protobuf:"varint,8,opt,name=IsTablet,proto3" json:"IsTablet,omitempty"`
	AllocationMode    string `protobuf:"bytes,9,opt,name=AllocationMode,proto3" json:"AllocationMode,omitempty"`
	SimNetwork        string `protobuf:"bytes,10,opt,name=SimNetwork,proto3" json:"SimNetwork,omitempty"`
	PhoneNumber       string `protobuf:"bytes,11,opt,name=PhoneNumber,proto3" json:"PhoneNumber,omitempty"`
	ScreenLock        string `protobuf:"bytes,12,opt,name=ScreenLock,proto3" json:"ScreenLock,omitempty"`
	VpnPackageName    string `protobuf:"bytes,13,opt,name=VpnPackageName,proto3" json:"VpnPackageName,omitempty"`
	PersistentApps    string `protobuf:"bytes,14,opt,name=PersistentApps,proto3" json:"PersistentApps,omitempty"`
	AgentVersion      string `protobuf:"bytes,15,opt,name=AgentVersion,proto3" json:"AgentVersion,omitempty"`
	FridaVersion      string `protobuf:"bytes,16,opt,name=FridaVersion,proto3" json:"FridaVersion,omitempty"`
	ExtraCapabilities string `protobuf:"bytes,17,opt,name=ExtraCapabilities,proto3" json:"ExtraCapabilities,omitempty"`
	DeviceIdentifier  string `protobuf:"bytes,18,opt,name=DeviceIdentifier,proto3" json:"DeviceIdentifier,omitempty"`
}

func (x *DeviceV2) Reset() {
	*x = DeviceV2{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bakerstreet_bakerstreet_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeviceV2) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeviceV2) ProtoMessage() {}

func (x *DeviceV2) ProtoReflect() protoreflect.Message {
	mi := &file_bakerstreet_bakerstreet_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeviceV2.ProtoReflect.Descriptor instead.
func (*DeviceV2) Descriptor() ([]byte, []int) {
	return file_bakerstreet_bakerstreet_proto_rawDescGZIP(), []int{4}
}

func (x *DeviceV2) GetPlatform() int32 {
	if x != nil {
		return x.Platform
	}
	return 0
}

func (x *DeviceV2) GetPlatformVersion() string {
	if x != nil {
		return x.PlatformVersion
	}
	return ""
}

func (x *DeviceV2) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *DeviceV2) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *DeviceV2) GetSerialNumber() string {
	if x != nil {
		return x.SerialNumber
	}
	return ""
}

func (x *DeviceV2) GetModel() string {
	if x != nil {
		return x.Model
	}
	return ""
}

func (x *DeviceV2) GetCpuArchitecture() string {
	if x != nil {
		return x.CpuArchitecture
	}
	return ""
}

func (x *DeviceV2) GetIsTablet() bool {
	if x != nil {
		return x.IsTablet
	}
	return false
}

func (x *DeviceV2) GetAllocationMode() string {
	if x != nil {
		return x.AllocationMode
	}
	return ""
}

func (x *DeviceV2) GetSimNetwork() string {
	if x != nil {
		return x.SimNetwork
	}
	return ""
}

func (x *DeviceV2) GetPhoneNumber() string {
	if x != nil {
		return x.PhoneNumber
	}
	return ""
}

func (x *DeviceV2) GetScreenLock() string {
	if x != nil {
		return x.ScreenLock
	}
	return ""
}

func (x *DeviceV2) GetVpnPackageName() string {
	if x != nil {
		return x.VpnPackageName
	}
	return ""
}

func (x *DeviceV2) GetPersistentApps() string {
	if x != nil {
		return x.PersistentApps
	}
	return ""
}

func (x *DeviceV2) GetAgentVersion() string {
	if x != nil {
		return x.AgentVersion
	}
	return ""
}

func (x *DeviceV2) GetFridaVersion() string {
	if x != nil {
		return x.FridaVersion
	}
	return ""
}

func (x *DeviceV2) GetExtraCapabilities() string {
	if x != nil {
		return x.ExtraCapabilities
	}
	return ""
}

func (x *DeviceV2) GetDeviceIdentifier() string {
	if x != nil {
		return x.DeviceIdentifier
	}
	return ""
}

type Finding struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title       string `protobuf:"bytes,1,opt,name=Title,proto3" json:"Title,omitempty"`
	Description string `protobuf:"bytes,2,opt,name=Description,proto3" json:"Description,omitempty"`
}

func (x *Finding) Reset() {
	*x = Finding{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bakerstreet_bakerstreet_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Finding) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Finding) ProtoMessage() {}

func (x *Finding) ProtoReflect() protoreflect.Message {
	mi := &file_bakerstreet_bakerstreet_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Finding.ProtoReflect.Descriptor instead.
func (*Finding) Descriptor() ([]byte, []int) {
	return file_bakerstreet_bakerstreet_proto_rawDescGZIP(), []int{5}
}

func (x *Finding) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Finding) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

type InstallReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	URL string `protobuf:"bytes,1,opt,name=URL,proto3" json:"URL,omitempty"`
}

func (x *InstallReq) Reset() {
	*x = InstallReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bakerstreet_bakerstreet_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InstallReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InstallReq) ProtoMessage() {}

func (x *InstallReq) ProtoReflect() protoreflect.Message {
	mi := &file_bakerstreet_bakerstreet_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InstallReq.ProtoReflect.Descriptor instead.
func (*InstallReq) Descriptor() ([]byte, []int) {
	return file_bakerstreet_bakerstreet_proto_rawDescGZIP(), []int{6}
}

func (x *InstallReq) GetURL() string {
	if x != nil {
		return x.URL
	}
	return ""
}

type ConfigProxyReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IP         string   `protobuf:"bytes,1,opt,name=IP,proto3" json:"IP,omitempty"`
	Port       string   `protobuf:"bytes,2,opt,name=Port,proto3" json:"Port,omitempty"`
	Hosts      []string `protobuf:"bytes,3,rep,name=Hosts,proto3" json:"Hosts,omitempty"`
	ChainHost  string   `protobuf:"bytes,4,opt,name=ChainHost,proto3" json:"ChainHost,omitempty"`
	ChainPort  string   `protobuf:"bytes,5,opt,name=ChainPort,proto3" json:"ChainPort,omitempty"`
	CaptureAll bool     `protobuf:"varint,6,opt,name=CaptureAll,proto3" json:"CaptureAll,omitempty"`
}

func (x *ConfigProxyReq) Reset() {
	*x = ConfigProxyReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bakerstreet_bakerstreet_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConfigProxyReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConfigProxyReq) ProtoMessage() {}

func (x *ConfigProxyReq) ProtoReflect() protoreflect.Message {
	mi := &file_bakerstreet_bakerstreet_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConfigProxyReq.ProtoReflect.Descriptor instead.
func (*ConfigProxyReq) Descriptor() ([]byte, []int) {
	return file_bakerstreet_bakerstreet_proto_rawDescGZIP(), []int{7}
}

func (x *ConfigProxyReq) GetIP() string {
	if x != nil {
		return x.IP
	}
	return ""
}

func (x *ConfigProxyReq) GetPort() string {
	if x != nil {
		return x.Port
	}
	return ""
}

func (x *ConfigProxyReq) GetHosts() []string {
	if x != nil {
		return x.Hosts
	}
	return nil
}

func (x *ConfigProxyReq) GetChainHost() string {
	if x != nil {
		return x.ChainHost
	}
	return ""
}

func (x *ConfigProxyReq) GetChainPort() string {
	if x != nil {
		return x.ChainPort
	}
	return ""
}

func (x *ConfigProxyReq) GetCaptureAll() bool {
	if x != nil {
		return x.CaptureAll
	}
	return false
}

type CleanOptions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DryRun bool `protobuf:"varint,1,opt,name=DryRun,proto3" json:"DryRun,omitempty"`
}

func (x *CleanOptions) Reset() {
	*x = CleanOptions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bakerstreet_bakerstreet_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CleanOptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CleanOptions) ProtoMessage() {}

func (x *CleanOptions) ProtoReflect() protoreflect.Message {
	mi := &file_bakerstreet_bakerstreet_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CleanOptions.ProtoReflect.Descriptor instead.
func (*CleanOptions) Descriptor() ([]byte, []int) {
	return file_bakerstreet_bakerstreet_proto_rawDescGZIP(), []int{8}
}

func (x *CleanOptions) GetDryRun() bool {
	if x != nil {
		return x.DryRun
	}
	return false
}

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bakerstreet_bakerstreet_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_bakerstreet_bakerstreet_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_bakerstreet_bakerstreet_proto_rawDescGZIP(), []int{9}
}

var File_bakerstreet_bakerstreet_proto protoreflect.FileDescriptor

var file_bakerstreet_bakerstreet_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x62, 0x61, 0x6b, 0x65, 0x72, 0x73, 0x74, 0x72, 0x65, 0x65, 0x74, 0x2f, 0x62, 0x61,
	0x6b, 0x65, 0x72, 0x73, 0x74, 0x72, 0x65, 0x65, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x17, 0x63, 0x6f, 0x6d, 0x2e, 0x61, 0x70, 0x70, 0x6b, 0x6e, 0x6f, 0x78, 0x2e, 0x62, 0x61, 0x6b,
	0x65, 0x72, 0x73, 0x74, 0x72, 0x65, 0x65, 0x74, 0x22, 0x1d, 0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x44, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x44, 0x61, 0x74, 0x61, 0x22, 0x36, 0x0a, 0x04, 0x41, 0x70, 0x70, 0x73, 0x12,
	0x2e, 0x0a, 0x03, 0x41, 0x70, 0x70, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x63,
	0x6f, 0x6d, 0x2e, 0x61, 0x70, 0x70, 0x6b, 0x6e, 0x6f, 0x78, 0x2e, 0x62, 0x61, 0x6b, 0x65, 0x72,
	0x73, 0x74, 0x72, 0x65, 0x65, 0x74, 0x2e, 0x41, 0x70, 0x70, 0x52, 0x03, 0x41, 0x70, 0x70, 0x22,
	0x19, 0x0a, 0x03, 0x41, 0x70, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0xd2, 0x01, 0x0a, 0x06, 0x44,
	0x65, 0x76, 0x69, 0x63, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x55, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x55, 0x75, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x49, 0x73, 0x54,
	0x61, 0x62, 0x6c, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x49, 0x73, 0x54,
	0x61, 0x62, 0x6c, 0x65, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72,
	0x6d, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72,
	0x6d, 0x12, 0x28, 0x0a, 0x0f, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x56, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x50, 0x6c, 0x61, 0x74,
	0x66, 0x6f, 0x72, 0x6d, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x28, 0x0a, 0x0f, 0x43,
	0x70, 0x75, 0x41, 0x72, 0x63, 0x68, 0x69, 0x74, 0x65, 0x63, 0x74, 0x75, 0x72, 0x65, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x43, 0x70, 0x75, 0x41, 0x72, 0x63, 0x68, 0x69, 0x74, 0x65,
	0x63, 0x74, 0x75, 0x72, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x4e,
	0x61, 0x6d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x22,
	0xfa, 0x04, 0x0a, 0x08, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x56, 0x32, 0x12, 0x1a, 0x0a, 0x08,
	0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08,
	0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x12, 0x28, 0x0a, 0x0f, 0x50, 0x6c, 0x61, 0x74,
	0x66, 0x6f, 0x72, 0x6d, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0f, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x56, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x12, 0x22, 0x0a, 0x0c, 0x53, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x53, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x4e, 0x75,
	0x6d, 0x62, 0x65, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x12, 0x28, 0x0a, 0x0f, 0x43, 0x70,
	0x75, 0x41, 0x72, 0x63, 0x68, 0x69, 0x74, 0x65, 0x63, 0x74, 0x75, 0x72, 0x65, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0f, 0x43, 0x70, 0x75, 0x41, 0x72, 0x63, 0x68, 0x69, 0x74, 0x65, 0x63,
	0x74, 0x75, 0x72, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x49, 0x73, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x74,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x49, 0x73, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x74,
	0x12, 0x26, 0x0a, 0x0e, 0x41, 0x6c, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x6f,
	0x64, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x41, 0x6c, 0x6c, 0x6f, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x6f, 0x64, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x53, 0x69, 0x6d, 0x4e,
	0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x53, 0x69,
	0x6d, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x12, 0x20, 0x0a, 0x0b, 0x50, 0x68, 0x6f, 0x6e,
	0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x50,
	0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x1e, 0x0a, 0x0a, 0x53, 0x63,
	0x72, 0x65, 0x65, 0x6e, 0x4c, 0x6f, 0x63, 0x6b, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x53, 0x63, 0x72, 0x65, 0x65, 0x6e, 0x4c, 0x6f, 0x63, 0x6b, 0x12, 0x26, 0x0a, 0x0e, 0x56, 0x70,
	0x6e, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x0d, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0e, 0x56, 0x70, 0x6e, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x26, 0x0a, 0x0e, 0x50, 0x65, 0x72, 0x73, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x74,
	0x41, 0x70, 0x70, 0x73, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x50, 0x65, 0x72, 0x73,
	0x69, 0x73, 0x74, 0x65, 0x6e, 0x74, 0x41, 0x70, 0x70, 0x73, 0x12, 0x22, 0x0a, 0x0c, 0x41, 0x67,
	0x65, 0x6e, 0x74, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0c, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x22,
	0x0a, 0x0c, 0x46, 0x72, 0x69, 0x64, 0x61, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x10,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x46, 0x72, 0x69, 0x64, 0x61, 0x56, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x12, 0x2c, 0x0a, 0x11, 0x45, 0x78, 0x74, 0x72, 0x61, 0x43, 0x61, 0x70, 0x61, 0x62,
	0x69, 0x6c, 0x69, 0x74, 0x69, 0x65, 0x73, 0x18, 0x11, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x45,
	0x78, 0x74, 0x72, 0x61, 0x43, 0x61, 0x70, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x69, 0x65, 0x73,
	0x12, 0x2a, 0x0a, 0x10, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69,
	0x66, 0x69, 0x65, 0x72, 0x18, 0x12, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x44, 0x65, 0x76, 0x69,
	0x63, 0x65, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x22, 0x41, 0x0a, 0x07,
	0x46, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x12, 0x14, 0x0a, 0x05, 0x54, 0x69, 0x74, 0x6c, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x20, 0x0a,
	0x0b, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22,
	0x1e, 0x0a, 0x0a, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x12, 0x10, 0x0a,
	0x03, 0x55, 0x52, 0x4c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x55, 0x52, 0x4c, 0x22,
	0xa6, 0x01, 0x0a, 0x0e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x52,
	0x65, 0x71, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x50, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x49, 0x50, 0x12, 0x12, 0x0a, 0x04, 0x50, 0x6f, 0x72, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x50, 0x6f, 0x72, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x48, 0x6f, 0x73, 0x74, 0x73, 0x18,
	0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x48, 0x6f, 0x73, 0x74, 0x73, 0x12, 0x1c, 0x0a, 0x09,
	0x43, 0x68, 0x61, 0x69, 0x6e, 0x48, 0x6f, 0x73, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x48, 0x6f, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x43, 0x68,
	0x61, 0x69, 0x6e, 0x50, 0x6f, 0x72, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x43,
	0x68, 0x61, 0x69, 0x6e, 0x50, 0x6f, 0x72, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x43, 0x61, 0x70, 0x74,
	0x75, 0x72, 0x65, 0x41, 0x6c, 0x6c, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x43, 0x61,
	0x70, 0x74, 0x75, 0x72, 0x65, 0x41, 0x6c, 0x6c, 0x22, 0x26, 0x0a, 0x0c, 0x43, 0x6c, 0x65, 0x61,
	0x6e, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x44, 0x72, 0x79, 0x52,
	0x75, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x44, 0x72, 0x79, 0x52, 0x75, 0x6e,
	0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x32, 0xd5, 0x07, 0x0a, 0x08, 0x4d, 0x6f,
	0x72, 0x69, 0x61, 0x72, 0x74, 0x79, 0x12, 0x4a, 0x0a, 0x04, 0x45, 0x63, 0x68, 0x6f, 0x12, 0x20,
	0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x61, 0x70, 0x70, 0x6b, 0x6e, 0x6f, 0x78, 0x2e, 0x62, 0x61, 0x6b,
	0x65, 0x72, 0x73, 0x74, 0x72, 0x65, 0x65, 0x74, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x1a, 0x20, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x61, 0x70, 0x70, 0x6b, 0x6e, 0x6f, 0x78, 0x2e, 0x62,
	0x61, 0x6b, 0x65, 0x72, 0x73, 0x74, 0x72, 0x65, 0x65, 0x74, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x12, 0x4b, 0x0a, 0x09, 0x4c, 0x61, 0x75, 0x6e, 0x63, 0x68, 0x41, 0x70, 0x70, 0x12,
	0x1c, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x61, 0x70, 0x70, 0x6b, 0x6e, 0x6f, 0x78, 0x2e, 0x62, 0x61,
	0x6b, 0x65, 0x72, 0x73, 0x74, 0x72, 0x65, 0x65, 0x74, 0x2e, 0x41, 0x70, 0x70, 0x1a, 0x20, 0x2e,
	0x63, 0x6f, 0x6d, 0x2e, 0x61, 0x70, 0x70, 0x6b, 0x6e, 0x6f, 0x78, 0x2e, 0x62, 0x61, 0x6b, 0x65,
	0x72, 0x73, 0x74, 0x72, 0x65, 0x65, 0x74, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12,
	0x4e, 0x0a, 0x0a, 0x43, 0x6c, 0x65, 0x61, 0x72, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x12, 0x1e, 0x2e,
	0x63, 0x6f, 0x6d, 0x2e, 0x61, 0x70, 0x70, 0x6b, 0x6e, 0x6f, 0x78, 0x2e, 0x62, 0x61, 0x6b, 0x65,
	0x72, 0x73, 0x74, 0x72, 0x65, 0x65, 0x74, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x20, 0x2e,
	0x63, 0x6f, 0x6d, 0x2e, 0x61, 0x70, 0x70, 0x6b, 0x6e, 0x6f, 0x78, 0x2e, 0x62, 0x61, 0x6b, 0x65,
	0x72, 0x73, 0x74, 0x72, 0x65, 0x65, 0x74, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12,
	0x4f, 0x0a, 0x0b, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x12, 0x1e,
	0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x61, 0x70, 0x70, 0x6b, 0x6e, 0x6f, 0x78, 0x2e, 0x62, 0x61, 0x6b,
	0x65, 0x72, 0x73, 0x74, 0x72, 0x65, 0x65, 0x74, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x20,
	0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x61, 0x70, 0x70, 0x6b, 0x6e, 0x6f, 0x78, 0x2e, 0x62, 0x61, 0x6b,
	0x65, 0x72, 0x73, 0x74, 0x72, 0x65, 0x65, 0x74, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x12, 0x4f, 0x0a, 0x0d, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67,
	0x65, 0x12, 0x1c, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x61, 0x70, 0x70, 0x6b, 0x6e, 0x6f, 0x78, 0x2e,
	0x62, 0x61, 0x6b, 0x65, 0x72, 0x73, 0x74, 0x72, 0x65, 0x65, 0x74, 0x2e, 0x41, 0x70, 0x70, 0x1a,
	0x20, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x61, 0x70, 0x70, 0x6b, 0x6e, 0x6f, 0x78, 0x2e, 0x62, 0x61,
	0x6b, 0x65, 0x72, 0x73, 0x74, 0x72, 0x65, 0x65, 0x74, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x12, 0x57, 0x0a, 0x0e, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x50, 0x61, 0x63, 0x6b,
	0x61, 0x67, 0x65, 0x12, 0x23, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x61, 0x70, 0x70, 0x6b, 0x6e, 0x6f,
	0x78, 0x2e, 0x62, 0x61, 0x6b, 0x65, 0x72, 0x73, 0x74, 0x72, 0x65, 0x65, 0x74, 0x2e, 0x49, 0x6e,
	0x73, 0x74, 0x61, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x1a, 0x20, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x61,
	0x70, 0x70, 0x6b, 0x6e, 0x6f, 0x78, 0x2e, 0x62, 0x61, 0x6b, 0x65, 0x72, 0x73, 0x74, 0x72, 0x65,
	0x65, 0x74, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x5b, 0x0a, 0x0e, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x65, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x12, 0x27, 0x2e, 0x63,
	0x6f, 0x6d, 0x2e, 0x61, 0x70, 0x70, 0x6b, 0x6e, 0x6f, 0x78, 0x2e, 0x62, 0x61, 0x6b, 0x65, 0x72,
	0x73, 0x74, 0x72, 0x65, 0x65, 0x74, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x50, 0x72, 0x6f,
	0x78, 0x79, 0x52, 0x65, 0x71, 0x1a, 0x20, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x61, 0x70, 0x70, 0x6b,
	0x6e, 0x6f, 0x78, 0x2e, 0x62, 0x61, 0x6b, 0x65, 0x72, 0x73, 0x74, 0x72, 0x65, 0x65, 0x74, 0x2e,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x51, 0x0a, 0x0f, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x75, 0x72, 0x65, 0x47, 0x61, 0x64, 0x67, 0x65, 0x74, 0x12, 0x1c, 0x2e, 0x63, 0x6f, 0x6d,
	0x2e, 0x61, 0x70, 0x70, 0x6b, 0x6e, 0x6f, 0x78, 0x2e, 0x62, 0x61, 0x6b, 0x65, 0x72, 0x73, 0x74,
	0x72, 0x65, 0x65, 0x74, 0x2e, 0x41, 0x70, 0x70, 0x1a, 0x20, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x61,
	0x70, 0x70, 0x6b, 0x6e, 0x6f, 0x78, 0x2e, 0x62, 0x61, 0x6b, 0x65, 0x72, 0x73, 0x74, 0x72, 0x65,
	0x65, 0x74, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x50, 0x0a, 0x05, 0x43, 0x6c,
	0x65, 0x61, 0x6e, 0x12, 0x25, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x61, 0x70, 0x70, 0x6b, 0x6e, 0x6f,
	0x78, 0x2e, 0x62, 0x61, 0x6b, 0x65, 0x72, 0x73, 0x74, 0x72, 0x65, 0x65, 0x74, 0x2e, 0x43, 0x6c,
	0x65, 0x61, 0x6e, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x1a, 0x20, 0x2e, 0x63, 0x6f, 0x6d,
	0x2e, 0x61, 0x70, 0x70, 0x6b, 0x6e, 0x6f, 0x78, 0x2e, 0x62, 0x61, 0x6b, 0x65, 0x72, 0x73, 0x74,
	0x72, 0x65, 0x65, 0x74, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x47, 0x0a, 0x04,
	0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1e, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x61, 0x70, 0x70, 0x6b, 0x6e,
	0x6f, 0x78, 0x2e, 0x62, 0x61, 0x6b, 0x65, 0x72, 0x73, 0x74, 0x72, 0x65, 0x65, 0x74, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x1a, 0x1f, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x61, 0x70, 0x70, 0x6b, 0x6e,
	0x6f, 0x78, 0x2e, 0x62, 0x61, 0x6b, 0x65, 0x72, 0x73, 0x74, 0x72, 0x65, 0x65, 0x74, 0x2e, 0x44,
	0x65, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4b, 0x0a, 0x06, 0x49, 0x6e, 0x66, 0x6f, 0x56, 0x32, 0x12,
	0x1e, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x61, 0x70, 0x70, 0x6b, 0x6e, 0x6f, 0x78, 0x2e, 0x62, 0x61,
	0x6b, 0x65, 0x72, 0x73, 0x74, 0x72, 0x65, 0x65, 0x74, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a,
	0x21, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x61, 0x70, 0x70, 0x6b, 0x6e, 0x6f, 0x78, 0x2e, 0x62, 0x61,
	0x6b, 0x65, 0x72, 0x73, 0x74, 0x72, 0x65, 0x65, 0x74, 0x2e, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65,
	0x56, 0x32, 0x12, 0x4d, 0x0a, 0x0c, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67,
	0x65, 0x73, 0x12, 0x1e, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x61, 0x70, 0x70, 0x6b, 0x6e, 0x6f, 0x78,
	0x2e, 0x62, 0x61, 0x6b, 0x65, 0x72, 0x73, 0x74, 0x72, 0x65, 0x65, 0x74, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x1a, 0x1d, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x61, 0x70, 0x70, 0x6b, 0x6e, 0x6f, 0x78,
	0x2e, 0x62, 0x61, 0x6b, 0x65, 0x72, 0x73, 0x74, 0x72, 0x65, 0x65, 0x74, 0x2e, 0x41, 0x70, 0x70,
	0x73, 0x42, 0x20, 0x5a, 0x1e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x61, 0x70, 0x70, 0x6b, 0x6e, 0x6f, 0x78, 0x2f, 0x62, 0x61, 0x6b, 0x65, 0x72, 0x73, 0x74, 0x72,
	0x65, 0x65, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_bakerstreet_bakerstreet_proto_rawDescOnce sync.Once
	file_bakerstreet_bakerstreet_proto_rawDescData = file_bakerstreet_bakerstreet_proto_rawDesc
)

func file_bakerstreet_bakerstreet_proto_rawDescGZIP() []byte {
	file_bakerstreet_bakerstreet_proto_rawDescOnce.Do(func() {
		file_bakerstreet_bakerstreet_proto_rawDescData = protoimpl.X.CompressGZIP(file_bakerstreet_bakerstreet_proto_rawDescData)
	})
	return file_bakerstreet_bakerstreet_proto_rawDescData
}

var file_bakerstreet_bakerstreet_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_bakerstreet_bakerstreet_proto_goTypes = []interface{}{
	(*Message)(nil),        // 0: com.appknox.bakerstreet.Message
	(*Apps)(nil),           // 1: com.appknox.bakerstreet.Apps
	(*App)(nil),            // 2: com.appknox.bakerstreet.App
	(*Device)(nil),         // 3: com.appknox.bakerstreet.Device
	(*DeviceV2)(nil),       // 4: com.appknox.bakerstreet.DeviceV2
	(*Finding)(nil),        // 5: com.appknox.bakerstreet.Finding
	(*InstallReq)(nil),     // 6: com.appknox.bakerstreet.InstallReq
	(*ConfigProxyReq)(nil), // 7: com.appknox.bakerstreet.ConfigProxyReq
	(*CleanOptions)(nil),   // 8: com.appknox.bakerstreet.CleanOptions
	(*Empty)(nil),          // 9: com.appknox.bakerstreet.Empty
}
var file_bakerstreet_bakerstreet_proto_depIdxs = []int32{
	2,  // 0: com.appknox.bakerstreet.Apps.App:type_name -> com.appknox.bakerstreet.App
	0,  // 1: com.appknox.bakerstreet.Moriarty.Echo:input_type -> com.appknox.bakerstreet.Message
	2,  // 2: com.appknox.bakerstreet.Moriarty.LaunchApp:input_type -> com.appknox.bakerstreet.App
	9,  // 3: com.appknox.bakerstreet.Moriarty.ClearProxy:input_type -> com.appknox.bakerstreet.Empty
	9,  // 4: com.appknox.bakerstreet.Moriarty.HealthCheck:input_type -> com.appknox.bakerstreet.Empty
	2,  // 5: com.appknox.bakerstreet.Moriarty.RemovePackage:input_type -> com.appknox.bakerstreet.App
	6,  // 6: com.appknox.bakerstreet.Moriarty.InstallPackage:input_type -> com.appknox.bakerstreet.InstallReq
	7,  // 7: com.appknox.bakerstreet.Moriarty.ConfigureProxy:input_type -> com.appknox.bakerstreet.ConfigProxyReq
	2,  // 8: com.appknox.bakerstreet.Moriarty.ConfigureGadget:input_type -> com.appknox.bakerstreet.App
	8,  // 9: com.appknox.bakerstreet.Moriarty.Clean:input_type -> com.appknox.bakerstreet.CleanOptions
	9,  // 10: com.appknox.bakerstreet.Moriarty.Info:input_type -> com.appknox.bakerstreet.Empty
	9,  // 11: com.appknox.bakerstreet.Moriarty.InfoV2:input_type -> com.appknox.bakerstreet.Empty
	9,  // 12: com.appknox.bakerstreet.Moriarty.ListPackages:input_type -> com.appknox.bakerstreet.Empty
	0,  // 13: com.appknox.bakerstreet.Moriarty.Echo:output_type -> com.appknox.bakerstreet.Message
	0,  // 14: com.appknox.bakerstreet.Moriarty.LaunchApp:output_type -> com.appknox.bakerstreet.Message
	0,  // 15: com.appknox.bakerstreet.Moriarty.ClearProxy:output_type -> com.appknox.bakerstreet.Message
	0,  // 16: com.appknox.bakerstreet.Moriarty.HealthCheck:output_type -> com.appknox.bakerstreet.Message
	0,  // 17: com.appknox.bakerstreet.Moriarty.RemovePackage:output_type -> com.appknox.bakerstreet.Message
	0,  // 18: com.appknox.bakerstreet.Moriarty.InstallPackage:output_type -> com.appknox.bakerstreet.Message
	0,  // 19: com.appknox.bakerstreet.Moriarty.ConfigureProxy:output_type -> com.appknox.bakerstreet.Message
	0,  // 20: com.appknox.bakerstreet.Moriarty.ConfigureGadget:output_type -> com.appknox.bakerstreet.Message
	0,  // 21: com.appknox.bakerstreet.Moriarty.Clean:output_type -> com.appknox.bakerstreet.Message
	3,  // 22: com.appknox.bakerstreet.Moriarty.Info:output_type -> com.appknox.bakerstreet.Device
	4,  // 23: com.appknox.bakerstreet.Moriarty.InfoV2:output_type -> com.appknox.bakerstreet.DeviceV2
	1,  // 24: com.appknox.bakerstreet.Moriarty.ListPackages:output_type -> com.appknox.bakerstreet.Apps
	13, // [13:25] is the sub-list for method output_type
	1,  // [1:13] is the sub-list for method input_type
	1,  // [1:1] is the sub-list for extension type_name
	1,  // [1:1] is the sub-list for extension extendee
	0,  // [0:1] is the sub-list for field type_name
}

func init() { file_bakerstreet_bakerstreet_proto_init() }
func file_bakerstreet_bakerstreet_proto_init() {
	if File_bakerstreet_bakerstreet_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_bakerstreet_bakerstreet_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Message); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_bakerstreet_bakerstreet_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Apps); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_bakerstreet_bakerstreet_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*App); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_bakerstreet_bakerstreet_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Device); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_bakerstreet_bakerstreet_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeviceV2); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_bakerstreet_bakerstreet_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Finding); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_bakerstreet_bakerstreet_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InstallReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_bakerstreet_bakerstreet_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConfigProxyReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_bakerstreet_bakerstreet_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CleanOptions); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_bakerstreet_bakerstreet_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Empty); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_bakerstreet_bakerstreet_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_bakerstreet_bakerstreet_proto_goTypes,
		DependencyIndexes: file_bakerstreet_bakerstreet_proto_depIdxs,
		MessageInfos:      file_bakerstreet_bakerstreet_proto_msgTypes,
	}.Build()
	File_bakerstreet_bakerstreet_proto = out.File
	file_bakerstreet_bakerstreet_proto_rawDesc = nil
	file_bakerstreet_bakerstreet_proto_goTypes = nil
	file_bakerstreet_bakerstreet_proto_depIdxs = nil
}
