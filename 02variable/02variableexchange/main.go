package main

import "fmt"

// =====================
// 元素交换
//======================

func main() {

	i, j := 2, 3
	fmt.Printf("i=%d,j=%d\n", i, j)

	// 1.变量交换
	i, j = j, i
	fmt.Printf("i=%d,j=%d\n", i, j)

	// 2.申明过的变量无法再申明
	// i,j  := 4, 5
	// 除非带一个未申明的就能通过编译
	i, k := 4, 5
	fmt.Printf("i=%d,k=%d\n", i, k)

}
