package main

import (
	"fmt"
	"sync"
	"time"
)

func main(){

	var lock sync.Mutex
	cond := sync.NewCond(&lock)

	var wg sync.WaitGroup
	wg.Add(5)
	// 创建一些协程打印
	for i := 0; i < 5; i++ {
		go func(i int) {
			defer wg.Done()
			// 必须加锁
			cond.L.Lock()
			defer cond.L.Unlock()
			cond.Wait()
			fmt.Println(i)
		}(i)
	}

	time.Sleep(time.Second*2)

	// 唤醒一个
	cond.Signal()

	time.Sleep(time.Second*2)

	// 唤醒一个
	cond.Signal()


	time.Sleep(time.Second*2)


	// 全唤醒
	cond.Broadcast()
	wg.Wait()

}
