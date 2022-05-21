package main

import (
	"fmt"
	"rabbitmq_test/03subscript/subscript"
	"strconv"
	"time"
)

func main() {

	i := 0
	for {
		err := subscript.PublicEx("test.hello", "fanout", "", "hello subscript"+strconv.Itoa(i))
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("发送消息到rabbitmq: ", i)
		i++
		time.Sleep(time.Second)
	}

}
