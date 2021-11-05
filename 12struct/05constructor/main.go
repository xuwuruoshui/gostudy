package main

import (
	"fmt"
	"unsafe"
)

/**
* @creator: xuwuruoshui
* @date: 2021-10-06 22:32:27
* @content: 构造函数
 */

func main() {

	user1 := newUser()
	user2 := newUser1()

	fmt.Printf("%#p\n", user1)
	fmt.Printf("%d\n", unsafe.Sizeof(user1))
	fmt.Printf("%d\n", unsafe.Sizeof(user2))
}

// 构造函数通常返回指针，一个指针只占8字节，而直接返回结构体特别占用内存。
// 构造函数是值类型！！！
func newUser() *user {
	return &user{1, "a"}
}

func newUser1() user {
	return user{1, "a"}
}

type user struct {
	id   int
	name string
}
