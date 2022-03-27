package main

import (
	"fmt"
	"rabbitmq_test/05topic/topic"
)

func main() {
	// 获取*.hello.*
	topic.ComsumerEx("test.hello.deadletter", "deadletter", "*.hello.*", func(msg string) {
		fmt.Println(msg)
	})
}
