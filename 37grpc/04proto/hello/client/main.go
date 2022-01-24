package main

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	"gostudy/36grpc/04proto/proto/hello"
)

const(
	Address = "127.0.0.1:5000"
)

func main(){
	// 1.连接rpc
	conn,err := grpc.Dial(Address,grpc.WithInsecure())
	if err!=nil{
		grpclog.Fatalln(err)
	}
	defer conn.Close()

	// 2.初始化客户端
	c := hello.NewHelloClient(conn)

	// 3.封装数据，调用rpc方法
	req := &hello.HelloRequest{Name: "gRPC"}
	res,err := c.SayHello(context.Background(),req)

	if err!=nil{
		grpclog.Fatalln(err)
	}

	grpclog.Infoln(res.Message)

}
