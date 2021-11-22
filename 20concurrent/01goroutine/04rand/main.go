package main

import (
	"fmt"
	"math/rand"
	"time"
)

/**
* @creator: xuwuruoshui
* @date: 2021-10-25 16:54:44
* @content: 随机数种子
 */
func test() {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 5; i++ {
		r1 := rand.Int()
		r2 := rand.Intn(10)
		fmt.Println(r1, r2)
	}

}

func main() {
	test()
}
