// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.6
// source: food.proto

package pb

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

// FoodServiceClient is the client API for FoodService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FoodServiceClient interface {
	SayName(ctx context.Context, in *FoodStreamRequest, opts ...grpc.CallOption) (FoodService_SayNameClient, error)
	PostName(ctx context.Context, opts ...grpc.CallOption) (FoodService_PostNameClient, error)
	FullStream(ctx context.Context, opts ...grpc.CallOption) (FoodService_FullStreamClient, error)
}

type foodServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewFoodServiceClient(cc grpc.ClientConnInterface) FoodServiceClient {
	return &foodServiceClient{cc}
}

func (c *foodServiceClient) SayName(ctx context.Context, in *FoodStreamRequest, opts ...grpc.CallOption) (FoodService_SayNameClient, error) {
	stream, err := c.cc.NewStream(ctx, &FoodService_ServiceDesc.Streams[0], "/FoodService/SayName", opts...)
	if err != nil {
		return nil, err
	}
	x := &foodServiceSayNameClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type FoodService_SayNameClient interface {
	Recv() (*FoodStreamResponse, error)
	grpc.ClientStream
}

type foodServiceSayNameClient struct {
	grpc.ClientStream
}

func (x *foodServiceSayNameClient) Recv() (*FoodStreamResponse, error) {
	m := new(FoodStreamResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *foodServiceClient) PostName(ctx context.Context, opts ...grpc.CallOption) (FoodService_PostNameClient, error) {
	stream, err := c.cc.NewStream(ctx, &FoodService_ServiceDesc.Streams[1], "/FoodService/PostName", opts...)
	if err != nil {
		return nil, err
	}
	x := &foodServicePostNameClient{stream}
	return x, nil
}

type FoodService_PostNameClient interface {
	Send(*FoodStreamRequest) error
	CloseAndRecv() (*FoodStreamResponse, error)
	grpc.ClientStream
}

type foodServicePostNameClient struct {
	grpc.ClientStream
}

func (x *foodServicePostNameClient) Send(m *FoodStreamRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *foodServicePostNameClient) CloseAndRecv() (*FoodStreamResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(FoodStreamResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *foodServiceClient) FullStream(ctx context.Context, opts ...grpc.CallOption) (FoodService_FullStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &FoodService_ServiceDesc.Streams[2], "/FoodService/FullStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &foodServiceFullStreamClient{stream}
	return x, nil
}

type FoodService_FullStreamClient interface {
	Send(*FoodStreamRequest) error
	Recv() (*FoodStreamResponse, error)
	grpc.ClientStream
}

type foodServiceFullStreamClient struct {
	grpc.ClientStream
}

func (x *foodServiceFullStreamClient) Send(m *FoodStreamRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *foodServiceFullStreamClient) Recv() (*FoodStreamResponse, error) {
	m := new(FoodStreamResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// FoodServiceServer is the server API for FoodService service.
// All implementations must embed UnimplementedFoodServiceServer
// for forward compatibility
type FoodServiceServer interface {
	SayName(*FoodStreamRequest, FoodService_SayNameServer) error
	PostName(FoodService_PostNameServer) error
	FullStream(FoodService_FullStreamServer) error
	mustEmbedUnimplementedFoodServiceServer()
}

// UnimplementedFoodServiceServer must be embedded to have forward compatible implementations.
type UnimplementedFoodServiceServer struct {
}

func (UnimplementedFoodServiceServer) SayName(*FoodStreamRequest, FoodService_SayNameServer) error {
	return status.Errorf(codes.Unimplemented, "method SayName not implemented")
}
func (UnimplementedFoodServiceServer) PostName(FoodService_PostNameServer) error {
	return status.Errorf(codes.Unimplemented, "method PostName not implemented")
}
func (UnimplementedFoodServiceServer) FullStream(FoodService_FullStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method FullStream not implemented")
}
func (UnimplementedFoodServiceServer) mustEmbedUnimplementedFoodServiceServer() {}

// UnsafeFoodServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FoodServiceServer will
// result in compilation errors.
type UnsafeFoodServiceServer interface {
	mustEmbedUnimplementedFoodServiceServer()
}

func RegisterFoodServiceServer(s grpc.ServiceRegistrar, srv FoodServiceServer) {
	s.RegisterService(&FoodService_ServiceDesc, srv)
}

func _FoodService_SayName_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(FoodStreamRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(FoodServiceServer).SayName(m, &foodServiceSayNameServer{stream})
}

type FoodService_SayNameServer interface {
	Send(*FoodStreamResponse) error
	grpc.ServerStream
}

type foodServiceSayNameServer struct {
	grpc.ServerStream
}

func (x *foodServiceSayNameServer) Send(m *FoodStreamResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _FoodService_PostName_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(FoodServiceServer).PostName(&foodServicePostNameServer{stream})
}

type FoodService_PostNameServer interface {
	SendAndClose(*FoodStreamResponse) error
	Recv() (*FoodStreamRequest, error)
	grpc.ServerStream
}

type foodServicePostNameServer struct {
	grpc.ServerStream
}

func (x *foodServicePostNameServer) SendAndClose(m *FoodStreamResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *foodServicePostNameServer) Recv() (*FoodStreamRequest, error) {
	m := new(FoodStreamRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _FoodService_FullStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(FoodServiceServer).FullStream(&foodServiceFullStreamServer{stream})
}

type FoodService_FullStreamServer interface {
	Send(*FoodStreamResponse) error
	Recv() (*FoodStreamRequest, error)
	grpc.ServerStream
}

type foodServiceFullStreamServer struct {
	grpc.ServerStream
}

func (x *foodServiceFullStreamServer) Send(m *FoodStreamResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *foodServiceFullStreamServer) Recv() (*FoodStreamRequest, error) {
	m := new(FoodStreamRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// FoodService_ServiceDesc is the grpc.ServiceDesc for FoodService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var FoodService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "FoodService",
	HandlerType: (*FoodServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "SayName",
			Handler:       _FoodService_SayName_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "PostName",
			Handler:       _FoodService_PostName_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "FullStream",
			Handler:       _FoodService_FullStream_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "food.proto",
}