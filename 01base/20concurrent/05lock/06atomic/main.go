package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

/**
* @creator: xuwuruoshui
* @date: 2021-11-03 16:22:35
* @content: 原子包
 */

type Counter interface{
	Inc()
	Load() int64
}

// 普通版
type CommonCounter struct{
	counter int64
}

func (c *CommonCounter) Inc(){
	c.counter++
}

func (c *CommonCounter)Load() int64{
	return c.counter
}

// 互斥锁版
type MutexCounter struct{
	counter int64
	lock sync.Mutex
}

func (m *MutexCounter) Inc(){
	m.lock.Lock()
	defer m.lock.Unlock()
	m.counter++
}

func (m *MutexCounter) Load() int64{
	m.lock.Lock()
	defer m.lock.Unlock()
	return m.counter
}

// 原子版
type AtomicCounter struct{
	counter int64
}

func (a *AtomicCounter) Inc(){
	atomic.AddInt64(&a.counter,1)
}

func (a *AtomicCounter) Load() int64{
	return atomic.LoadInt64(&a.counter)
}

func test(c Counter){
	var wg sync.WaitGroup
	start := time.Now()
	for i := 0; i < 10000000; i++ {
		wg.Add(1)
		go func ()  {
			c.Inc()
			wg.Done()
		}()
	}
	wg.Wait()
	end := time.Now()
	fmt.Println(c.Load(),end.Sub(start))
}


func main(){

	// 非并发安全
	c1 := CommonCounter{}
	test(&c1)
	
	// 使用互斥锁
	c2 := MutexCounter{}
	test(&c2)

	// 原子操作
	c3 := AtomicCounter{}
	test(&c3)
}