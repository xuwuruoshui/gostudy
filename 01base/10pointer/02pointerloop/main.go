package main

import "fmt"

// =====================
// 指针的访问,及其值的遍历
//======================

func main() {

	// 1.遍历赋值指针
	a := [5]int{1, 2, 3, 4, 5}
	var b [5]*int

	for i := 0; i < 5; i++ {
		b[i] = &a[i]
	}

	// 2.遍历打印指针值
	for i := 0; i < 5; i++ {
		fmt.Println(*b[i])
	}
}
