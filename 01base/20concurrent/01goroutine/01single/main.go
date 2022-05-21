package main

import (
	"fmt"
	"time"
)
/**
* @creator: xuwuruoshui
* @date: 2021-11-02 17:02:32
* @content: 单个goroutine
*/


func hello(){
	fmt.Println("Hello Goroutine")
}

func main() {
	// 1.main结束,hello可能还未打印
	// main结束goroutine不管执行完没有，都会结束
	go hello()
	fmt.Println("main Goroutine done")
	
	// 可以等待1s,等hello执行
	time.Sleep(time.Second)

}
