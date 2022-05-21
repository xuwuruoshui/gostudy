package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"net"

	"google.golang.org/grpc"

	helloworld "08error/proto"
)

type Server struct {
	helloworld.UnimplementedGreeterServer
}

func (s *Server) SayHello(ctx context.Context, request *helloworld.HelloRequest) (*helloworld.HelloReply, error) {
	return &helloworld.HelloReply{
		Message: fmt.Sprintf("Hello GRPC!!![%s]", request.GetName()),
	}, status.Error(codes.NotFound, "can't find "+request.GetName())
}

func main() {
	g := grpc.NewServer()
	helloworld.RegisterGreeterServer(g, &Server{})
	listener, err := net.Listen("tcp", "127.0.0.1:8080")
	if err != nil {
		panic("faild to listen:" + err.Error())
	}
	err = g.Serve(listener)
	if err != nil {
		panic(fmt.Sprintf("faild to start grpc:%s", err))
	}
}
