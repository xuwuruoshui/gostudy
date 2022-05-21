package main

import (
	helloworld "06interceptor/proto"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"time"
)

// interceptor的实现

func main() {

	interceptor := func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		start := time.Now()
		err := invoker(ctx, method, req, reply, cc, opts...)
		log.Printf("耗时%s\n", time.Since(start))
		return err
	}

	opt := grpc.WithUnaryInterceptor(interceptor)

	conn, err := grpc.Dial("127.0.0.1:8080", grpc.WithInsecure(), opt)
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
