package main

import (
	"fmt"
	"rabbitmq_test/05topic/topic"
)

func main() {
	// 获取b.*
	topic.ComsumerEx("test.hello.deadletter", "deadletter", "b.*.*", func(msg string) {
		fmt.Println(msg)
	})
}
