// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.14.0
// source: application/application.proto

package application

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

// ApplicationClient is the client API for Application service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ApplicationClient interface {
	Save(ctx context.Context, in *SaveRequest, opts ...grpc.CallOption) (*SaveResponse, error)
	Edit(ctx context.Context, in *EditRequest, opts ...grpc.CallOption) (*EditResponse, error)
	Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error)
}

type applicationClient struct {
	cc grpc.ClientConnInterface
}

func NewApplicationClient(cc grpc.ClientConnInterface) ApplicationClient {
	return &applicationClient{cc}
}

func (c *applicationClient) Save(ctx context.Context, in *SaveRequest, opts ...grpc.CallOption) (*SaveResponse, error) {
	out := new(SaveResponse)
	err := c.cc.Invoke(ctx, "/application.Application/Save", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *applicationClient) Edit(ctx context.Context, in *EditRequest, opts ...grpc.CallOption) (*EditResponse, error) {
	out := new(EditResponse)
	err := c.cc.Invoke(ctx, "/application.Application/Edit", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *applicationClient) Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error) {
	out := new(GetResponse)
	err := c.cc.Invoke(ctx, "/application.Application/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ApplicationServer is the server API for Application service.
// All implementations must embed UnimplementedApplicationServer
// for forward compatibility
type ApplicationServer interface {
	Save(context.Context, *SaveRequest) (*SaveResponse, error)
	Edit(context.Context, *EditRequest) (*EditResponse, error)
	Get(context.Context, *GetRequest) (*GetResponse, error)
	mustEmbedUnimplementedApplicationServer()
}

// UnimplementedApplicationServer must be embedded to have forward compatible implementations.
type UnimplementedApplicationServer struct {
}

func (UnimplementedApplicationServer) Save(context.Context, *SaveRequest) (*SaveResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Save not implemented")
}
func (UnimplementedApplicationServer) Edit(context.Context, *EditRequest) (*EditResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Edit not implemented")
}
func (UnimplementedApplicationServer) Get(context.Context, *GetRequest) (*GetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedApplicationServer) mustEmbedUnimplementedApplicationServer() {}

// UnsafeApplicationServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ApplicationServer will
// result in compilation errors.
type UnsafeApplicationServer interface {
	mustEmbedUnimplementedApplicationServer()
}

func RegisterApplicationServer(s grpc.ServiceRegistrar, srv ApplicationServer) {
	s.RegisterService(&Application_ServiceDesc, srv)
}

func _Application_Save_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SaveRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApplicationServer).Save(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/application.Application/Save",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApplicationServer).Save(ctx, req.(*SaveRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Application_Edit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EditRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApplicationServer).Edit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/application.Application/Edit",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApplicationServer).Edit(ctx, req.(*EditRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Application_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApplicationServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/application.Application/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApplicationServer).Get(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Application_ServiceDesc is the grpc.ServiceDesc for Application service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Application_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "application.Application",
	HandlerType: (*ApplicationServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Save",
			Handler:    _Application_Save_Handler,
		},
		{
			MethodName: "Edit",
			Handler:    _Application_Edit_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _Application_Get_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "application/application.proto",
}