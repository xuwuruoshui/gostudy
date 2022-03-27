package main

import (
	"fmt"
	"rabbitmq_test/05topic/topic"
)

func main() {
	// 获取所有路由
	topic.ComsumerEx("test.hello.deadletter", "deadletter", "#", func(msg string) {
		fmt.Println(msg)
	})
}
