package main

import (
	"fmt"
	"sync/atomic"
)

/**
* @creator: xuwuruoshui
* @date: 2021-11-25 16:00:07
* @content: CAS
 */

func main(){
	var a int64 = 10
	ok:=atomic.CompareAndSwapInt64(&a,10,20)
	ok=atomic.CompareAndSwapInt64(&a,10,20)
	fmt.Println(ok,a)
}