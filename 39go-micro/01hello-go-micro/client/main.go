package main

import (
	"01hello-go-micro/proto/hellogomicro"
	"context"
	"fmt"
	"go-micro.dev/v4"
)

func main() {
	service := micro.NewService(micro.Name("go-micro-client"))
	service.Init()
	server := hellogomicro.NewGreeterService("go-micro-server", service.Client())

	rsp, err := server.Hello(context.Background(), &hellogomicro.Request{Name: "ffff"})
	if err != nil {
		fmt.Println("call server failed,err:", err)
		return
	}
	fmt.Println(rsp.Msg)

}
