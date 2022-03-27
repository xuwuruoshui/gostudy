package main

import (
	"rabbitmq_test/06deadletter/deadletter"
	"strconv"
	"time"
)

func main() {

	count := 0
	for {
		// msg->A->B
		deadletter.PublicDlx("test.dlx.a","deadletter test: "+strconv.Itoa(count))
		// msg->B
		deadletter.PublicEx("test.dlx.b","fanout","","normal exchange test: "+strconv.Itoa(count))
		count++
		time.Sleep(time.Second)
	}
}
