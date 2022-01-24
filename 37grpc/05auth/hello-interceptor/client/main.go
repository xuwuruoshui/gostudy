package main

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/grpclog"
	"gostudy/36grpc/05auth/proto/hello"
	"time"
)

const (
	Address = "127.0.0.1:5000"
	OpenTLS = true
)

type customCredential struct {}

func (c customCredential)GetRequestMetadata(ctx context.Context,uri ...string)(map[string]string,error){
	return map[string]string{
		"appid":"101010",
		"appkey":"i am key",
	},nil
}

// RequireTransportSecurity 自定义认证是否开启TLS
func (c customCredential)RequireTransportSecurity()bool{
	return OpenTLS
}

func main(){
	var err error
	var opts []grpc.DialOption

	if OpenTLS{
		// TLS连接
		creds,err := credentials.NewClientTLSFromFile("../../keys/server.pem","www.haha.com")
		if err!=nil{
			grpclog.Fatalln("Faild to create TLS credentials:",err)
		}
		opts = append(opts,grpc.WithTransportCredentials(creds))
	}else{
		opts = append(opts,grpc.WithInsecure())
	}

	// 指定自定义认证
	opts = append(opts,grpc.WithPerRPCCredentials(new(customCredential)))
	// 指定客户端interceptor
	opts = append(opts,grpc.WithUnaryInterceptor(interceport))

	conn,err := grpc.Dial(Address,opts...)
	if err != nil {
		grpclog.Fatalln(err)
	}
	defer conn.Close()

	// 初始化客户端
	c := hello.NewHelloClient(conn)

	// 调用方法
	req := &hello.HelloRequest{Name: "gRPC Interceptor",}
	res,err := c.SayHello(context.Background(),req)
	if err!=nil{
		grpclog.Fatalln(err)
	}
	grpclog.Infoln(res.Message)
}

// 客户端接受服务端拦截器
func interceport(ctx context.Context,method string,req,reply interface{},conn *grpc.ClientConn,invoker grpc.UnaryInvoker,opts ...grpc.CallOption)(error){
	start := time.Now()
	err := invoker(ctx,method,req,reply,conn,opts...)
	grpclog.Infof("method=%s req=%v rep=%v duration=%s error=%v\n", method, req, reply, time.Since(start), err)
	return err
}