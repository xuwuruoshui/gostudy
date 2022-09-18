package main

import (
	"02grpc/07metadata/proto/pb"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func main(){
	conn, err := grpc.Dial("127.0.0.1:9094",grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	client := pb.NewTodoServiceClient(conn)

	// 必须是ascii码
	//md1 := metadata.New(map[string]string{
	//	"name": "metadata try",
	//})
	md1 := metadata.Pairs("name","microserver","key","value")

	ctx := metadata.NewOutgoingContext(context.Background(), md1)
	resp, err := client.DoWork(ctx, &pb.TodoRequest{Msg: "看一下要学什么"})
	if err != nil {
		panic(err)
	}

	fmt.Println(resp.ResMsg)
}
