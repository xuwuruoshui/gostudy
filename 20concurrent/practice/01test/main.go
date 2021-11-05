package main

import (
	"fmt"
	"sync"
)

/**
* @creator: xuwuruoshui
* @date: 2021-11-03 17:00:50
* @content:
* 使用goroutine和channel实现一个计算int64随机数各位数和的程序。
* 	1.开启一个goroutine循环生成int64类型的随机数，发送到jobChan
* 	2.开启24个goroutine从jobChan中取出随机数计算各位数的和，将结果发送到resultChan
* 	3.主goroutine从resultChan取出结果并打印到终端输出
 */

func main() {

	jobChan := make(chan int64)
	resultChan := make(chan int64, 1)
	wg := sync.WaitGroup{}
	i := 0

	go func() {
		resultChan <- 0
		for {
			i++
			v := int64(i)
			jobChan <- v
			fmt.Printf("%d\n", v)
		}
	}()

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			v := <-jobChan
			temp := <-resultChan
			resultChan <- (v + temp)
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println(<-resultChan)

}
