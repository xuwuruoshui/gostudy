package main

import (
	helloworld "05metadata/proto"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

// grpc metadata 类似http中header的token
// tcp地址写全

func main() {
	conn, err := grpc.Dial("127.0.0.1:9000", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	client := helloworld.NewGreeterClient(conn)

	// metadata使用
	// 方法1
	//metadata.Pairs("time", time.Now().String())
	// 方法2
	md := metadata.New(map[string]string{
		"name":     "zhangsan",
		"password": "123456",
	})
	ctx := metadata.NewOutgoingContext(context.Background(), md)

	hello, err := client.SayHello(ctx, &helloworld.HelloRequest{Name: "grpc client"})
	if err != nil {
		panic(err)
	}
	fmt.Println(hello.Message)
}
