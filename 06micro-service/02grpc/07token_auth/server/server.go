package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"log"
	"net"

	"google.golang.org/grpc"

	helloworld "07token_auth/proto"
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
	// 该函数实现UnaryServerInterceptor
	interceptor := func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {

		md, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			return resp, status.Error(codes.Unauthenticated, "无token信息,未认证")
		}
		var (
			appid  string
			appkey string
		)

		if k, ok := md["appid"]; ok {
			appid = k[0]
		}

		if k, ok := md["appkey"]; ok {
			appkey = k[0]
		}

		if appid != "2333" || appkey != "I am key" {
			return resp, status.Error(codes.Unauthenticated, "无token信息,未认证")
		}
		log.Printf("appid=%s,appkey=%s", appid, appkey)
		res, err := handler(ctx, req)

		return res, err
	}
	opt := grpc.UnaryInterceptor(interceptor)

	g := grpc.NewServer(opt)
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
