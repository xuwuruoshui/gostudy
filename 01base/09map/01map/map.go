package main

import "fmt"

// =====================
// map
//======================

func main() {

	// map是引用类型
	var map1 map[string]int

	// 1.未初始化无法赋值
	if map1 == nil {
		// map1["tes"] = 1
		fmt.Println(map1)
	}

	map2 := make(map[string]int)
	map2["aa"] = 1
	map2["bb"] = 2

	fmt.Println(map2)

}
