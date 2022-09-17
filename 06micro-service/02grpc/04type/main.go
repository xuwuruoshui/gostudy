package main

import (
	"02grpc/04type/proto/pb"
	"fmt"
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"
)

func main(){
	res := &pb.TodoRequest{
		Todo: "test",
		Week: pb.Week_Friday,
		BookMap: map[string]string{"a":"artist"},
		DoneTime: timestamppb.New(time.Now()),
	}


	fmt.Println(res)
}
