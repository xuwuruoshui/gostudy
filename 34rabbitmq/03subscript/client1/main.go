package main

import (
	"fmt"
	"rabbitmq_test/03subscript/subscript"
)

func main() {
	subscript.ComsumerEx("test.hello","fanout","", func(s string) {
		fmt.Println(s)
	})
}
