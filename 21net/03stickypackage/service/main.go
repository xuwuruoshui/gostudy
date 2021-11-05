package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
)

/**
* @creator: xuwuruoshui
* @date: 2021-11-05 13:53:47
* @content: 粘包service
服务端粘包，是指TCP会把接收的数据先放到缓冲区中,再通知应用层取数据,应用层可能没有及时取出数据
 */

// 解决办法


 func handle(conn net.Conn){
		defer conn.Close()
		reader := bufio.NewReader(conn)
		content := make([]byte,1024)
		
		for{
			n,err := reader.Read(content)
			if err == io.EOF{
				break
			}

			if err!=nil{
				fmt.Println("读取客户端数据失败:",err)
				break
			}

			fmt.Println("收到来自client的数据",string(content[:n]))
		}
 }

func main(){
	listen,err := net.Listen("tcp",":8080")
	if err!=nil{
		log.Panic("监听失败:",err)
	}
	defer listen.Close()

	for{
		conn,err := listen.Accept()
		if err!=nil{
			log.Println("接收失败:",err)
			continue
		}
		go handle(conn)
	}
	
}

