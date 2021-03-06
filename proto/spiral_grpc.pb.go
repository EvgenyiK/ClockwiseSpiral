// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package spiral

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

// ClockWiseSpiralServiceClient is the client API for ClockWiseSpiralService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ClockWiseSpiralServiceClient interface {
	GenerateSpiral(ctx context.Context, in *NumbRequest, opts ...grpc.CallOption) (*NumbResponse, error)
}

type clockWiseSpiralServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewClockWiseSpiralServiceClient(cc grpc.ClientConnInterface) ClockWiseSpiralServiceClient {
	return &clockWiseSpiralServiceClient{cc}
}

func (c *clockWiseSpiralServiceClient) GenerateSpiral(ctx context.Context, in *NumbRequest, opts ...grpc.CallOption) (*NumbResponse, error) {
	out := new(NumbResponse)
	err := c.cc.Invoke(ctx, "/spiral.ClockWiseSpiralService/GenerateSpiral", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ClockWiseSpiralServiceServer is the server API for ClockWiseSpiralService service.
// All implementations must embed UnimplementedClockWiseSpiralServiceServer
// for forward compatibility
type ClockWiseSpiralServiceServer interface {
	GenerateSpiral(context.Context, *NumbRequest) (*NumbResponse, error)
	mustEmbedUnimplementedClockWiseSpiralServiceServer()
}

// UnimplementedClockWiseSpiralServiceServer must be embedded to have forward compatible implementations.
type UnimplementedClockWiseSpiralServiceServer struct {
}

func (UnimplementedClockWiseSpiralServiceServer) GenerateSpiral(context.Context, *NumbRequest) (*NumbResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GenerateSpiral not implemented")
}
func (UnimplementedClockWiseSpiralServiceServer) mustEmbedUnimplementedClockWiseSpiralServiceServer() {
}

// UnsafeClockWiseSpiralServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ClockWiseSpiralServiceServer will
// result in compilation errors.
type UnsafeClockWiseSpiralServiceServer interface {
	mustEmbedUnimplementedClockWiseSpiralServiceServer()
}

func RegisterClockWiseSpiralServiceServer(s grpc.ServiceRegistrar, srv ClockWiseSpiralServiceServer) {
	s.RegisterService(&ClockWiseSpiralService_ServiceDesc, srv)
}

func _ClockWiseSpiralService_GenerateSpiral_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NumbRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClockWiseSpiralServiceServer).GenerateSpiral(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/spiral.ClockWiseSpiralService/GenerateSpiral",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClockWiseSpiralServiceServer).GenerateSpiral(ctx, req.(*NumbRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ClockWiseSpiralService_ServiceDesc is the grpc.ServiceDesc for ClockWiseSpiralService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ClockWiseSpiralService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "spiral.ClockWiseSpiralService",
	HandlerType: (*ClockWiseSpiralServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GenerateSpiral",
			Handler:    _ClockWiseSpiralService_GenerateSpiral_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/spiral.proto",
}
