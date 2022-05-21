package main

import (
	"fmt"
	"time"
)

/**
* @creator: xuwuruoshui
* @date: 2021-10-06 22:32:27
* @content: 编写程序统计一段代码的执行耗时时间，单位精确到微秒。
 */

func main(){

	time1 := time.Now()
	a :=10
	b :=20
	c :=a+b
	fmt.Printf("c = %d\n", c)
	// 停一秒
	time.Sleep(time.Second)
	time2 := time.Now()

	fmt.Println(time2.Sub(time1).Microseconds())
}