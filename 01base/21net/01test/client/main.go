package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

/**
* @creator: xuwuruoshui
* @date: 2021-11-04 14:03:28
* @content: 客户端
 */

func main() {

	// 1.连接127.0.0.1:20000
	conn, err := net.Dial("tcp", "127.0.0.1:20000")
	if err != nil {
		fmt.Println("err:", err)
		return
	}

	defer conn.Close()

	// 等待用户输入
	inputReader := bufio.NewReader(os.Stdin)

	for {
		// 按行读
		input, _ := inputReader.ReadString('\n')
		inputInfo := strings.Trim(input, "\r\n")

		// 读到Q时结束
		if strings.ToUpper(inputInfo) == "Q" {
			return
		}

		// 将读取的信息写入
		_, err = conn.Write([]byte(inputInfo))

		if err != nil {
			return
		}

		buf := [512]byte{}

		// 读取服务端的消息
		n, err := conn.Read(buf[:])
		if err != nil {
			fmt.Println("recv faild:", err)
			return
		}

		fmt.Println(string(buf[:n]))
	}
}
