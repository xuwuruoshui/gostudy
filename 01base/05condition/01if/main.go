package main

import "fmt"

// =====================
// if条件判断
//======================

func main() {
	a := 10

	// 1.if else
	if a > 20 {
		fmt.Println(a)
	} else {
		fmt.Println("a<20")
	}

	// 2.if elseif else
	if b := 15; b > 20 {
		fmt.Println("b>20")
	} else if b < 10 {
		fmt.Println("b<20")
	} else {
		fmt.Println("10<b<20")
	}
	//在if判断时赋值，作用域只能在if中有效
	//fmt.Println(b)
}
