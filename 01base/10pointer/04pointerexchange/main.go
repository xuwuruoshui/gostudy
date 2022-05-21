package main

import "fmt"

// =====================
// 指针交换元素
//======================

func main() {
	a := 100
	b := 200
	// 1.获取a,b的地址
	swap(&a, &b)
	fmt.Println(a, b)
}

func swap(a *int, b *int) {
	// 指针所指向的值，互相交换赋值
	*a, *b = *b, *a
}
