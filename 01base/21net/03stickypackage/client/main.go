package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
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
func Encode(msg string) ([]byte,error){

	// 1.将消息长度作为消息头(4字节)
	length := int32(len(msg))

	// 2.创建一个字节buffer
	pkg := new(bytes.Buffer)

	// 3.消息头写入buffer
	err :=binary.Write(pkg,binary.LittleEndian,length)
	if err!=nil {
		log.Println("消息头写入异常",err)
	}

	// 4.消息体写入buffer
	err = binary.Write(pkg,binary.LittleEndian,[]byte(msg))
	if err!=nil{
		log.Println("消息体写入异常",err)
	}

	return pkg.Bytes(),err
}


func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8080")
	if err != nil {
		log.Panic("连接失败", err)
	}
	defer conn.Close()

	// for i := 1; i <= 20; i++ {
	// 	conn.Write([]byte("Hello,How are you today?"))
	// }

	for i := 0; i < 20; i++ {
		msg := `Hello,Hello.How are you?`
		data,err := Encode(msg)
		if err!=nil{
			fmt.Println("封包失败:",err)
		}
		conn.Write(data)
	}
}
