package main

import (
	"log"
	"net"
	"net/rpc"
	"pack_rpc/handler"
	"pack_rpc/server_proxy"
)

func main() {
	// 1.向rpc中注册服务
	rect := new(handler.RectService)
	err := server_proxy.RegisterRectService(rect)
	if err != nil {
		log.Println(err)
	}

	//3.监听服务
	listener, err := net.Listen("tcp", ":8000")
	if err != nil {
		log.Panicln(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Panicln(err)
		}
		rpc.ServeConn(conn)
	}

}
