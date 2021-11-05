package main

import "fmt"

// =====================
// 指针中的指针
//======================

func main() {

	// 1.定义a的指针aP
	a := 100
	var aP *int = &a
	// 2. 定制aP的指针aPP
	var aPP **int = &aP
	fmt.Println(*aP)
	fmt.Println(**aPP)
}
