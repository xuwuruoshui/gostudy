package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

// 参数
type Params struct {
	Width,Height int
}

// 矩形
type Rect struct{}

// RPC服务端方法，求矩形面积
func (r *Rect) Area(p Params, result *int) error {
	*result = p.Height * p.Width
	return nil
}

// RPC服务端方法，求矩形周长
func (r *Rect) Perimeter(p Params, result *int) error {
	*result = (p.Height + p.Width) * 2
	return nil
}

func main() {
	// 1.注册到rpc
	rpc.Register(new(Rect))
	listener, err := net.Listen("tcp", ":8000")
	if err != nil {
		log.Fatalln(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}

		go func(conn net.Conn) {
			fmt.Println("new hello", conn.RemoteAddr())
			jsonrpc.ServeConn(conn)
		}(conn)
	}

}
