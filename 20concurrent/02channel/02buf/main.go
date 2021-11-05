package main

import (
	"fmt"
)

/**
* @creator: xuwuruoshui
* @date: 2021-11-03 10:56:10
* @content: 缓冲通道
 */

func recv(ch chan int) {
	a := <-ch
	fmt.Println("接收成功", a)
}

func main() {
	// 1.无缓冲通道
	ch := make(chan int)
	// 只发不接,deadlock,需要先接收阻塞，再传输数据
	go recv(ch)
	ch <- 10
	fmt.Println("ch:发送完成")

	// 2.缓存通道
	ch2 := make(chan int, 1)
	ch2 <- 20
	fmt.Println("ch2:发送完成")
}
