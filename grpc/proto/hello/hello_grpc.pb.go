// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.1
// source: proto/hello.proto

package hello

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	Emaple_ServerReply_FullMethodName = "/Emaple/ServerReply"
)

// EmapleClient is the client API for Emaple service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type EmapleClient interface {
	ServerReply(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloResponse, error)
}

type emapleClient struct {
	cc grpc.ClientConnInterface
}

func NewEmapleClient(cc grpc.ClientConnInterface) EmapleClient {
	return &emapleClient{cc}
}

func (c *emapleClient) ServerReply(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(HelloResponse)
	err := c.cc.Invoke(ctx, Emaple_ServerReply_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EmapleServer is the server API for Emaple service.
// All implementations must embed UnimplementedEmapleServer
// for forward compatibility.
type EmapleServer interface {
	ServerReply(context.Context, *HelloRequest) (*HelloResponse, error)
	mustEmbedUnimplementedEmapleServer()
}

// UnimplementedEmapleServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedEmapleServer struct{}

func (UnimplementedEmapleServer) ServerReply(context.Context, *HelloRequest) (*HelloResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ServerReply not implemented")
}
func (UnimplementedEmapleServer) mustEmbedUnimplementedEmapleServer() {}
func (UnimplementedEmapleServer) testEmbeddedByValue()                {}

// UnsafeEmapleServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to EmapleServer will
// result in compilation errors.
type UnsafeEmapleServer interface {
	mustEmbedUnimplementedEmapleServer()
}

func RegisterEmapleServer(s grpc.ServiceRegistrar, srv EmapleServer) {
	// If the following call pancis, it indicates UnimplementedEmapleServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Emaple_ServiceDesc, srv)
}

func _Emaple_ServerReply_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EmapleServer).ServerReply(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Emaple_ServerReply_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EmapleServer).ServerReply(ctx, req.(*HelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Emaple_ServiceDesc is the grpc.ServiceDesc for Emaple service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Emaple_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Emaple",
	HandlerType: (*EmapleServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ServerReply",
			Handler:    _Emaple_ServerReply_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/hello.proto",
}
