package main

import (
	"fmt"
	//"time"
)

/**
* @creator: xuwuruoshui
* @date: 2021-11-03 14:17:43
* @content: 线程池
 */

func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Printf("worker:%d start job:%d\n", id, j)
		//time.Sleep(time.Second)
		fmt.Printf("worker:%d end job:%d\n", id, j)
		results <- j * 2
	}
}

func main() {
	jobs := make(chan int, 100)
	results := make(chan int, 100)

	// 开启3个goroutine
	for i := 1; i <= 3; i++ {
		go worker(i, jobs, results)
	}

	// 5个任务
	for i := 1; i <= 5; i++ {
		jobs <- i
	}

	close(jobs)

	// 打印最终获取的值
	for i := 1; i <= 5; i++ {
		fmt.Println(<-results)
	}
}
