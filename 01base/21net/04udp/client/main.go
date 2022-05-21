package main

import (
	"fmt"
	"net"
)

/**
* @creator: xuwuruoshui
* @date: 2021-11-08 09:24:57
* @content: udp客户端
 */

func main(){
	// 1.连接
	socket,err := net.DialUDP("udp",nil,&net.UDPAddr{
		IP: net.IPv4(127,0,0,1),
		Port: 30000,
	})

	if err!=nil{
		fmt.Println("连接服务端失败:",err)
	}
	defer socket.Close()

	// 2.发送数据到服务端
	sendData := []byte("Hello server")
	_,err = socket.Write(sendData)
	if err!=nil{
		fmt.Println("发送数据失败:",err)
	}

	// 3.读取服务端返回数据
	data := make([]byte,4096)
	n,remoteAddr,err := socket.ReadFromUDP(data)
	if err!=nil{
		fmt.Println("接收数据失败",err)
		return
	}

	fmt.Printf("服务端:%v addr:%v count:%v\n",string(data[:n]),remoteAddr,n )
}