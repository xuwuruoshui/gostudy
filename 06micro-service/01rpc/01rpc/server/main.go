package main

import (
	"fmt"
	"net"
	"net/rpc"
)

type FoodService struct {

}

func (f *FoodService)SayName(req string,resp *string)error{

	*resp = "您点的菜是:"+req
	return nil
}


func main(){

	listen, err := net.Listen("tcp", ":9092")
	if err!=nil{
		fmt.Println(err)
		return
	}

	err = rpc.RegisterName("FoodService", &FoodService{})
	if err!=nil{
		fmt.Println(err)
		return
	}
	conn, err := listen.Accept()
	if err!=nil{
		fmt.Println(err)
		return
	}
	rpc.ServeConn(conn)

}


