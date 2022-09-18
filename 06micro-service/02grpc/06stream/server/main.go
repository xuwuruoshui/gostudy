package main

import (
	"02grpc/06stream/proto/pb"
	"fmt"
	"google.golang.org/grpc"
	"net"
	"sync"
)

type FoodInfo struct {
	pb.UnimplementedFoodServiceServer
}

// 服务端
func (f *FoodInfo) SayName(req *pb.FoodStreamRequest,server pb.FoodService_SayNameServer) error {
	err := server.Send(&pb.FoodStreamResponse{Msg: "您点的菜已经做好了:" + req.Name})
	if err!=nil{
		panic(err)
	}
	return nil
}

// 客户端
func (f *FoodInfo) PostName(server pb.FoodService_PostNameServer) error {
	for{
		recv, err := server.Recv()
		if err!=nil{
			fmt.Println(err)
			return err
		}
		fmt.Println(recv.Name)
	}
}

// 双向通信
func (f *FoodInfo) FullStream(server pb.FoodService_FullStreamServer) error {

	var wg sync.WaitGroup

	foods := make(chan string,5)

	wg.Add(2)

	// 收
	go func() {
		defer wg.Done()
		for{
			recv, err := server.Recv()
			if err!=nil{
				fmt.Println(err)
				break
			}
			foods<-recv.Name
			fmt.Println(recv.Name)
		}
	}()

	// 发
	go func() {
		defer wg.Done()
		for{
			select {
			case food :=<-foods:
				err := server.Send(&pb.FoodStreamResponse{Msg: "先生你点的菜做好了:"+food})
				if err!=nil{
					fmt.Println(err)
					goto End
				}
			}
		}
		End:
	}()

	wg.Wait()
	return nil
}

func main() {

	server := grpc.NewServer()
	pb.RegisterFoodServiceServer(server,&FoodInfo{})
	listen, err := net.Listen("tcp", ":9091")
	if err!=nil{
		panic(err)
	}
	err = server.Serve(listen)
	if err!=nil{
		panic(err)
	}

}
