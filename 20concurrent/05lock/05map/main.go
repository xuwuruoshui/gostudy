package main

import (
	"fmt"
	"strconv"
	"sync"
)

/**
* @creator: xuwuruoshui
* @date: 2021-11-03 16:06:08
* @content: 并发map
 */

//var m1 = make(map[string]int)

// func get(key string) int{
// 	return m1[key]
// }

// func set(key string, value int){
// 	m1[key] = value
// }

// 使用普通的map,多线程下读写会报:fatal error: concurrent map writes
var m2 = sync.Map{}

func main() {
	wg := sync.WaitGroup{}

	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func(n int) {
			key := strconv.Itoa(n)
			// set(key,n)
			// fmt.Printf("key=%s,value=%d\n",key,get(key))
			m2.Store(key, n)
			value, _ := m2.Load(key)
			fmt.Printf("key=%s,value=%d\n", key, value)
			wg.Done()
		}(i)
	}

	wg.Wait()
}
