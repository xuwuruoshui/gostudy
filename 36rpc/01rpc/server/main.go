package main

import (
	"log"
	"net"
	"net/rpc"
)

// 参数
type Params struct {
	Width,Height int
}

// 矩形
type Rect struct {}

// RPC服务端方法，求矩形面积
func (r *Rect) Area(p Params,result *int)error{
	*result = p.Height * p.Width
	return nil
}

// RPC服务端方法，求矩形周长
func (r *Rect) Perimeter(p Params,result *int)error{
	*result = (p.Height + p.Width)*2
	return nil
}


func main(){
	// 1.向rpc中注册服务
	rect := new(Rect)
	rpc.Register(rect)
	// 自定义名字
	//rpc.RegisterName("haha",rect)


	//3.监听服务
	listener ,err := net.Listen("tcp",":8000")
	if err!=nil{
		log.Panicln(err)
	}

	for{
		conn,err := listener.Accept()
		if err!=nil{
			log.Panicln(err)
		}
		rpc.ServeConn(conn)
	}
	


}
