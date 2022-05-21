package main

import (
	"fmt"
	"rabbitmq_test/04router/router"
)

func main() {
	router.ComsumerEx("test.hello.deadletter", "direct", "one", func(msg string) {
		fmt.Println(msg)
	})
}
