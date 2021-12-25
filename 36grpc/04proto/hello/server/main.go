package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	"gostudy/36grpc/04proto/proto/hello"
	"net"
)

const(
	Address = "127.0.0.1:5000"
)

type helloService struct {}

var HelloService = helloService{}

func(h helloService)SayHello(ctx context.Context,in *hello.HelloRequest)(*hello.HelloResponse,error){
	resp := new(hello.HelloResponse)
	resp.Message = fmt.Sprintf("Hello %s",in.Name)

	return resp,nil
}

func main(){

	// 1.开启监听端口
	listen,err := net.Listen("tcp",Address)
	if err!=nil{
		grpclog.Fatalln("Faild to listen:",err)
	}

	// 2.创建grpc服务端实例
	s := grpc.NewServer()

	// 3.注册HelloService原型到grpc服务端实例中
	hello.RegisterHelloServer(s,HelloService)

	// 需要先设置export(set) GRPC_GO_LOG_SEVERITY_LEVEL=INFO, 因为默认日志级别是Error
	// 再go run main.go
	grpclog.Infoln("Listen on",Address)

	// 4.绑定端口，开启服务
	s.Serve(listen)
}



