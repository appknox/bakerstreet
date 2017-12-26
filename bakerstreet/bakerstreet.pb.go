// Code generated by protoc-gen-go. DO NOT EDIT.
// source: bakerstreet/bakerstreet.proto

/*
Package com_appknox_bakerstreet is a generated protocol buffer package.

It is generated from these files:
	bakerstreet/bakerstreet.proto

It has these top-level messages:
	Message
	Packages
	Package
	Device
	Finding
	InstallReq
	ConfigProxyReq
	Empty
*/
package com_appknox_bakerstreet

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Message struct {
	Data string `protobuf:"bytes,1,opt,name=Data" json:"Data,omitempty"`
}

func (m *Message) Reset()                    { *m = Message{} }
func (m *Message) String() string            { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()               {}
func (*Message) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Message) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

type Packages struct {
	Package []*Package `protobuf:"bytes,1,rep,name=Package" json:"Package,omitempty"`
}

func (m *Packages) Reset()                    { *m = Packages{} }
func (m *Packages) String() string            { return proto.CompactTextString(m) }
func (*Packages) ProtoMessage()               {}
func (*Packages) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Packages) GetPackage() []*Package {
	if m != nil {
		return m.Package
	}
	return nil
}

type Package struct {
	Name string `protobuf:"bytes,1,opt,name=Name" json:"Name,omitempty"`
}

func (m *Package) Reset()                    { *m = Package{} }
func (m *Package) String() string            { return proto.CompactTextString(m) }
func (*Package) ProtoMessage()               {}
func (*Package) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Package) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type Device struct {
	Uuid            string `protobuf:"bytes,1,opt,name=Uuid" json:"Uuid,omitempty"`
	IsTablet        bool   `protobuf:"varint,2,opt,name=IsTablet" json:"IsTablet,omitempty"`
	Platform        int32  `protobuf:"varint,3,opt,name=Platform" json:"Platform,omitempty"`
	PlatformVersion string `protobuf:"bytes,4,opt,name=PlatformVersion" json:"PlatformVersion,omitempty"`
}

func (m *Device) Reset()                    { *m = Device{} }
func (m *Device) String() string            { return proto.CompactTextString(m) }
func (*Device) ProtoMessage()               {}
func (*Device) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *Device) GetUuid() string {
	if m != nil {
		return m.Uuid
	}
	return ""
}

func (m *Device) GetIsTablet() bool {
	if m != nil {
		return m.IsTablet
	}
	return false
}

func (m *Device) GetPlatform() int32 {
	if m != nil {
		return m.Platform
	}
	return 0
}

func (m *Device) GetPlatformVersion() string {
	if m != nil {
		return m.PlatformVersion
	}
	return ""
}

type Finding struct {
	Title       string `protobuf:"bytes,1,opt,name=Title" json:"Title,omitempty"`
	Description string `protobuf:"bytes,2,opt,name=Description" json:"Description,omitempty"`
}

func (m *Finding) Reset()                    { *m = Finding{} }
func (m *Finding) String() string            { return proto.CompactTextString(m) }
func (*Finding) ProtoMessage()               {}
func (*Finding) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *Finding) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Finding) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

type InstallReq struct {
	URL string `protobuf:"bytes,1,opt,name=URL" json:"URL,omitempty"`
}

func (m *InstallReq) Reset()                    { *m = InstallReq{} }
func (m *InstallReq) String() string            { return proto.CompactTextString(m) }
func (*InstallReq) ProtoMessage()               {}
func (*InstallReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *InstallReq) GetURL() string {
	if m != nil {
		return m.URL
	}
	return ""
}

type ConfigProxyReq struct {
}

func (m *ConfigProxyReq) Reset()                    { *m = ConfigProxyReq{} }
func (m *ConfigProxyReq) String() string            { return proto.CompactTextString(m) }
func (*ConfigProxyReq) ProtoMessage()               {}
func (*ConfigProxyReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

type Empty struct {
}

func (m *Empty) Reset()                    { *m = Empty{} }
func (m *Empty) String() string            { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()               {}
func (*Empty) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func init() {
	proto.RegisterType((*Message)(nil), "com.appknox.bakerstreet.Message")
	proto.RegisterType((*Packages)(nil), "com.appknox.bakerstreet.Packages")
	proto.RegisterType((*Package)(nil), "com.appknox.bakerstreet.Package")
	proto.RegisterType((*Device)(nil), "com.appknox.bakerstreet.Device")
	proto.RegisterType((*Finding)(nil), "com.appknox.bakerstreet.Finding")
	proto.RegisterType((*InstallReq)(nil), "com.appknox.bakerstreet.InstallReq")
	proto.RegisterType((*ConfigProxyReq)(nil), "com.appknox.bakerstreet.ConfigProxyReq")
	proto.RegisterType((*Empty)(nil), "com.appknox.bakerstreet.Empty")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Moriarty service

type MoriartyClient interface {
	Echo(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Message, error)
	LaunchApp(ctx context.Context, in *Package, opts ...grpc.CallOption) (*Message, error)
	ClearProxy(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Message, error)
	HealthCheck(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Message, error)
	RemovePackage(ctx context.Context, in *Package, opts ...grpc.CallOption) (*Message, error)
	InstallPackage(ctx context.Context, in *InstallReq, opts ...grpc.CallOption) (*Message, error)
	ConfigureProxy(ctx context.Context, in *ConfigProxyReq, opts ...grpc.CallOption) (*Message, error)
	ConfigureGadget(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Message, error)
	Info(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Device, error)
	ListPackages(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Packages, error)
}

type moriartyClient struct {
	cc *grpc.ClientConn
}

func NewMoriartyClient(cc *grpc.ClientConn) MoriartyClient {
	return &moriartyClient{cc}
}

func (c *moriartyClient) Echo(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Message, error) {
	out := new(Message)
	err := grpc.Invoke(ctx, "/com.appknox.bakerstreet.Moriarty/Echo", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *moriartyClient) LaunchApp(ctx context.Context, in *Package, opts ...grpc.CallOption) (*Message, error) {
	out := new(Message)
	err := grpc.Invoke(ctx, "/com.appknox.bakerstreet.Moriarty/LaunchApp", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *moriartyClient) ClearProxy(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Message, error) {
	out := new(Message)
	err := grpc.Invoke(ctx, "/com.appknox.bakerstreet.Moriarty/ClearProxy", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *moriartyClient) HealthCheck(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Message, error) {
	out := new(Message)
	err := grpc.Invoke(ctx, "/com.appknox.bakerstreet.Moriarty/HealthCheck", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *moriartyClient) RemovePackage(ctx context.Context, in *Package, opts ...grpc.CallOption) (*Message, error) {
	out := new(Message)
	err := grpc.Invoke(ctx, "/com.appknox.bakerstreet.Moriarty/RemovePackage", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *moriartyClient) InstallPackage(ctx context.Context, in *InstallReq, opts ...grpc.CallOption) (*Message, error) {
	out := new(Message)
	err := grpc.Invoke(ctx, "/com.appknox.bakerstreet.Moriarty/InstallPackage", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *moriartyClient) ConfigureProxy(ctx context.Context, in *ConfigProxyReq, opts ...grpc.CallOption) (*Message, error) {
	out := new(Message)
	err := grpc.Invoke(ctx, "/com.appknox.bakerstreet.Moriarty/ConfigureProxy", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *moriartyClient) ConfigureGadget(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Message, error) {
	out := new(Message)
	err := grpc.Invoke(ctx, "/com.appknox.bakerstreet.Moriarty/ConfigureGadget", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *moriartyClient) Info(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Device, error) {
	out := new(Device)
	err := grpc.Invoke(ctx, "/com.appknox.bakerstreet.Moriarty/Info", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *moriartyClient) ListPackages(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Packages, error) {
	out := new(Packages)
	err := grpc.Invoke(ctx, "/com.appknox.bakerstreet.Moriarty/ListPackages", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Moriarty service

type MoriartyServer interface {
	Echo(context.Context, *Message) (*Message, error)
	LaunchApp(context.Context, *Package) (*Message, error)
	ClearProxy(context.Context, *Empty) (*Message, error)
	HealthCheck(context.Context, *Empty) (*Message, error)
	RemovePackage(context.Context, *Package) (*Message, error)
	InstallPackage(context.Context, *InstallReq) (*Message, error)
	ConfigureProxy(context.Context, *ConfigProxyReq) (*Message, error)
	ConfigureGadget(context.Context, *Message) (*Message, error)
	Info(context.Context, *Empty) (*Device, error)
	ListPackages(context.Context, *Empty) (*Packages, error)
}

func RegisterMoriartyServer(s *grpc.Server, srv MoriartyServer) {
	s.RegisterService(&_Moriarty_serviceDesc, srv)
}

func _Moriarty_Echo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Message)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MoriartyServer).Echo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/com.appknox.bakerstreet.Moriarty/Echo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MoriartyServer).Echo(ctx, req.(*Message))
	}
	return interceptor(ctx, in, info, handler)
}

func _Moriarty_LaunchApp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Package)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MoriartyServer).LaunchApp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/com.appknox.bakerstreet.Moriarty/LaunchApp",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MoriartyServer).LaunchApp(ctx, req.(*Package))
	}
	return interceptor(ctx, in, info, handler)
}

func _Moriarty_ClearProxy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MoriartyServer).ClearProxy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/com.appknox.bakerstreet.Moriarty/ClearProxy",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MoriartyServer).ClearProxy(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Moriarty_HealthCheck_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MoriartyServer).HealthCheck(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/com.appknox.bakerstreet.Moriarty/HealthCheck",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MoriartyServer).HealthCheck(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Moriarty_RemovePackage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Package)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MoriartyServer).RemovePackage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/com.appknox.bakerstreet.Moriarty/RemovePackage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MoriartyServer).RemovePackage(ctx, req.(*Package))
	}
	return interceptor(ctx, in, info, handler)
}

func _Moriarty_InstallPackage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InstallReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MoriartyServer).InstallPackage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/com.appknox.bakerstreet.Moriarty/InstallPackage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MoriartyServer).InstallPackage(ctx, req.(*InstallReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Moriarty_ConfigureProxy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConfigProxyReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MoriartyServer).ConfigureProxy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/com.appknox.bakerstreet.Moriarty/ConfigureProxy",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MoriartyServer).ConfigureProxy(ctx, req.(*ConfigProxyReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Moriarty_ConfigureGadget_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Message)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MoriartyServer).ConfigureGadget(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/com.appknox.bakerstreet.Moriarty/ConfigureGadget",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MoriartyServer).ConfigureGadget(ctx, req.(*Message))
	}
	return interceptor(ctx, in, info, handler)
}

func _Moriarty_Info_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MoriartyServer).Info(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/com.appknox.bakerstreet.Moriarty/Info",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MoriartyServer).Info(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Moriarty_ListPackages_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MoriartyServer).ListPackages(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/com.appknox.bakerstreet.Moriarty/ListPackages",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MoriartyServer).ListPackages(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _Moriarty_serviceDesc = grpc.ServiceDesc{
	ServiceName: "com.appknox.bakerstreet.Moriarty",
	HandlerType: (*MoriartyServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Echo",
			Handler:    _Moriarty_Echo_Handler,
		},
		{
			MethodName: "LaunchApp",
			Handler:    _Moriarty_LaunchApp_Handler,
		},
		{
			MethodName: "ClearProxy",
			Handler:    _Moriarty_ClearProxy_Handler,
		},
		{
			MethodName: "HealthCheck",
			Handler:    _Moriarty_HealthCheck_Handler,
		},
		{
			MethodName: "RemovePackage",
			Handler:    _Moriarty_RemovePackage_Handler,
		},
		{
			MethodName: "InstallPackage",
			Handler:    _Moriarty_InstallPackage_Handler,
		},
		{
			MethodName: "ConfigureProxy",
			Handler:    _Moriarty_ConfigureProxy_Handler,
		},
		{
			MethodName: "ConfigureGadget",
			Handler:    _Moriarty_ConfigureGadget_Handler,
		},
		{
			MethodName: "Info",
			Handler:    _Moriarty_Info_Handler,
		},
		{
			MethodName: "ListPackages",
			Handler:    _Moriarty_ListPackages_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "bakerstreet/bakerstreet.proto",
}

func init() { proto.RegisterFile("bakerstreet/bakerstreet.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 461 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x94, 0xcf, 0x6e, 0xd3, 0x40,
	0x10, 0xc6, 0xe5, 0xe6, 0x6f, 0x27, 0xd0, 0x56, 0x2b, 0x24, 0xac, 0x4a, 0x2d, 0xc6, 0x1c, 0xf0,
	0xc9, 0x48, 0xe5, 0xc6, 0xad, 0x4a, 0xda, 0x12, 0x94, 0x96, 0x60, 0x1a, 0x38, 0x70, 0x9a, 0x38,
	0x13, 0x67, 0x15, 0xdb, 0xeb, 0xee, 0x6e, 0xaa, 0xe6, 0xc2, 0xe3, 0xf1, 0x5c, 0xc8, 0x8e, 0x1d,
	0x42, 0x25, 0xcb, 0x3e, 0xe4, 0xf6, 0xad, 0xbf, 0x9d, 0xdf, 0xec, 0xec, 0x8c, 0x17, 0xce, 0xa6,
	0xb8, 0x24, 0xa9, 0xb4, 0x24, 0xd2, 0x1f, 0x76, 0xb4, 0x9b, 0x48, 0xa1, 0x05, 0x7b, 0xed, 0x8b,
	0xc8, 0xc5, 0x24, 0x59, 0xc6, 0xe2, 0xc9, 0xdd, 0xb1, 0xed, 0x33, 0xe8, 0xdc, 0x92, 0x52, 0x18,
	0x10, 0x63, 0xd0, 0x1c, 0xa0, 0x46, 0xd3, 0xb0, 0x0c, 0xe7, 0xd0, 0xcb, 0xb4, 0x7d, 0x0d, 0xdd,
	0x31, 0xfa, 0x4b, 0x0c, 0x48, 0xb1, 0x4f, 0xd0, 0xc9, 0xb5, 0x69, 0x58, 0x0d, 0xa7, 0x77, 0x61,
	0xb9, 0x25, 0x54, 0x37, 0xdf, 0xe7, 0x15, 0x01, 0x69, 0x9a, 0x5c, 0xa6, 0x69, 0xee, 0x30, 0xa2,
	0x22, 0x4d, 0xaa, 0xed, 0xdf, 0xd0, 0x1e, 0xd0, 0x23, 0xf7, 0x33, 0x77, 0xb2, 0xe2, 0xb3, 0xc2,
	0x4d, 0x35, 0x3b, 0x85, 0xee, 0x50, 0xdd, 0xe3, 0x34, 0x24, 0x6d, 0x1e, 0x58, 0x86, 0xd3, 0xf5,
	0xb6, 0xeb, 0xd4, 0x1b, 0x87, 0xa8, 0xe7, 0x42, 0x46, 0x66, 0xc3, 0x32, 0x9c, 0x96, 0xb7, 0x5d,
	0x33, 0x07, 0x8e, 0x0b, 0xfd, 0x83, 0xa4, 0xe2, 0x22, 0x36, 0x9b, 0x19, 0xf6, 0xf9, 0x67, 0xfb,
	0x12, 0x3a, 0xd7, 0x3c, 0x9e, 0xf1, 0x38, 0x60, 0xaf, 0xa0, 0x75, 0xcf, 0x75, 0x58, 0x9c, 0x6f,
	0xb3, 0x60, 0x16, 0xf4, 0x06, 0xa4, 0x7c, 0xc9, 0x13, 0x9d, 0x62, 0x0e, 0x32, 0x6f, 0xf7, 0x93,
	0x7d, 0x0e, 0x30, 0x8c, 0x95, 0xc6, 0x30, 0xf4, 0xe8, 0x81, 0x9d, 0x40, 0x63, 0xe2, 0x8d, 0x72,
	0x46, 0x2a, 0xed, 0x13, 0x38, 0xea, 0x8b, 0x78, 0xce, 0x83, 0xb1, 0x14, 0x4f, 0x6b, 0x8f, 0x1e,
	0xec, 0x0e, 0xb4, 0xae, 0xa2, 0x44, 0xaf, 0x2f, 0xfe, 0xb4, 0xa1, 0x7b, 0x2b, 0x24, 0x47, 0xa9,
	0xd7, 0xec, 0x0b, 0x34, 0xaf, 0xfc, 0x85, 0x60, 0xe5, 0x97, 0x9b, 0xf7, 0xeb, 0xb4, 0x72, 0x07,
	0xfb, 0x0a, 0x87, 0x23, 0x5c, 0xc5, 0xfe, 0xe2, 0x32, 0x49, 0x58, 0x65, 0xb7, 0x6a, 0x00, 0xef,
	0x00, 0xfa, 0x21, 0xa1, 0xcc, 0x6a, 0x60, 0xe7, 0xa5, 0xfb, 0xb3, 0xba, 0x6a, 0x1d, 0xb0, 0xf7,
	0x99, 0x30, 0xd4, 0x8b, 0xfe, 0x82, 0xfc, 0xe5, 0x1e, 0x80, 0xdf, 0xe1, 0xa5, 0x47, 0x91, 0x78,
	0xa4, 0x62, 0xda, 0xf6, 0x51, 0xf5, 0x4f, 0x38, 0xca, 0x5b, 0x5b, 0x50, 0xdf, 0x95, 0xc6, 0xfc,
	0x9b, 0x81, 0x1a, 0xe0, 0x5f, 0xc5, 0x4c, 0xac, 0x24, 0x6d, 0xae, 0xf4, 0x7d, 0x69, 0xcc, 0xff,
	0xc3, 0x53, 0x03, 0x3e, 0x81, 0xe3, 0x2d, 0xfc, 0x06, 0x67, 0x01, 0xe9, 0xbd, 0xcc, 0xd4, 0x0d,
	0x34, 0x87, 0xf1, 0x5c, 0x54, 0xf6, 0xea, 0x4d, 0xa9, 0x9f, 0xff, 0xe9, 0xdf, 0xe0, 0xc5, 0x88,
	0x2b, 0xbd, 0x7d, 0x5e, 0xaa, 0x80, 0x6f, 0xab, 0x3a, 0xa9, 0xa6, 0xed, 0xec, 0xb1, 0xfb, 0xf8,
	0x37, 0x00, 0x00, 0xff, 0xff, 0x0d, 0xc6, 0x09, 0xc3, 0x0d, 0x05, 0x00, 0x00,
}
