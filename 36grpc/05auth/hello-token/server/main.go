package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"gostudy/36grpc/05auth/proto/hello"
	"net"
)

const Address = "127.0.0.1:5000"

type helloService struct {}

var HelloService = helloService{}

func (h helloService)SayHello(ctx context.Context,in *hello.HelloRequest)(*hello.HelloResponse,error){
	// 1.解析数据
	md,ok := metadata.FromIncomingContext(ctx)
	if !ok{
		return nil,status.Errorf(codes.Unauthenticated,"token is null")
	}

	var(
		appid string
		appkey string
	)
	if val,ok := md["appid"];ok{
		appid = val[0]
	}

	if val,ok := md["appkey"];ok{
		appkey = val[0]
	}

	if appid !="101010" || appkey!= "i am key"{
		return nil,status.Errorf(codes.Unauthenticated,"invalid token,appid=%s,appkey=%s",appid,appkey)
	}

	resp := new(hello.HelloResponse)
	resp.Message = fmt.Sprintf("Hello %s.\n Token info:appid:%s,appkey:%s",in.Name,appid,appkey)
	return resp,nil
}
func main(){
	listen,err := net.Listen("tcp",Address)
	if err!=nil{
		grpclog.Fatalln("failed to listen:%v",err)
	}

	// TLS认证
	creds,err := credentials.NewServerTLSFromFile("../../keys/server.pem","../../keys/server.key")
	if err!=nil{
		grpclog.Fatalln("Failed to generate credentails:",err)
	}

	// 实例化grpc Server,并开启TLS认证
	s := grpc.NewServer(grpc.Creds(creds))

	// 注册HelloService
	hello.RegisterHelloServer(s,HelloService)

	grpclog.Infoln("Listen on",Address,"with TLS + Token")
	s.Serve(listen)
}
