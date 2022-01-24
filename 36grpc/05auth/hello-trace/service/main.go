package main

import (
	"context"
	"fmt"
	"golang.org/x/net/trace"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"gostudy/36grpc/05auth/proto/hello"
	"net"
	"net/http"
)

const(
	Address = "127.0.0.1:5000"
)

type helloService struct {
}

var HelloService = helloService{}

func (h helloService)SayHello(ctx context.Context,req *hello.HelloRequest)(*hello.HelloResponse,error){
	resp := new(hello.HelloResponse)
	resp.Message = fmt.Sprintf("Hello %s",req.Name)
	return resp,nil
}

func main(){
	listen,err := net.Listen("tcp",Address)
	if err!=nil{
		grpclog.Fatalln("Failed to listen:",err)
	}

	var opts []grpc.ServerOption

	// TLS认证
	creds,err := credentials.NewServerTLSFromFile("../../keys/server.pem","../../keys/server.key")
	if err!=nil{
		grpclog.Fatalln("Failed to generate credentials:",err)
	}

	opts = append(opts,grpc.Creds(creds))

	// 注册interceptor
	opts = append(opts,grpc.UnaryInterceptor(interceptor))

	// 实例化grpc Server
	s := grpc.NewServer(opts...)

	hello.RegisterHelloServer(s,HelloService)
	go startTrace()
	grpclog.Infoln("Listen on"+Address+"with TLS + Token + Interceptor")
	s.Serve(listen)
}

func startTrace(){
	trace.AuthRequest = func(req *http.Request) (any, sensitive bool) {
		return true,true
	}

	// 浏览器访问localhost:5001查看结果
	go http.ListenAndServe(":5001",nil)
	grpclog.Infoln("Trace listen on 5001")
}

func auth(ctx context.Context)error{
	md,ok := metadata.FromIncomingContext(ctx)
	if !ok{
		return status.Errorf(codes.Unauthenticated,"无Token认证信息")
	}

	var (
		appid string
		appkey string
	)

	if val,ok := md["appid"];ok{
		appid = val[0]
	}

	if val,ok := md["appkey"];ok{
		appkey = val[0]
	}

	if appid!="101010" || appkey !="i am key"{
		return status.Errorf(codes.Unauthenticated,"Token认证信息无效: appid=%s, appkey=%s", appid, appkey)
	}

	return nil
}

func interceptor(ctx context.Context,req interface{},info *grpc.UnaryServerInfo,handler grpc.UnaryHandler)(interface{},error){
	// 服务端经过token验证
	err := auth(ctx)
	if err!=nil{
		return nil,err
	}

	// 继续处理请求
	return handler(ctx,req)
}