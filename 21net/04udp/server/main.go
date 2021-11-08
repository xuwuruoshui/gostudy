package main

import (
	"fmt"
	"net"
)

/**
* @creator: xuwuruoshui
* @date: 2021-11-08 09:15:49
* @content: udp服务端
 */

func main() {
	// 1.监听
	listen, err := net.ListenUDP("udp", &net.UDPAddr{
		IP:   net.IPv4(127, 0, 0, 1),
		Port: 30000,
	})
	if err != nil {
		fmt.Println("监听失败:", err)
		return
	}
	defer listen.Close()

	// 2. 循环读取
	for {
		var data [1024]byte
		// 读取到data,返回客户端的地址
		n, addr, err := listen.ReadFromUDP(data[:])
		if err != nil {
			fmt.Println("读取服务端数据失败:", err)
			continue
		}
		fmt.Printf("客户端:%v addr:%v count:%v\n", string(data[:n]), addr, n)

		// 读取完成后返回给客户端
		_, err = listen.WriteToUDP([]byte("Hi,I am Server"), addr)
		if err != nil {
			fmt.Println("向客户端写数据失败", err)
			continue
		}
	}
}
