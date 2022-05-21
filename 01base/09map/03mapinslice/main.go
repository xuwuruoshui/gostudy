package main

import "fmt"

// =====================
// creator: xuwuruoshui
// createTime:
// content: list中存map
//======================
func main() {
	slice1 := make([]map[string]int, 5, 10)
	map1 := make(map[string]int)
	slice1[0] = map1
	slice1[0]["test"] = 1
	slice1[0]["test2"] = 2
	map2 := make(map[string]int)
	slice1[1] = map2
	slice1[1]["try"] = 1
	slice1[1]["try2"] = 2

	fmt.Println(slice1)
}
