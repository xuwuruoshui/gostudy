package main

import "fmt"

/**
* @creator: xuwuruoshui
* @date: 2021-10-16 17:34:49
* @content: 匿名字段的结构体
 */

type Student struct {
	int
	string
}

func main() {
	// 学号和姓名匿名字段, 如果再加一个string就会出现问题, 这种并不推荐使用
	s1 := Student{1, "zhangsan"}
	fmt.Println(s1)
	fmt.Println(s1.int)
	fmt.Println(s1.string)
}
