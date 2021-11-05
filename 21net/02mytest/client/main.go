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
* @date: 2021-11-04 16:38:58
* @content: client
 */

func main() {
	// 1、连接，关闭
	connect, err := net.Dial("tcp", "127.0.0.1:8080")
	if err != nil {
		fmt.Println("连接失败:", err)
	}
	defer connect.Close()

	// 2、等待用户输入阻塞
	reader := bufio.NewReader(os.Stdin)

	for {

		// 读一行
		str, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("读取错误:", str)
		}

		content := strings.Trim(str, "\r\n")
		if strings.ToUpper(content) == "Q" {
			fmt.Println("结束会话")
			break
		}

		// 4、连接中，传输用户输入内容到服务端
		_, err = connect.Write([]byte(content))
		if err != nil {
			fmt.Println("发送至服务端失败:", err)
		}

		// 5、连接中，读取服务端消息(读取也会被阻塞)
		// msg := make([]byte,1024)
		// n,err := connect.Read(msg)
		// if err!=nil{
		// 	fmt.Println("接收服务端消息失败:",err)
		// }

		// fmt.Println(string(msg[:n]))
	}
}
