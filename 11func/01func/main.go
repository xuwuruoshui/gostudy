package main

import "fmt"

/**
* @creator: xuwuruoshui
* @date: 2021-10-06 22:32:27
* @content: 函数使用
 */

func main() {
	fmt.Println(sum(1, 2))
	fmt.Println(sum2(2, 3))
	fmt.Println(sum3(2, 3, "What's","up?"))
	sum4("可变参数")
	sum4("可变参数",1,2,3,4,5)
}

func sum(a int, b int) int {
	return a + b
}

// 返回参数可以有名字，最后一个return
func sum2(a int, b int) (result int) {
	result = a + b
	return
}

// 传参类型相同可以简写
func sum3(a , b int,c,d string) int{
	fmt.Println(c+d)
	return a+b
}

// 可变长度参数
func sum4(a string,b ...int){
	fmt.Println(a)
	fmt.Println(b)
}
