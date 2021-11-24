package main

import (
	"fmt"
	"sync"
)

/**
* @creator: xuwuruoshui
* @date: 2021-11-04 16:39:10
* @content: service
 */
 var wg sync.WaitGroup

func work(i int, ch <-chan int, ch2 chan<- int) {
	for v := range ch {
		ch2 <- v * 2
		fmt.Println("第", i, "个goroutine,start执行任务", v)
		fmt.Println("第", i, "个goroutine,end执行任务", v)
	}
}

func main() {

	ch := make(chan int, 10)
	ch2 := make(chan int, 10)

	for i := 1; i <= 3; i++ {
		go work(i, ch, ch2)
	}

	for i := 1; i <= 5; i++ {
		ch <- i
	}

	for i := 0; i < 5; i++ {
		fmt.Println(<-ch2)
	}

}
