// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.25.1
// source: abc.proto

package abc

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

// AbcClient is the client API for Abc service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AbcClient interface {
	Ping(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
}

type abcClient struct {
	cc grpc.ClientConnInterface
}

func NewAbcClient(cc grpc.ClientConnInterface) AbcClient {
	return &abcClient{cc}
}

func (c *abcClient) Ping(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/abc.Abc/Ping", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AbcServer is the server API for Abc service.
// All implementations must embed UnimplementedAbcServer
// for forward compatibility
type AbcServer interface {
	Ping(context.Context, *Request) (*Response, error)
	mustEmbedUnimplementedAbcServer()
}

// UnimplementedAbcServer must be embedded to have forward compatible implementations.
type UnimplementedAbcServer struct {
}

func (UnimplementedAbcServer) Ping(context.Context, *Request) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}
func (UnimplementedAbcServer) mustEmbedUnimplementedAbcServer() {}

// UnsafeAbcServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AbcServer will
// result in compilation errors.
type UnsafeAbcServer interface {
	mustEmbedUnimplementedAbcServer()
}

func RegisterAbcServer(s grpc.ServiceRegistrar, srv AbcServer) {
	s.RegisterService(&Abc_ServiceDesc, srv)
}

func _Abc_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AbcServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/abc.Abc/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AbcServer).Ping(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

// Abc_ServiceDesc is the grpc.ServiceDesc for Abc service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Abc_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "abc.Abc",
	HandlerType: (*AbcServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _Abc_Ping_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "abc.proto",
}
