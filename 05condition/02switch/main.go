package main

import "fmt"

// =====================
// switch
//======================

func main() {

	// 1.基本使用
	a := 11
	switch a {
	case 1, 3, 5, 7:
		fmt.Println("a是奇数")
	case 2, 4, 6, 8, 10:
		fmt.Println("a是偶数")
	default:
		fmt.Println("a大于10")
	}

	// 2.使用判断
	switch {
	case 0 < a && a <= 10:
		fmt.Println("0<a<=10")
	case a > 10:
		fmt.Println("a>10")
	}

}
