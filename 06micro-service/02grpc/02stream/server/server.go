package main

import (
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	stream "stream/proto"
	"sync"
	"time"
)

type Server struct {
	stream.UnimplementedGreeterServer
}

func (s *Server) GetStream(req *stream.StreamReqData, res stream.Greeter_GetStreamServer) error {
	i := 0
	for i < 10 {
		i++
		_ = res.Send(&stream.StreamResData{Data: time.Now().String()})
		time.Sleep(time.Second)
	}

	return nil
}
func (s *Server) PostStream(cli stream.Greeter_PostStreamServer) error {

	for {
		recv, err := cli.Recv()
		if err != nil {
			log.Println(err)
			break
		}
		log.Println(recv.Data)
	}

	return nil
}
func (s *Server) AllStream(all stream.Greeter_AllStreamServer) error {

	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		i := 0
		for i < 10 {
			i++
			all.Send(&stream.StreamResData{Data: "我是客户端"})
			time.Sleep(time.Second)
		}
	}()

	go func() {
		defer wg.Done()
		i := 0
		for i < 10 {
			recv, err := all.Recv()
			if err != nil {
				log.Println(err)
				break
			}
			log.Println("服务端消息:", recv.Data)
		}
	}()
	wg.Wait()
	return nil
}

func main() {
	g := grpc.NewServer()
	listen, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}
	stream.RegisterGreeterServer(g, &Server{})
	err = g.Serve(listen)
	if err != nil {
		panic(fmt.Sprintf("faild to start grpc:%s", err))
	}
}
