package main

import "fmt"

// =====================
// 简单的变量申明与赋值
//======================

func main() {

	// 1. 简单定义
	a := 1
	b := "simple"
	a1, b2 := 1, "haha"
	fmt.Println("a=", a, "b=", b)
	fmt.Println("a1=", a1, "b2=", b2)
	fmt.Println("========================")

	// 2. 使用var
	var c = 2
	var d = "var"
	fmt.Println("c=", c, "d=", d)
	fmt.Println("========================")

	// 3. 使用var指定类型
	var e int = 3
	var f string = "specify type"
	var e1, f1 int = 3, 5
	fmt.Println("e=", e, "f=", f)
	fmt.Println("e1=", e1, "f1=", f1)
	fmt.Println("========================")

	// 4. 简写var
	var (
		g int    = 4
		h string = "simplify"
	)
	// 以及C语言的打印方式
	fmt.Printf("g=%d,h=%s\n", g, h)
	fmt.Println("========================")

	// 5.补充一种定义int类型的方法
	var other = int(10)
	fmt.Println(other)

}
