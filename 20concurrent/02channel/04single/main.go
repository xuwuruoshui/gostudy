package main

import "fmt"

/**
* @creator: xuwuruoshui
* @date: 2021-11-03 13:33:35
* @content: 单向通道，只能发送或者只能接受
 */

// ch1只接收
func counter(out chan<- int) {
	for i := 0; i < 100; i++ {
		out <- i
	}
	close(out)
}

// ch2只接收 , ch1只发送
func squarer(out chan<- int, in <-chan int) {
	for i := range in {
		out <- i * i
	}
	close(out)
}

// ch2只发送
func printer(in <-chan int) {
	for i := range in {
		fmt.Println(i)
	}
}

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go counter(ch1)
	go squarer(ch2, ch1)
	printer(ch2)
}
