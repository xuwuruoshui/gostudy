package main

import (
	"fmt"
	"sync"
)

/**
* @creator: xuwuruoshui
* @date: 2021-11-02 17:02:00
* @content: 多个goroutine
 */

// 等待一组线程结束，父线程调用add方法设置等待的线程数量，被等待的线程结束时调用Done方法
// 存在意义:因为并不知到程序何时结束，长时间睡眠又不符合实际，所以就可以使用WaitGroup
var wg sync.WaitGroup

func hello(i int) {
	defer wg.Done()
	fmt.Println("Hello Goroutine!", i)
}

func main() {
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go hello(i)
	}

	wg.Wait()
}
