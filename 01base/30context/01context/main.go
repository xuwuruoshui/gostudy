package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

/**
* @creator: xuwuruoshui
* @date: 2021-11-30 20:25:19
* @content: context,
 */

var wg sync.WaitGroup

func worker2(ctx context.Context) {

	for {
		fmt.Println("worker2")
		time.Sleep(time.Second)
		select {
		case <-ctx.Done():
			return
		default:
		}
	}
}

func worker(ctx context.Context) {
	defer wg.Done()
	go worker2(ctx)
	for {
		fmt.Println("worker1")
		time.Sleep(time.Second)
		select {
		case <-ctx.Done():
			return
		default:
		}
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	wg.Add(1)
	go worker(ctx)
	time.Sleep(time.Second * 5)
	cancel()
	wg.Wait()
	fmt.Println("结束")
}
