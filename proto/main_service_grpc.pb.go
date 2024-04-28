// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v5.26.1
// source: proto/main_service.proto

package proto

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

// MainServiceClient is the client API for MainService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MainServiceClient interface {
	SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloResponse, error)
}

type mainServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMainServiceClient(cc grpc.ClientConnInterface) MainServiceClient {
	return &mainServiceClient{cc}
}

func (c *mainServiceClient) SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloResponse, error) {
	out := new(HelloResponse)
	err := c.cc.Invoke(ctx, "/autopal.MainService/SayHello", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MainServiceServer is the server API for MainService service.
// All implementations must embed UnimplementedMainServiceServer
// for forward compatibility
type MainServiceServer interface {
	SayHello(context.Context, *HelloRequest) (*HelloResponse, error)
	mustEmbedUnimplementedMainServiceServer()
}

// UnimplementedMainServiceServer must be embedded to have forward compatible implementations.
type UnimplementedMainServiceServer struct {
}

func (UnimplementedMainServiceServer) SayHello(context.Context, *HelloRequest) (*HelloResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SayHello not implemented")
}
func (UnimplementedMainServiceServer) mustEmbedUnimplementedMainServiceServer() {}

// UnsafeMainServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MainServiceServer will
// result in compilation errors.
type UnsafeMainServiceServer interface {
	mustEmbedUnimplementedMainServiceServer()
}

func RegisterMainServiceServer(s grpc.ServiceRegistrar, srv MainServiceServer) {
	s.RegisterService(&MainService_ServiceDesc, srv)
}

func _MainService_SayHello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MainServiceServer).SayHello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/autopal.MainService/SayHello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MainServiceServer).SayHello(ctx, req.(*HelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// MainService_ServiceDesc is the grpc.ServiceDesc for MainService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MainService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "autopal.MainService",
	HandlerType: (*MainServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayHello",
			Handler:    _MainService_SayHello_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/main_service.proto",
}
