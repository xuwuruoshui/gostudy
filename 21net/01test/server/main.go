package main

import (
	"bufio"
	"fmt"
	"net"
)

/**
* @creator: xuwuruoshui
* @date: 2021-11-04 13:42:49
* @content: 服务端
 */

func process(conn net.Conn) {
	defer conn.Close()
	for {
		reader := bufio.NewReader(conn)
		var buf [128]byte
		n, err := reader.Read(buf[:])
		if err != nil {
			fmt.Println("read from client faild:", err)
			break
		}
		recStr := string(buf[:n])
		fmt.Println(recStr)
		conn.Write([]byte(recStr))
	}
}

func main() {

	// 1.监听20000端口
	listen, err := net.Listen("tcp", "127.0.0.1:20000")
	if err != nil {
		fmt.Println("listen faild:", err)
		return
	}

	// 2.接收数据
	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("accept faild:", err)
			continue
		}
		go process(conn)
	}
}
