package main

import (
	"fmt"
	"rabbitmq_test/04router/router"
)

func main() {
	router.ComsumerEx("test.hello.deadletter", "direct", "two", func(msg string) {
		fmt.Println(msg)
	})
}
