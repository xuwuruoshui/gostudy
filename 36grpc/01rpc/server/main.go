package main

import (
	"log"
	"net/http"
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

	// 2.服务处理绑定到http协议上
	rpc.HandleHTTP()

	//3.监听服务
	err := http.ListenAndServe(":8000",nil)
	if err!=nil{
		log.Panicln(err)
	}
}
