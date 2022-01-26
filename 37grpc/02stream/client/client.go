package main

import (
	"context"
	"google.golang.org/grpc"
	"log"
	stream "stream/proto"
	"sync"
	"time"
)

func main() {
	conn, err := grpc.Dial(":8080", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	// 服务端->客户端
	client := stream.NewGreeterClient(conn)
	res, _ := client.GetStream(context.Background(), &stream.StreamReqData{})

	for {
		recv, err := res.Recv()
		if err != nil {
			log.Println(err)
			break
		}
		log.Println(recv.Data)
	}

	// 客户端->服务端
	client1, _ := client.PostStream(context.Background())
	i := 0
	for i < 10 {
		i++
		_ = client1.Send(&stream.StreamReqData{Data: "客户端到此一游"})
		time.Sleep(time.Second)
	}

	// 客户端<->服务端
	client2, err := client.AllStream(context.Background())
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		i := 0
		for i < 10 {
			i++
			client2.Send(&stream.StreamReqData{Data: "我是服务端"})
			time.Sleep(time.Second)
		}
	}()

	go func() {
		defer wg.Done()
		i := 0
		for i < 10 {
			i++
			recv, err := client2.Recv()
			if err != nil {
				log.Println(err)
				break
			}
			log.Println("客户端消息:", recv.Data)
		}
	}()
	wg.Wait()
}
