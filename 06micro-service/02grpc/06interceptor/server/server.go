package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"

	helloworld "06interceptor/proto"
)

// interceptor的实现

type Server struct {
	helloworld.UnimplementedGreeterServer
}

func (s *Server) SayHello(ctx context.Context, request *helloworld.HelloRequest) (*helloworld.HelloReply, error) {
	return &helloworld.HelloReply{
		Message: fmt.Sprintf("Hello GRPC!!![%s]", request.GetName()),
	}, nil
}

func main() {

	// 该函数实现UnaryServerInterceptor
	interceptor := func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		log.Println("grpc interceptor start!!! ")
		res, err := handler(ctx, req)
		log.Println("grpc interceptor end!!! ")
		return res, err
	}
	opt := grpc.UnaryInterceptor(interceptor)

	g := grpc.NewServer(opt)
	helloworld.RegisterGreeterServer(g, &Server{})
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic("faild to listen:" + err.Error())
	}
	err = g.Serve(listener)
	if err != nil {
		panic(fmt.Sprintf("faild to start grpc:%s", err))
	}
}
