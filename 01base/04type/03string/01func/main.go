package main

import (
	"fmt"
	"strings"
)

// =====================
// 字符串相关操作函数
//======================

func main() {
	// 1.字符拼接 Sprintf
	a := "Hello,"
	b := "Golang"
	result := fmt.Sprintf("%s%s", a, b)
	fmt.Println(result)

	// 2.分割 split
	result1 := strings.Split(result,",")
	fmt.Println(result1)

	// 3.包含contain
	result2 := strings.Contains(result,"ell")
	fmt.Println(result2)

	// 4.前缀和后缀 HasPrefix/HasSuffix
	result3 := strings.HasPrefix(result,"He")
	result4 := strings.HasSuffix(result,"ng")
	fmt.Println(result3,result4)

	// 5.字符串穿线的位置index/lastIndex
	result5 := strings.Index(result,"ell")
	result6 := strings.LastIndex(result,"an")
	fmt.Println(result5,result6)

	// 6.拼接
	strArr := []string{"Go","Java","Python","C","C++"}
	result7 := strings.Join(strArr,",")
	fmt.Println(result7)

}