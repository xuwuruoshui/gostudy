package main

import "fmt"

/**
* @creator: xuwuruoshui
* @date: 2021-10-06 22:32:27
* @content: 闭包, 函数中返回函数
 */

func sum() func(int) int {
	a := 100
	return func(b int) int {
		a += b
		return a
	}
}

func main() {
	// 同一个函数调用时共用变量a
	f := sum()
	fmt.Println(f(10))
	fmt.Println(f(10))
}
