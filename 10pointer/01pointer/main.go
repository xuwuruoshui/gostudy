package main

import "fmt"

// =====================
// 指针的简单使用
//======================

func main() {
	// 1.指针的赋值操作
	a := 10
	var b *int = &a
	*b = 2
	fmt.Println(a)

	// 2.为赋值的指针为空
	var c *string
	if c == nil {
		fmt.Println(c)
	}

	// 3.空指针无法赋值,因为内存地址为空
	var d *int
	//*d=100
	fmt.Println(d)

	// make和new的区别
	// 1.make和new都是用来申请内存
	// 2.new用得少，通常为基本数据类型申请内存，string、int,返回的是对应类型的指针(*stirng,*int)
	// 3.make是用来对slice、map、chan申请内存的,make函数返回的是对应的这三个类型本身
}
