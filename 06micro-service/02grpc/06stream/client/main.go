package main

import (
	"02grpc/06stream/proto/pb"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"sync"
	"time"
)

func main(){
	conn, err := grpc.Dial("localhost:9091", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	client := pb.NewFoodServiceClient(conn)
	// 1. 服务端流模式
	//ServerStream(err, hello)

	// 2. 客户端流模式
	//ClientStream(err, hello)

	// 3. 双向模式
	fullClient, err := client.FullStream(context.Background())
	if err!=nil{
		panic(err)
	}

	wg := FullStream(err, fullClient)

	wg.Wait()
}

func FullStream(err error, fullClient pb.FoodService_FullStreamClient) sync.WaitGroup {
	var wg sync.WaitGroup
	if err != nil {
		panic(err)
	}
	wg.Add(2)
	go func() {
		defer wg.Done()
		for {
			recv, err := fullClient.Recv()
			if err != nil {
				fmt.Println(err)
				break
			}
			fmt.Println(recv.Msg)
		}
	}()

	go func() {
		defer wg.Done()
		foods := []string{"东坡肘子", "回锅肉", "叫花鸡", "伤心凉粉", "担担面"}
		for _, item := range foods {
			err := fullClient.Send(&pb.FoodStreamRequest{Name: item})
			if err != nil {
				fmt.Println(err)
				break
			}
		}
	}()
	return wg
}

func ClientStream(err error, client pb.FoodServiceClient) {
	ctx, cancelFunc := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancelFunc()
	clientPost, err := client.PostName(ctx)
	if err != nil {
		panic(err)
	}

	foods := []string{"东坡肘子", "回锅肉", "叫花鸡", "伤心凉粉", "担担面"}

	for _, item := range foods {
		err = clientPost.Send(&pb.FoodStreamRequest{Name: item})
		time.Sleep(time.Second)
		if err != nil {
			fmt.Println(err)
			break
		}
	}
}

func ServerStream(err error, client pb.FoodServiceClient) {
	res, err := client.SayName(context.Background(), &pb.FoodStreamRequest{Name: "奥里给"})
	if err != nil {
		panic(err)
	}

	for {
		recv, err := res.Recv()
		if err != nil {
			fmt.Println(err)
			break
		}
		fmt.Println(recv.Msg)
	}
}
