// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v5.29.1
// source: bakerstreet/bakerstreet.proto

package bakerstreet

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// MoriartyClient is the client API for Moriarty service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MoriartyClient interface {
	Echo(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Message, error)
	LaunchApp(ctx context.Context, in *App, opts ...grpc.CallOption) (*Message, error)
	ClearProxy(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Message, error)
	HealthCheck(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Message, error)
	RemovePackage(ctx context.Context, in *App, opts ...grpc.CallOption) (*Message, error)
	InstallPackage(ctx context.Context, in *InstallReq, opts ...grpc.CallOption) (*Message, error)
	ConfigureProxy(ctx context.Context, in *ConfigProxyReq, opts ...grpc.CallOption) (*Message, error)
	ConfigureGadget(ctx context.Context, in *App, opts ...grpc.CallOption) (*Message, error)
	StartAutopilot(ctx context.Context, in *AutoPilotConfig, opts ...grpc.CallOption) (*Message, error)
	Clean(ctx context.Context, in *CleanOptions, opts ...grpc.CallOption) (*Message, error)
	Info(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Device, error)
	InfoV2(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*DeviceV2, error)
	ListPackages(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Apps, error)
}

type moriartyClient struct {
	cc grpc.ClientConnInterface
}

func NewMoriartyClient(cc grpc.ClientConnInterface) MoriartyClient {
	return &moriartyClient{cc}
}

func (c *moriartyClient) Echo(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Message, error) {
	out := new(Message)
	err := c.cc.Invoke(ctx, "/com.appknox.bakerstreet.Moriarty/Echo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *moriartyClient) LaunchApp(ctx context.Context, in *App, opts ...grpc.CallOption) (*Message, error) {
	out := new(Message)
	err := c.cc.Invoke(ctx, "/com.appknox.bakerstreet.Moriarty/LaunchApp", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *moriartyClient) ClearProxy(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Message, error) {
	out := new(Message)
	err := c.cc.Invoke(ctx, "/com.appknox.bakerstreet.Moriarty/ClearProxy", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *moriartyClient) HealthCheck(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Message, error) {
	out := new(Message)
	err := c.cc.Invoke(ctx, "/com.appknox.bakerstreet.Moriarty/HealthCheck", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *moriartyClient) RemovePackage(ctx context.Context, in *App, opts ...grpc.CallOption) (*Message, error) {
	out := new(Message)
	err := c.cc.Invoke(ctx, "/com.appknox.bakerstreet.Moriarty/RemovePackage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *moriartyClient) InstallPackage(ctx context.Context, in *InstallReq, opts ...grpc.CallOption) (*Message, error) {
	out := new(Message)
	err := c.cc.Invoke(ctx, "/com.appknox.bakerstreet.Moriarty/InstallPackage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *moriartyClient) ConfigureProxy(ctx context.Context, in *ConfigProxyReq, opts ...grpc.CallOption) (*Message, error) {
	out := new(Message)
	err := c.cc.Invoke(ctx, "/com.appknox.bakerstreet.Moriarty/ConfigureProxy", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *moriartyClient) ConfigureGadget(ctx context.Context, in *App, opts ...grpc.CallOption) (*Message, error) {
	out := new(Message)
	err := c.cc.Invoke(ctx, "/com.appknox.bakerstreet.Moriarty/ConfigureGadget", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *moriartyClient) StartAutopilot(ctx context.Context, in *AutoPilotConfig, opts ...grpc.CallOption) (*Message, error) {
	out := new(Message)
	err := c.cc.Invoke(ctx, "/com.appknox.bakerstreet.Moriarty/StartAutopilot", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *moriartyClient) Clean(ctx context.Context, in *CleanOptions, opts ...grpc.CallOption) (*Message, error) {
	out := new(Message)
	err := c.cc.Invoke(ctx, "/com.appknox.bakerstreet.Moriarty/Clean", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *moriartyClient) Info(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Device, error) {
	out := new(Device)
	err := c.cc.Invoke(ctx, "/com.appknox.bakerstreet.Moriarty/Info", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *moriartyClient) InfoV2(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*DeviceV2, error) {
	out := new(DeviceV2)
	err := c.cc.Invoke(ctx, "/com.appknox.bakerstreet.Moriarty/InfoV2", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *moriartyClient) ListPackages(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Apps, error) {
	out := new(Apps)
	err := c.cc.Invoke(ctx, "/com.appknox.bakerstreet.Moriarty/ListPackages", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MoriartyServer is the server API for Moriarty service.
// All implementations must embed UnimplementedMoriartyServer
// for forward compatibility
type MoriartyServer interface {
	Echo(context.Context, *Message) (*Message, error)
	LaunchApp(context.Context, *App) (*Message, error)
	ClearProxy(context.Context, *Empty) (*Message, error)
	HealthCheck(context.Context, *Empty) (*Message, error)
	RemovePackage(context.Context, *App) (*Message, error)
	InstallPackage(context.Context, *InstallReq) (*Message, error)
	ConfigureProxy(context.Context, *ConfigProxyReq) (*Message, error)
	ConfigureGadget(context.Context, *App) (*Message, error)
	StartAutopilot(context.Context, *AutoPilotConfig) (*Message, error)
	Clean(context.Context, *CleanOptions) (*Message, error)
	Info(context.Context, *Empty) (*Device, error)
	InfoV2(context.Context, *Empty) (*DeviceV2, error)
	ListPackages(context.Context, *Empty) (*Apps, error)
	mustEmbedUnimplementedMoriartyServer()
}

// UnimplementedMoriartyServer must be embedded to have forward compatible implementations.
type UnimplementedMoriartyServer struct {
}

func (UnimplementedMoriartyServer) Echo(context.Context, *Message) (*Message, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Echo not implemented")
}
func (UnimplementedMoriartyServer) LaunchApp(context.Context, *App) (*Message, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LaunchApp not implemented")
}
func (UnimplementedMoriartyServer) ClearProxy(context.Context, *Empty) (*Message, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ClearProxy not implemented")
}
func (UnimplementedMoriartyServer) HealthCheck(context.Context, *Empty) (*Message, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HealthCheck not implemented")
}
func (UnimplementedMoriartyServer) RemovePackage(context.Context, *App) (*Message, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemovePackage not implemented")
}
func (UnimplementedMoriartyServer) InstallPackage(context.Context, *InstallReq) (*Message, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InstallPackage not implemented")
}
func (UnimplementedMoriartyServer) ConfigureProxy(context.Context, *ConfigProxyReq) (*Message, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ConfigureProxy not implemented")
}
func (UnimplementedMoriartyServer) ConfigureGadget(context.Context, *App) (*Message, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ConfigureGadget not implemented")
}
func (UnimplementedMoriartyServer) StartAutopilot(context.Context, *AutoPilotConfig) (*Message, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StartAutopilot not implemented")
}
func (UnimplementedMoriartyServer) Clean(context.Context, *CleanOptions) (*Message, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Clean not implemented")
}
func (UnimplementedMoriartyServer) Info(context.Context, *Empty) (*Device, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Info not implemented")
}
func (UnimplementedMoriartyServer) InfoV2(context.Context, *Empty) (*DeviceV2, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InfoV2 not implemented")
}
func (UnimplementedMoriartyServer) ListPackages(context.Context, *Empty) (*Apps, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListPackages not implemented")
}
func (UnimplementedMoriartyServer) mustEmbedUnimplementedMoriartyServer() {}

// UnsafeMoriartyServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MoriartyServer will
// result in compilation errors.
type UnsafeMoriartyServer interface {
	mustEmbedUnimplementedMoriartyServer()
}

func RegisterMoriartyServer(s grpc.ServiceRegistrar, srv MoriartyServer) {
	s.RegisterService(&Moriarty_ServiceDesc, srv)
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
	in := new(App)
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
		return srv.(MoriartyServer).LaunchApp(ctx, req.(*App))
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
	in := new(App)
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
		return srv.(MoriartyServer).RemovePackage(ctx, req.(*App))
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
	in := new(App)
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
		return srv.(MoriartyServer).ConfigureGadget(ctx, req.(*App))
	}
	return interceptor(ctx, in, info, handler)
}

func _Moriarty_StartAutopilot_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AutoPilotConfig)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MoriartyServer).StartAutopilot(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/com.appknox.bakerstreet.Moriarty/StartAutopilot",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MoriartyServer).StartAutopilot(ctx, req.(*AutoPilotConfig))
	}
	return interceptor(ctx, in, info, handler)
}

func _Moriarty_Clean_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CleanOptions)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MoriartyServer).Clean(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/com.appknox.bakerstreet.Moriarty/Clean",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MoriartyServer).Clean(ctx, req.(*CleanOptions))
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

func _Moriarty_InfoV2_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MoriartyServer).InfoV2(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/com.appknox.bakerstreet.Moriarty/InfoV2",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MoriartyServer).InfoV2(ctx, req.(*Empty))
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

// Moriarty_ServiceDesc is the grpc.ServiceDesc for Moriarty service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Moriarty_ServiceDesc = grpc.ServiceDesc{
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
			MethodName: "StartAutopilot",
			Handler:    _Moriarty_StartAutopilot_Handler,
		},
		{
			MethodName: "Clean",
			Handler:    _Moriarty_Clean_Handler,
		},
		{
			MethodName: "Info",
			Handler:    _Moriarty_Info_Handler,
		},
		{
			MethodName: "InfoV2",
			Handler:    _Moriarty_InfoV2_Handler,
		},
		{
			MethodName: "ListPackages",
			Handler:    _Moriarty_ListPackages_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "bakerstreet/bakerstreet.proto",
}
