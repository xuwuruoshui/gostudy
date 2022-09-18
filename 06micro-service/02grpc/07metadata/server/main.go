package main

import (
	"02grpc/07metadata/proto/pb"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"net"
)

type Todo struct {
	pb.UnimplementedTodoServiceServer
}

func(t *Todo) DoWork(ctx context.Context,req *pb.TodoRequest) (*pb.TodoResponse, error){

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok{
		fmt.Println("metadata不存在")
		return nil, nil
	}
	for k,v := range md {
		fmt.Printf("%s:%s\n",k,v)
	}

	fmt.Println("服务端已接受到客户端消息: ",req.Msg)
	return &pb.TodoResponse{ResMsg: "你要学习: go微服务,grpc"},nil
}

func main(){




	server := grpc.NewServer()
	pb.RegisterTodoServiceServer(server,&Todo{})

	listen, err := net.Listen("tcp", ":9094")
	if err!=nil{
		panic(err)
	}
	err = server.Serve(listen)
	if err!=nil{
		panic(err)
	}
}
