package main

import "fmt"

// =====================
// map判断是否包含key、删除操作
//======================

func main() {

	map1 := map[string]string{"first": "一", "second": "二", "third": "三"}
	for key, value := range map1 {
		fmt.Println(key, value)
	}

	// 1.是否包含key
	isContainValue(map1, "first")
	isContainValue(map1, "aaa")

	// 2.删除
	delete(map1, "second")
	fmt.Println(map1)
}

// 判断是否包含value
func isContainValue(map1 map[string]string, str string) {
	value, ok := map1[str]
	if ok {
		fmt.Println(value)
	} else {
		fmt.Println("the value is empty")
	}
}
