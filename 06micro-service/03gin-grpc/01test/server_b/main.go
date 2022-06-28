package main

import (
	"context"
	"fmt"
	"net"
	"net/http"
	pb "server_b/pb/protos"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

type Server struct {
	pb.UnimplementedUserServiceServer
}

func (s *Server) SendUserInfo(ctx context.Context, request *pb.UserRequest) (*pb.UserResponse, error) {
	return &pb.UserResponse{Id: request.Id, Name: "zhangsan", Age: 12}, nil
}

func main() {

	// 1. grpc服务端 初始化
	g := grpc.NewServer()
	pb.RegisterUserServiceServer(g, &Server{})
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		panic("faild to listen:" + err.Error())
	}
	err = g.Serve(listener)
	if err != nil {
		panic(fmt.Sprintf("faild to start grpc:%s", err))
	}

	// 2.gin 初始化
	r := gin.Default()
	// 测试一个get请求
	r.GET("/test", func(c *gin.Context) {

		c.JSON(http.StatusOK, gin.H{
			"test": "fffff",
		})
	})

	// Run http server
	if err := r.Run(":8053"); err != nil {
		fmt.Printf("could not run server: %v", err)
	}

}
