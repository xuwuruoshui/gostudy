package main

import (
	"fmt"
	"runtime"
	"time"
)

/**
* @creator: xuwuruoshui
* @date: 2021-11-03 09:59:31
* @content: GOMAXPROCS
 */

func a() {
	for i := 1; i < 500; i++ {
		fmt.Println("A", i)
	}
}

func b() {
	for i := 1; i < 500; i++ {
		fmt.Println("B", i)
	}
}

func main() {
	// 配置n个OS线程同时执行Go代码, 默认为机器上的核心数
	//runtime.GOMAXPROCS(1)
	runtime.GOMAXPROCS(2)
	go a()
	go b()
	time.Sleep(time.Second)

}
