package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc/metadata"
	"log"
	"net"

	"google.golang.org/grpc"

	helloworld "05metadata/proto"
)

// grpc metadata 类似http中header的token

type Server struct {
	helloworld.UnimplementedGreeterServer
}

func (s *Server) SayHello(ctx context.Context, request *helloworld.HelloRequest) (*helloworld.HelloReply, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		log.Println("get meatadata error")
	}
	for k, v := range md {
		log.Println(k, v)
	}
	return &helloworld.HelloReply{
		Message: fmt.Sprintf("Hello GRPC!!![%s]", request.GetName()),
	}, nil
}

func main() {
	g := grpc.NewServer()
	helloworld.RegisterGreeterServer(g, &Server{})
	listener, err := net.Listen("tcp", "127.0.0.1:9000")
	if err != nil {
		panic("faild to listen:" + err.Error())
	}
	err = g.Serve(listener)
	if err != nil {
		panic(fmt.Sprintf("faild to start grpc:%s", err))
	}
}
