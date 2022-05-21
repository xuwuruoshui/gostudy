package main

import "fmt"

// =====================
// 强制类型转换
//======================

func main() {

	sum := 17
	count := 5
	var mean float32

	// 乘法除法都需要强转
	mean = float32(sum) / float32(count)
	fmt.Printf("mean 的值: %f\n", mean)

}
