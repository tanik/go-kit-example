// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: proto/v1/string.proto

package string

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

const (
	StringService_Uppercase_FullMethodName = "/proto.v1.StringService/Uppercase"
	StringService_Count_FullMethodName     = "/proto.v1.StringService/Count"
)

// StringServiceClient is the client API for StringService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type StringServiceClient interface {
	Uppercase(ctx context.Context, in *UppercaseRequest, opts ...grpc.CallOption) (*UppercaseResponse, error)
	Count(ctx context.Context, in *CountRequest, opts ...grpc.CallOption) (*CountResponse, error)
}

type stringServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewStringServiceClient(cc grpc.ClientConnInterface) StringServiceClient {
	return &stringServiceClient{cc}
}

func (c *stringServiceClient) Uppercase(ctx context.Context, in *UppercaseRequest, opts ...grpc.CallOption) (*UppercaseResponse, error) {
	out := new(UppercaseResponse)
	err := c.cc.Invoke(ctx, StringService_Uppercase_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *stringServiceClient) Count(ctx context.Context, in *CountRequest, opts ...grpc.CallOption) (*CountResponse, error) {
	out := new(CountResponse)
	err := c.cc.Invoke(ctx, StringService_Count_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StringServiceServer is the server API for StringService service.
// All implementations must embed UnimplementedStringServiceServer
// for forward compatibility
type StringServiceServer interface {
	Uppercase(context.Context, *UppercaseRequest) (*UppercaseResponse, error)
	Count(context.Context, *CountRequest) (*CountResponse, error)
	mustEmbedUnimplementedStringServiceServer()
}

// UnimplementedStringServiceServer must be embedded to have forward compatible implementations.
type UnimplementedStringServiceServer struct {
}

func (UnimplementedStringServiceServer) Uppercase(context.Context, *UppercaseRequest) (*UppercaseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Uppercase not implemented")
}
func (UnimplementedStringServiceServer) Count(context.Context, *CountRequest) (*CountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Count not implemented")
}
func (UnimplementedStringServiceServer) mustEmbedUnimplementedStringServiceServer() {}

// UnsafeStringServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to StringServiceServer will
// result in compilation errors.
type UnsafeStringServiceServer interface {
	mustEmbedUnimplementedStringServiceServer()
}

func RegisterStringServiceServer(s grpc.ServiceRegistrar, srv StringServiceServer) {
	s.RegisterService(&StringService_ServiceDesc, srv)
}

func _StringService_Uppercase_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UppercaseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StringServiceServer).Uppercase(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StringService_Uppercase_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StringServiceServer).Uppercase(ctx, req.(*UppercaseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StringService_Count_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StringServiceServer).Count(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StringService_Count_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StringServiceServer).Count(ctx, req.(*CountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// StringService_ServiceDesc is the grpc.ServiceDesc for StringService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var StringService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.v1.StringService",
	HandlerType: (*StringServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Uppercase",
			Handler:    _StringService_Uppercase_Handler,
		},
		{
			MethodName: "Count",
			Handler:    _StringService_Count_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/v1/string.proto",
}
