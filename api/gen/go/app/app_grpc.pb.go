// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: app/app.proto

package app

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

// AppServiceClient is the client API for AppService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AppServiceClient interface {
	List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListResponse, error)
	Install(ctx context.Context, in *AppOpRequest, opts ...grpc.CallOption) (*AppOpReply, error)
	Uninstall(ctx context.Context, in *AppOpRequest, opts ...grpc.CallOption) (*AppOpReply, error)
	Start(ctx context.Context, in *AppStatusRequest, opts ...grpc.CallOption) (*AppStatusReply, error)
	Stop(ctx context.Context, in *AppStatusRequest, opts ...grpc.CallOption) (*AppStatusReply, error)
}

type appServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAppServiceClient(cc grpc.ClientConnInterface) AppServiceClient {
	return &appServiceClient{cc}
}

func (c *appServiceClient) List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListResponse, error) {
	out := new(ListResponse)
	err := c.cc.Invoke(ctx, "/app.AppService/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appServiceClient) Install(ctx context.Context, in *AppOpRequest, opts ...grpc.CallOption) (*AppOpReply, error) {
	out := new(AppOpReply)
	err := c.cc.Invoke(ctx, "/app.AppService/Install", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appServiceClient) Uninstall(ctx context.Context, in *AppOpRequest, opts ...grpc.CallOption) (*AppOpReply, error) {
	out := new(AppOpReply)
	err := c.cc.Invoke(ctx, "/app.AppService/Uninstall", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appServiceClient) Start(ctx context.Context, in *AppStatusRequest, opts ...grpc.CallOption) (*AppStatusReply, error) {
	out := new(AppStatusReply)
	err := c.cc.Invoke(ctx, "/app.AppService/Start", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appServiceClient) Stop(ctx context.Context, in *AppStatusRequest, opts ...grpc.CallOption) (*AppStatusReply, error) {
	out := new(AppStatusReply)
	err := c.cc.Invoke(ctx, "/app.AppService/Stop", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AppServiceServer is the server API for AppService service.
// All implementations must embed UnimplementedAppServiceServer
// for forward compatibility
type AppServiceServer interface {
	List(context.Context, *ListRequest) (*ListResponse, error)
	Install(context.Context, *AppOpRequest) (*AppOpReply, error)
	Uninstall(context.Context, *AppOpRequest) (*AppOpReply, error)
	Start(context.Context, *AppStatusRequest) (*AppStatusReply, error)
	Stop(context.Context, *AppStatusRequest) (*AppStatusReply, error)
	mustEmbedUnimplementedAppServiceServer()
}

// UnimplementedAppServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAppServiceServer struct {
}

func (UnimplementedAppServiceServer) List(context.Context, *ListRequest) (*ListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedAppServiceServer) Install(context.Context, *AppOpRequest) (*AppOpReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Install not implemented")
}
func (UnimplementedAppServiceServer) Uninstall(context.Context, *AppOpRequest) (*AppOpReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Uninstall not implemented")
}
func (UnimplementedAppServiceServer) Start(context.Context, *AppStatusRequest) (*AppStatusReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Start not implemented")
}
func (UnimplementedAppServiceServer) Stop(context.Context, *AppStatusRequest) (*AppStatusReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Stop not implemented")
}
func (UnimplementedAppServiceServer) mustEmbedUnimplementedAppServiceServer() {}

// UnsafeAppServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AppServiceServer will
// result in compilation errors.
type UnsafeAppServiceServer interface {
	mustEmbedUnimplementedAppServiceServer()
}

func RegisterAppServiceServer(s grpc.ServiceRegistrar, srv AppServiceServer) {
	s.RegisterService(&AppService_ServiceDesc, srv)
}

func _AppService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/app.AppService/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppServiceServer).List(ctx, req.(*ListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AppService_Install_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AppOpRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppServiceServer).Install(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/app.AppService/Install",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppServiceServer).Install(ctx, req.(*AppOpRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AppService_Uninstall_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AppOpRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppServiceServer).Uninstall(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/app.AppService/Uninstall",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppServiceServer).Uninstall(ctx, req.(*AppOpRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AppService_Start_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AppStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppServiceServer).Start(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/app.AppService/Start",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppServiceServer).Start(ctx, req.(*AppStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AppService_Stop_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AppStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppServiceServer).Stop(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/app.AppService/Stop",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppServiceServer).Stop(ctx, req.(*AppStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AppService_ServiceDesc is the grpc.ServiceDesc for AppService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AppService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "app.AppService",
	HandlerType: (*AppServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "List",
			Handler:    _AppService_List_Handler,
		},
		{
			MethodName: "Install",
			Handler:    _AppService_Install_Handler,
		},
		{
			MethodName: "Uninstall",
			Handler:    _AppService_Uninstall_Handler,
		},
		{
			MethodName: "Start",
			Handler:    _AppService_Start_Handler,
		},
		{
			MethodName: "Stop",
			Handler:    _AppService_Stop_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "app/app.proto",
}
