package main

import (
	"log"
	"net"
)

/**
* @creator: xuwuruoshui
* @date: 2021-11-05 15:54:19
* @content: 客户端粘包
* 客户端粘包，是指数据发送时，TCP并不会立即发送，会先等待一段时间，看是否还有要发送的数据，有就一起发送
 */

// encode 将消息编码
func Encode(){
	
}


func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8080")
	if err != nil {
		log.Panic("连接失败", err)
	}
	defer conn.Close()

	for i := 1; i <= 20; i++ {
		conn.Write([]byte("Hello,How ary your today?"))
	}
}
