package main

import (
	"rabbitmq_test/05topic/topic"
	"strconv"
	"time"
)

func main() {

	count := 0
	for {
		if count%3 == 0 {
			topic.PublicEx("test.hello.deadletter", "deadletter", "a.hello.name", "a.hello.name deadletter: "+strconv.Itoa(count))
		} else if count%3 == 1 {
			topic.PublicEx("test.hello.deadletter", "deadletter", "b.hello.uid", "b.hello.uid: "+strconv.Itoa(count))
		} else {
			topic.PublicEx("test.hello.deadletter", "deadletter", "b.haha.uid", "b.haha.uid: "+strconv.Itoa(count))
		}
		count++
		time.Sleep(time.Second)
	}
}
