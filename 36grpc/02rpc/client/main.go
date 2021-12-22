package main

import (
	"fmt"
	"log"
	"net/rpc/jsonrpc"
)

// 参数
type Params struct {
	Width,Height int
}



func main(){
	conn,err := jsonrpc.Dial("tcp","120.78.159.42:8000")
	if err != nil {
		log.Fatalln(err)
	}

	result := 0
	err = conn.Call("Rect.Area",Params{50,100},&result)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("面积:",result)

	result = 0
	err = conn.Call("Rect.Perimeter",Params{50,100},&result)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("周长:",result)
}
