package main

import "fmt"

// =====================
// 二维数组
//======================

func main() {

	// 1.多维数组赋值/访问
	arr := [3][3]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	fmt.Println(arr)
	fmt.Println(arr[1][1])
	arr[2][2] = 0
	fmt.Println(arr)

	// 2.切片多维赋值/访问
	arr1 := [][]int{}
	row1 := []int{1, 2, 3}
	row2 := []int{4, 5, 6}
	arr1 = append(arr1, row1)
	arr1 = append(arr1, row2)

	fmt.Println(arr1)
	test(arr1)
	fmt.Println(arr1)

	// 这是数组不是切片，所以不能用append
	// arr1 := [2][2]int{}
	// row1 := []int{1,2,3}
	// row2 := []int{4,5,6}
	// arr1 = append(arr1, row1)
	// arr1 = append(arr1, row2)

	// fmt.Println(arr1)
	// test1(arr1)
	// fmt.Println(arr1)
}

func test(arr [][]int) {
	arr[1][1] = 2
}

func test1(arr [][]int) {
	arr[1][1] = 2
}
