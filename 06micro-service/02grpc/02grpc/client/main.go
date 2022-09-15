package main

import (
	"02grpc/01protobuf/proto/pb"
	"context"
	"fmt"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:9090", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	client := pb.NewStudyClient(conn)
	resp, err := client.Study(context.Background(), &pb.BookRequest{Name: "Client: I want to study go micro"})
	if err != nil {
		panic(err)
	}
	fmt.Println(resp.Msg)
}
