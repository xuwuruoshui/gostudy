package main

import (
	"01hello-go-micro/proto/hellogomicro"
	"context"
	"go-micro.dev/v4"
	"log"
)

type Greeter struct {
}

func (h *Greeter) Hello(ctx context.Context, req *hellogomicro.Request, rsp *hellogomicro.Response) error {
	rsp.Msg = "Hello " + req.Name
	return nil
}

func main() {
	service := micro.NewService(micro.Name("go-micro-server"))
	service.Init()
	err := hellogomicro.RegisterGreeterHandler(service.Server(), new(Greeter))
	if err != nil {
		log.Println("register server error:", err)
	}
	err = service.Run()
	if err != nil {
		log.Println(err)
	}
}
