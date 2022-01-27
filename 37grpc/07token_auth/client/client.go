package main

import (
	helloworld "07token_auth/proto"
	"context"
	"fmt"
	"google.golang.org/grpc"
)

type customCredentials struct {
}

func (c customCredentials) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
	return map[string]string{
		"appid":  "2333",
		"appkey": "I am key",
	}, nil
}

func (c customCredentials) RequireTransportSecurity() bool {
	return false
}

func main() {
	// 方法1：interceptor+metadata实现认证
	//interceptor := func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	//	md := metadata.New(map[string]string{
	//		"appid":  "2333",
	//		"appkey": "I am key",
	//	})
	//	ctx = metadata.NewOutgoingContext(context.Background(), md)
	//	err := invoker(ctx, method, req, reply, cc, opts...)
	//	return err
	//}
	//opt := grpc.WithUnaryInterceptor(interceptor)

	// 方法2: 自带的,需要实现PerRPCCredentials
	credentials := grpc.WithPerRPCCredentials(&customCredentials{})

	conn, err := grpc.Dial("127.0.0.1:8080", grpc.WithInsecure(), credentials)
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	client := helloworld.NewGreeterClient(conn)
	hello, err := client.SayHello(context.Background(), &helloworld.HelloRequest{Name: "grpc client"})
	if err != nil {
		panic(err)
	}
	fmt.Println(hello.Message)
}
