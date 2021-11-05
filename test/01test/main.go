package main

import (
	"fmt"
	"sync"
	"time"
)

/**
* @creator: xuwuruoshui
* @date: 2021-10-25 16:54:44
* @content:
 */

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	wg := sync.WaitGroup{}

	wg.Add(2)
	go func(){
		ch1 <- 10
		ch2 <- 20
		wg.Done()
	}()

	go func ()  {
		time.Sleep(time.Second)
		ch2 <- (<-ch1 + <-ch2)
		wg.Done()
	}()

	wg.Wait()

	fmt.Println(<-ch2)
}
