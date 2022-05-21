package main

import (
	"rabbitmq_test/04router/router"
	"strconv"
	"time"
)

func main() {

	count := 0
	for {
		if count&1 == 1 {
			router.PublicEx("test.hello.deadletter", "direct", "one", "test hello: "+strconv.Itoa(count))
		} else {
			router.PublicEx("test.hello.deadletter", "direct", "two", "test hello: "+strconv.Itoa(count))
		}
		count++
		time.Sleep(time.Second)
	}
}
