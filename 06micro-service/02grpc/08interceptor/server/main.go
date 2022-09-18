package main

import (
	"02grpc/08interceptor/proto/pb"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"net"
	"time"
)

type Todo struct {
	pb.UnimplementedTodoServiceServer
}

func (t Todo) DoWork(ctx context.Context,req *pb.TodoRequest) (*pb.TodoResponse, error) {

	msg := req.GetMsg()
	fmt.Println(msg)
	return &pb.TodoResponse{ResMsg: "hahaha"},nil
}

func MyInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error){
	t1 := time.Now()
	resp, err = handler(ctx, req)
	take := time.Now().Sub(t1)
	fmt.Printf("执行时间:%d",take.Milliseconds())
	return
}

func main(){

	interceptor := grpc.UnaryInterceptor(MyInterceptor)
	server := grpc.NewServer(interceptor)

	pb.RegisterTodoServiceServer(server,&Todo{})

	listen, err := net.Listen("tcp", ":9095")
	if err != nil {
		panic(err)
	}
	err = server.Serve(listen)
	if err != nil {
		panic(err)
	}

}
