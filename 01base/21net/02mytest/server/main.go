package main

import (
	"bufio"
	"fmt"
	"net"
)

/**
* @creator: xuwuruoshui
* @date: 2021-11-25 18:20:28
* @content: 服务器
 */

func main() {
	listen, err := net.Listen("tcp", "127.0.0.1:8080")
	if err != nil {
		panic(err)
	}
	defer listen.Close()


	for{
		conn, err := listen.Accept()
		if err != nil {
			panic(err)
		}
		go handle(conn)
	}
}

func handle(conn net.Conn) {
	defer conn.Close()
	for{
		reader := bufio.NewReader(conn)
		content := make([]byte, 1024)
		n, err := reader.Read(content)
		if err != nil {
			panic(err)
		}
		fmt.Printf("客户端%v: %s\n", conn.RemoteAddr(), string(content[:n]))
	}	
}
