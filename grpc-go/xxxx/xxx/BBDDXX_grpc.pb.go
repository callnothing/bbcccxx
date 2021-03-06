// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.5.0
// source: BBDDXX.proto

package xxx

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

// BBDDXXClient is the client API for BBDDXX service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BBDDXXClient interface {
	SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error)
}

type bBDDXXClient struct {
	cc grpc.ClientConnInterface
}

func NewBBDDXXClient(cc grpc.ClientConnInterface) BBDDXXClient {
	return &bBDDXXClient{cc}
}

func (c *bBDDXXClient) SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error) {
	out := new(HelloReply)
	err := c.cc.Invoke(ctx, "/xxxx.xxx.BBDDXX/SayHello", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BBDDXXServer is the server API for BBDDXX service.
// All implementations must embed UnimplementedBBDDXXServer
// for forward compatibility
type BBDDXXServer interface {
	SayHello(context.Context, *HelloRequest) (*HelloReply, error)
	mustEmbedUnimplementedBBDDXXServer()
}

// UnimplementedBBDDXXServer must be embedded to have forward compatible implementations.
type UnimplementedBBDDXXServer struct {
}

func (UnimplementedBBDDXXServer) SayHello(context.Context, *HelloRequest) (*HelloReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SayHello not implemented")
}
func (UnimplementedBBDDXXServer) mustEmbedUnimplementedBBDDXXServer() {}

// UnsafeBBDDXXServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BBDDXXServer will
// result in compilation errors.
type UnsafeBBDDXXServer interface {
	mustEmbedUnimplementedBBDDXXServer()
}

func RegisterBBDDXXServer(s grpc.ServiceRegistrar, srv BBDDXXServer) {
	s.RegisterService(&BBDDXX_ServiceDesc, srv)
}

func _BBDDXX_SayHello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BBDDXXServer).SayHello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/xxxx.xxx.BBDDXX/SayHello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BBDDXXServer).SayHello(ctx, req.(*HelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// BBDDXX_ServiceDesc is the grpc.ServiceDesc for BBDDXX service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BBDDXX_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "xxxx.xxx.BBDDXX",
	HandlerType: (*BBDDXXServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayHello",
			Handler:    _BBDDXX_SayHello_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "BBDDXX.proto",
}
