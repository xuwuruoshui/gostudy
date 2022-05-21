package main

import (
	"fmt"
	"unicode"
)

// =====================
// 遍历 条件判断练习
//======================

func main() {

	// 1.99乘法表
	for i := 1; i < 10; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%dX%d=%d\t", i, j, i*j)
		}
		fmt.Println()
	}

	// 2.反向99乘法表
	for i := 9; i >= 0; i-- {
		for j := 1; j <= i; j++ {
			fmt.Printf("%dX%d=%d\t", i, j, i*j)
		}
		fmt.Println()
	}

	// 3.打印出现的中文个数
	// go语言中UTF中文占3个字节
	str := "你好,Golang!"
	count := 0
	for _, v := range str {
		if unicode.Is(unicode.Han,v){
			count++
		}
	}
	fmt.Println(count)

}
