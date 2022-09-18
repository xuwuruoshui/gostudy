package main

import (
	"02grpc/08interceptor/proto/pb"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"time"
)


func MyInterceptor(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error{
	t1 := time.Now()
	err := invoker(ctx, method, req, reply, cc, opts...)
	if err!=nil{
		panic(err)
	}
	take := time.Since(t1)
	fmt.Printf("客户端执行时间:%d\n",take.Milliseconds())
	return err
}

func main(){
	opt := grpc.WithUnaryInterceptor(MyInterceptor)
	conn, err := grpc.Dial("127.0.0.1:9095", grpc.WithInsecure(), opt)
	if err!=nil{
		panic(err)
	}
	defer conn.Close()

	client := pb.NewTodoServiceClient(conn)
	res, err := client.DoWork(context.Background(), &pb.TodoRequest{Msg: "I am so happy!!!"})
	if err != nil {
		panic(err)
	}
	fmt.Println(res.ResMsg)
}
