package main

import (
	"fmt"

	"github.com/nats-io/nats.go"
)

/**
* @creator: xuwuruoshui
* @date: 2022-06-24 17:09:43
* @content: 订阅
 */

func main() {
	conn, err := nats.Connect("nats://127.0.0.1:4222")
	if err != nil {
		return
	}

	// 开两个消费者
	for i := 1; i <= 2; i++ {
		dummy := i
		conn.Subscribe("hello", func(msg *nats.Msg) {
			fmt.Printf("消费者[%d]收到：%s\n", dummy, string(msg.Data))
			msg.Respond([]byte("I am nats learner!!!"))
		})
	}
	select {}
}
