package main

import (
	"fmt"
	"log"
	"pack_rpc/client/entity"
	"pack_rpc/client_proxy"
)

// 主函数
func main() {
	service := client_proxy.NewRectService("tcp", ":8000")
	var result int
	err := service.Area(entity.Params{Width: 100, Height: 5}, &result)
	if err != nil {
		log.Println(err)
	}
	fmt.Println("面积:", result)

}
