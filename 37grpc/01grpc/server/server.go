package main

import (
	"context"
	"fmt"
	"net"

	"google.golang.org/grpc"

	helloworld "01grpc/proto"
)

type Server struct {
	helloworld.UnimplementedGreeterServer
}

func (s *Server) SayHello(ctx context.Context, request *helloworld.HelloRequest) (*helloworld.HelloReply, error) {
	return &helloworld.HelloReply{
		Message: fmt.Sprintf("Hello GRPC!!![%s]", request.GetName()),
	}, nil
}

func main() {
	g := grpc.NewServer()
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
