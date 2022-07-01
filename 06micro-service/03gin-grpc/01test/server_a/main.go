package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	pb "server_a/pb/protos"
	"strconv"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

func main() {

	// 1. grpc客户端 初始化
	conn, err := grpc.Dial("127.0.0.1:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err.Error())
	}
	defer conn.Close()

	// 2.gin 初始化
	r := gin.Default()
	// 测试一个get请求
	r.GET("/user/:id", func(c *gin.Context) {

		idStr := c.Param("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			fmt.Println(err)
		}

		// 获取client,发送id到server_b

		uc := pb.NewUserServiceClient(conn)
		res, err := uc.SendUserInfo(context.Background(), &pb.UserRequest{Id: int64(id)})
		if err != nil {
			fmt.Println(err)
		}

		c.JSON(http.StatusOK, gin.H{
			"user": res,
		})
	})
	// Run http server
	if err := r.Run(":8052"); err != nil {
		fmt.Printf("could not run server: %v", err)
	}

}
