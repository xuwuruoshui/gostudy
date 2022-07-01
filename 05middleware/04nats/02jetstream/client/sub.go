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
	// Connect to NATS
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		fmt.Println(err)
	}

	// Create JetStream Context
	// 创建消费者
	js, _ := nc.JetStream()
	js.AddConsumer("ORDERS", &nats.ConsumerConfig{
		Durable: "MONITOR",
	})
	sub, err := js.PullSubscribe("ORDERS.*", "MONITOR")
	if err != nil {
		fmt.Println(err)
	}

	// 分批从stream拉取消息一次10个
	msgs, err := sub.Fetch(10)
	if err != nil {
		fmt.Println(err)
	}
	for _, msg := range msgs {
		err = msg.Ack()
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(string(msg.Data))
	}
}
