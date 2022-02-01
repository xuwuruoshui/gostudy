package main

import (
	helloworld "08error/proto"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:8080", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	client := helloworld.NewGreeterClient(conn)
	hello, err := client.SayHello(context.Background(), &helloworld.HelloRequest{Name: "grpc client"})
	if err != nil {
		st, ok := status.FromError(err)
		if !ok {
			panic("解析错误！！！")
		}
		fmt.Println(st.Code())
		fmt.Println(st.Message())
		return
	}
	fmt.Println(hello.Message)
}
