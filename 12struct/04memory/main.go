package main

import "fmt"

/**
* @creator: xuwuruoshui
* @date: 2021-10-06 22:32:27
* @content: 结构体内存
 */

func main() {
	test := A{1, 2, 3, "fff"}
	fmt.Println(&(test.a))
	fmt.Println(&(test.b))
	fmt.Println(&(test.c))
	// 内存对齐
	fmt.Println(&(test.d))

}

type A struct {
	a int8
	b int8
	c int8
	d string
}
