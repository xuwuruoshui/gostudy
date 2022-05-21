package main

import (
	"fmt"
	"sync"
	"time"
)

/**
* @creator: xuwuruoshui
* @date: 2021-11-03 15:27:27
* @content: 读写互斥锁
 */

var (
	x int
	wg sync.WaitGroup
	rwlock sync.RWMutex
	lock sync.Mutex
)

func write(){
	rwlock.Lock()
	//lock.Lock()
	x++
	time.Sleep(10*time.Millisecond)
	//lock.Unlock()
	rwlock.Unlock()
	wg.Done()
}

func read(){
	rwlock.RLock()
	//lock.Lock()
	time.Sleep(10*time.Millisecond)
	//lock.Unlock()
	rwlock.RUnlock()
	wg.Done()
}

func main(){
	start := time.Now()

	// 写10次
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go write()
	}

	// 读1000次
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go read()
	}
	wg.Wait()
	end := time.Now()
	fmt.Println(end.Sub(start))
	fmt.Println(x)
}