package main

import (
	"fmt"
	"net/rpc/jsonrpc"
)

/**
* @creator: xuwuruoshui
* @date: 2022-01-24 16:22:20
* @content: http rpc client
 */

func main(){
	client,err := jsonrpc.Dial("tcp",":8000")
	if err!=nil{
		panic(err)
	}
	slice := []int{1,2,3,4}
	var result int 
	client.Call("HelloRpc.Add",slice,&result)
	fmt.Println(result)
}