package main

import (
	"fmt"
	"sync"
)

/**
* @creator: xuwuruoshui
* @date: 2021-11-03 15:06:27
* @content: 并发问题
 */

var x int
var wg sync.WaitGroup

func add() {
	for i := 0; i < 100000; i++ {
		x++
	}

	wg.Done()
}

func main() {
	wg.Add(2)
	go add()
	go add()
	wg.Wait()

	// 最终x应该为2万
	fmt.Println(x)
}
