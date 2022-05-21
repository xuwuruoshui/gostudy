package main

import (
	"fmt"
	"rabbitmq_test/06deadletter/deadletter"
)

func main(){
	deadletter.ComsumerDlx("test.dlx.a","test_dlx_a","test.dlx.b","test_dlx_b",1000, func(callBack string) {
		fmt.Println(callBack)
	})
}
