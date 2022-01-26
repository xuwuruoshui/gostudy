package main

import (
	helloworld "01grpc/proto"
	"context"
	"fmt"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial(":8080", grpc.WithInsecure())
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
