package main

import (
	"bufio"
	"fmt"
	"net"
)

/**
* @creator: xuwuruoshui
* @date: 2021-11-04 16:39:10
* @content: service
 */

func main() {
	//1、监听
	listen, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("监听失败:", err)
	}

	for {
		//2、接收客户端数据
		connect, err := listen.Accept()
		if err != nil {
			fmt.Println("连接失败:", err)
		}
		go handle(connect)
	}

}

func handle(connect net.Conn) {
	// 3、从连接中读取客户端消息(连接转reader)
	reader := bufio.NewReader(connect)
	content := make([]byte, 1024)
	n, err := reader.Read(content)
	if err != nil {
		fmt.Println("读取消息失败:", err)
	}
	fmt.Println(string(content[:n]))

	// 4、从连接中，返回给客户端消息
	// n,err = connect.Write(content)
	// if err != nil {
	//         fmt.Println("返回失败:", err)
	//     }
}
