package main

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"log"
	"net"
)

/**
* @creator: xuwuruoshui
* @date: 2021-11-05 13:53:47
* @content: 粘包service
* 服务端粘包，是指TCP会把接收的数据先放到缓冲区中,再通知应用层取数据,应用层可能没有及时取出数据
 */

// 解决办法
func Decode(reader *bufio.Reader)(string,error){
	
	// 1.读取前4字节
	lable,err := reader.Peek(4)
	if err!=nil{
		log.Println("读取包头失败",lable)
	}
	lenghtBuff := bytes.NewBuffer(lable)

	// 2.创建一个int32,并接收上述4字节的接收成功表明数据是正确的
	var length int32
	err = binary.Read(lenghtBuff,binary.LittleEndian,&length)
	if err!=nil{
		return "",err
	}
	
	// 3.将所有数据转为int32判断字节数是否大于8字节,小于表示无数据
	if int32(reader.Buffered())<length+4{
		return "",err
	}
	
	// 4.读取真实数据
	pkg :=make([]byte,int(4+length))
	_,err = reader.Read(pkg)
	if err!=nil{
		return "",err
	}
	return string(pkg[4:]),nil
}

func process(conn net.Conn){
	defer conn.Close()
	reader := bufio.NewReader(conn)
	for{
		msg,err := Decode(reader)
		if err==io.EOF{
			return
		}
		if err!=nil{
			fmt.Println("解包出错:",err)
		}
		fmt.Println("收到来自client的数据",msg)
	}
}

// func handle(conn net.Conn) {
// 	defer conn.Close()
// 	reader := bufio.NewReader(conn)
// 	content := make([]byte, 1024)

// 	for {
// 		n, err := reader.Read(content)
// 		if err == io.EOF {
// 			break
// 		}

// 		if err != nil {
// 			fmt.Println("读取客户端数据失败:", err)
// 			break
// 		}

// 		fmt.Println("收到来自client的数据", string(content[:n]))
// 	}
// }

func main() {
	listen, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Panic("监听失败:", err)
	}
	defer listen.Close()

	for {
		conn, err := listen.Accept()
		if err != nil {
			log.Println("接收失败:", err)
			continue
		}
		//go handle(conn)
		go process(conn)
	}

}
