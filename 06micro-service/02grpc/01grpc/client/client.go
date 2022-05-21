package main

import (
	helloworld "01grpc/proto"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"time"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:8080", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	client := helloworld.NewGreeterClient(conn)
	//hello, err := client.SayHello(context.Background(), &helloworld.HelloRequest{Name: "grpc client"})
	// 超时
	ctx, _ := context.WithTimeout(context.Background(), time.Second*3)
	hello, err := client.SayHello(ctx, &helloworld.HelloRequest{Name: "grpc client"})
	if err != nil {
		panic(err)
	}
	fmt.Println(hello.Message)
}
