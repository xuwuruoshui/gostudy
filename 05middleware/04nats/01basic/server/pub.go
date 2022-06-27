package main

import (
	"fmt"
	"github.com/nats-io/nats.go"
)

/**
* @creator: xuwuruoshui
* @date: 2022-06-24 17:01:28
* @content: 发布
 */

func main() {
	conn, err := nats.Connect("nats://127.0.0.1:4222")
	if err != nil {
		fmt.Println(err)
		return
	}
	err = conn.Publish("hello", []byte("你好呀，我Nats"))
	if err != nil {
		fmt.Println(err)
	}
	select {}
}
