package main

import "fmt"

/**
* @creator: xuwuruoshui
* @date: 2021-10-06 22:32:27
* @content: type用法
 */

 // 自定义类型
type myInt int
// 类型别名
type yourInt = int

func main(){
	var a myInt = 10
	fmt.Println(a)
	fmt.Printf("%T\n", a)

	var b yourInt = 100
	fmt.Println(b)
	fmt.Printf("%T\n", b)
}