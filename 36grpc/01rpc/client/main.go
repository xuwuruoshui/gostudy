package main

import (
	"fmt"
	"log"
	"net/rpc"
)

// 参数
type Params struct {
	Width,Height int
}

// 主函数
func main(){
	// 1.连接远程rpc服务
	conn,err := rpc.DialHTTP("tcp",":8000")
	if err != nil {
		log.Fatalln(err)
	}

	// 2.远程调用方法
	// 面积计算
	result := 0
	err = conn.Call("Rect.Area",Params{50,100},&result)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("面积:",result)

	// 周长
	result = 0
	err = conn.Call("Rect.Perimeter",Params{50,100},&result)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("周长:",result)
}
