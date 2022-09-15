package main

import (
	"02grpc/01protobuf/proto/pb"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"net"
)

type BookInfo struct {
	pb.UnimplementedStudyServer
}

func (b *BookInfo) Study(ct context.Context, req *pb.BookRequest) (*pb.BookResponse, error) {
	fmt.Println(req.Name)
	return &pb.BookResponse{Msg: "Server: Welcome to GRPC!!!"}, nil
}

func main() {

	// 1.创建服务
	server := grpc.NewServer()
	pb.RegisterStudyServer(server, &BookInfo{})
	listen, err := net.Listen("tcp", "0.0.0.0:9090")
	if err != nil {
		panic(err)
	}
	err = server.Serve(listen)
	if err != nil {
		panic(err)
	}
}
