package main

import (
	"fmt"
	"math/rand"
	"sync"

	//"sync"
	"time"
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

 // Tip:代码中使用的无限循环。如果是固定循环100个任务,接收时，要保证发送的channel关闭后再接受,
 // 否则可能在接受101时直接deadlock。当然也可以不用关闭，接受只接受100次就行了。
type Consumer struct{
	result int64
	origin int64
}

var wg sync.WaitGroup
func f1(jobChain chan int64) {
	rand.Seed(time.Now().UnixNano())
	for{
		jobChain <- int64(rand.Int63())
	}
}

func f2(jobChain chan int64, resultChain chan *Consumer) {
	// lock.Lock()
	// defer lock.Unlock()
	defer wg.Done()
	for {
		value, ok := <-jobChain
		if !ok {
			return
		}
		con := Consumer{origin:value}
		var result int64
		for {
			result += value % 10
			if value/10 < 10 {
				time.Sleep(time.Millisecond*100)
				con.result = result
				resultChain <- &con
				break
			}
			value = value / 10
		}
	}
}

func main() {
	jobChain := make(chan int64, 100)
	resultChain := make(chan *Consumer, 100)


	go f1(jobChain)
	wg.Add(24)
	for i := 0; i < 24; i++ {
		go f2(jobChain, resultChain)
	}

	for v := range resultChain {
		time.Sleep(time.Millisecond*100)
		fmt.Println(v.origin,v.result)
	}

	wg.Wait()

	


}
