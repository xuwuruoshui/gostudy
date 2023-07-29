// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.6
// source: machine.proto

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// StatusServiceClient is the hello API for StatusService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type StatusServiceClient interface {
	Alive(ctx context.Context, in *timestamppb.Timestamp, opts ...grpc.CallOption) (*Response, error)
}

type statusServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewStatusServiceClient(cc grpc.ClientConnInterface) StatusServiceClient {
	return &statusServiceClient{cc}
}

func (c *statusServiceClient) Alive(ctx context.Context, in *timestamppb.Timestamp, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/StatusService/alive", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StatusServiceServer is the server API for StatusService service.
// All implementations must embed UnimplementedStatusServiceServer
// for forward compatibility
type StatusServiceServer interface {
	Alive(context.Context, *timestamppb.Timestamp) (*Response, error)
	mustEmbedUnimplementedStatusServiceServer()
}

// UnimplementedStatusServiceServer must be embedded to have forward compatible implementations.
type UnimplementedStatusServiceServer struct {
}

func (UnimplementedStatusServiceServer) Alive(context.Context, *timestamppb.Timestamp) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Alive not implemented")
}
func (UnimplementedStatusServiceServer) mustEmbedUnimplementedStatusServiceServer() {}

// UnsafeStatusServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to StatusServiceServer will
// result in compilation errors.
type UnsafeStatusServiceServer interface {
	mustEmbedUnimplementedStatusServiceServer()
}

func RegisterStatusServiceServer(s grpc.ServiceRegistrar, srv StatusServiceServer) {
	s.RegisterService(&StatusService_ServiceDesc, srv)
}

func _StatusService_Alive_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(timestamppb.Timestamp)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StatusServiceServer).Alive(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/StatusService/alive",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StatusServiceServer).Alive(ctx, req.(*timestamppb.Timestamp))
	}
	return interceptor(ctx, in, info, handler)
}

// StatusService_ServiceDesc is the grpc.ServiceDesc for StatusService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var StatusService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "StatusService",
	HandlerType: (*StatusServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "alive",
			Handler:    _StatusService_Alive_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "machine.proto",
}
