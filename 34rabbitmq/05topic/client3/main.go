package main

import (
	"fmt"
	"rabbitmq_test/05topic/topic"
)

func main() {
	topic.ComsumerEx("test.hello.topic", "topic", "*.hello.*", func(msg string) {
		fmt.Println(msg)
	})
}
