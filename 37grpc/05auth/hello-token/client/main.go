package main

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/grpclog"
	"gostudy/36grpc/05auth/proto/hello"
)

const (
	Address = "127.0.0.1:5000"

	// 是否开启TLS认证
	OpenTLS = true
)

type customCredential struct {

}

// 客制化认证数据
func (c customCredential)GetRequestMetadata(ctx context.Context,uri ...string)(cred map[string]string,err error){
	cred = map[string]string{
		"appid": "101010",
		"appkey": "i am key",
	}
	return
}

// 是否开启TLS
func (c customCredential)RequireTransportSecurity() bool{
	return OpenTLS
}

func main(){
	var err error
	// RPC认证方式
	var opts []grpc.DialOption

	if OpenTLS{
		creds,err := credentials.NewClientTLSFromFile("../../keys/server.pem","www.haha.com")
		if err!=nil{
			grpclog.Fatalln("Failed to create TLS credential: ",err)
		}
		opts = append(opts,grpc.WithTransportCredentials(creds))
	}else {
		opts = append(opts,grpc.WithInsecure())
	}

	// 使用自定义认证
	opts = append(opts,grpc.WithPerRPCCredentials(new (customCredential)))

	// 1.连接服务端，并设置上认证方式
	conn,err := grpc.Dial(Address,opts...)

	if err!=nil{
		grpclog.Fatalln("Fail to connet Server:",err)
	}

	defer conn.Close()

	// 2.初始化客户端
	c := hello.NewHelloClient(conn)

	// 3.封装数据并调用方法
	req := &hello.HelloRequest{Name: "golang RPC"}
	res,err := c.SayHello(context.Background(),req)
	if err!=nil{
		grpclog.Fatalln("Failed to call remote:",err)
	}
	grpclog.Infoln(res.Message)
}
