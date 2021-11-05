package main

import "fmt"

// =====================
// 字符串修改
//======================


func main() {
	s1 := "白萝卜"
	// rune == int32,只是名字不一样
	s2 := []rune(s1)
	//s2 := []int32(s1)
	s2[0] = '红'

	fmt.Println(string(s2))
}