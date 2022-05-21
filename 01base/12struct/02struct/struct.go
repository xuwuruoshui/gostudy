package main

import "fmt"

/**
* @creator: xuwuruoshui
* @date: 2021-10-06 22:32:27
* @content: 结构体
 */

func main() {

	// 1.值一一对应结构体初始化
	book1 := Books{"Go语言学习", "xuwuruoshui", "Go语言的相关基础知识", 1}
	fmt.Println(book1)

	// 2.类似与json的初始化
	book2 := Books{author: "Go语言学习", title: "aa", id: 1}
	fmt.Println(book2)
	book2.author = "zhangsan"
	book2.title = "Golang的学习之旅"
	book2.subject = "Go语言进阶"
	book2.id = 2
	fmt.Println(book2)
	fmt.Println(book2.id)

	// 2.匿名结构体(很少用)

	var noname struct {
		name string
		age  int
	}
	noname.name = "haha"
	noname.age = 10
	fmt.Println(noname)
}

type Books struct {
	title   string
	author  string
	subject string
	id      int
}
