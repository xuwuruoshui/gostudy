package main

import (
	"fmt"
	"net/rpc"
)

func main(){
	client, err := rpc.Dial("tcp", ":9092")
	if err != nil {
		fmt.Println(err)
		return
	}

	reply := ""
	err = client.Call("FoodService.SayName", "北京烤鸭", &reply)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(reply)
}
