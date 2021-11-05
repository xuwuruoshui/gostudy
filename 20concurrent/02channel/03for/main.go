package main

import "fmt"

/**
* @creator: xuwuruoshui
* @date: 2021-11-03 11:20:55
* @content: 通道循环取值
 */

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	// 1.开启goroutine将0-100的数发送到ch1
	go func() {
		for i := 0; i < 100; i++ {
			ch1 <- i
		}
		close(ch1)
	}()

	// 2.再开一个goroutine从ch1中值的平方
	go func() {
		for {
			i, ok := <-ch1
			// ok为false表示已经取完了, i此时为0
			if !ok {
				break
			}
			ch2 <- i * i
		}
		close(ch2)
	}()

	// 3.打印ch2
	for v := range ch2 {
		fmt.Println(v)
	}
}
