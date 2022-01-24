package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/grpclog"
	"gostudy/36grpc/05auth/proto/hello"
	"net"
)

const (
	Address = "127.0.0.1:5000"
)

// 定义Service
type helloService struct {

}

var HelloService = helloService{}

func (h helloService)SayHello(ctx context.Context, in *hello.HelloRequest) (*hello.HelloResponse,error){
	resp := new(hello.HelloResponse)
	resp.Message = fmt.Sprintf("Hello %s",in.Name)
	return resp,nil
}

func main(){
	// 1.开启监听端口
	listener,err := net.Listen("tcp",Address)
	if err!=nil{
		grpclog.Fatalln("Fail to listen:",err)
	}

	// 2. TLS认证
	creds,err := credentials.NewServerTLSFromFile("../../keys/server.pem","../../keys/server.key")
	if err!=nil{
		grpclog.Fatalln("Fail to generate TSL:",err)
	}
	// 2.创建grpc服务端实例
	s := grpc.NewServer(grpc.Creds(creds))

	// 3.注册HelloService原型到grpc服务端实例中
	hello.RegisterHelloServer(s,HelloService)

	// 4.绑定端口，开启服务
	s.Serve(listener)
}
