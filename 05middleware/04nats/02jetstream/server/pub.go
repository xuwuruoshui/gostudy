package main

import (
	"fmt"
	"github.com/nats-io/nats.go"
	"time"
)

/**
* @creator: xuwuruoshui
* @date: 2022-06-24 17:09:43
* @content: 发布
* @start: docker run  -p 4222:4222 -d nats -js
 */

func main() {
	nc, _ := nats.Connect(nats.DefaultURL)
	js, _ := nc.JetStream(nats.PublishAsyncMaxPending(256))

	// Create a Stream
	js.AddStream(&nats.StreamConfig{
		Name:     "ORDERS",
		Subjects: []string{"ORDERS.*"},
	})

	// Simple Async Stream Publisher
	for i := 0; i < 9; i++ {
		js.PublishAsync("ORDERS.scratch", []byte("hello"))
	}

	select {
	case temp := <-js.PublishAsyncComplete():
		fmt.Println(temp)
	case <-time.After(5 * time.Second):
		fmt.Println("Did not resolve in time")
	}

}
