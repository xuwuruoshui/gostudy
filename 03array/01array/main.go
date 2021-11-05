package main

import "fmt"

// =====================
// 数组的基本用法
//======================

func main() {

	// 1. 数组的赋值/便利
	var arr = [...]int{1, 2, 3, 4, 5}
	fmt.Println(arr)

	for i := 0; i < 5; i++ {
		fmt.Println(arr[i])
	}

	//2. 访问数组
	var arr2 [10]int

	for i := 0; i < 10; i++ {
		arr2[i] = i
	}

	fmt.Println(arr2)

	//3. 数组是值类型(切片才是引用类型)
	var arr3 = [10]int{1: 2, 3: 5, 6: 7}
	fmt.Println(arr3)
	test(arr3)
	fmt.Println(arr3)
}

func test(arr [10]int) {
	arr[0] = 1
	arr[1] = 3
}
