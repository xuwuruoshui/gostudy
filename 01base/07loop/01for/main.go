package main

import "fmt"

// =====================
// 遍历元素
//======================

func main() {
	// 1.普通遍历方式
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	// 2.类似while循环
	i := 10
	for i >= 0 {
		fmt.Println(1)
		i--
	}

	// 3.使用range遍历
	strings := []string{"a", "b"}
	//使用两个变量接受。如果有不想要的返回参数,直接用_接受
	for key, value := range strings {
		fmt.Printf("key=%d,value=%s\n", key, value)
	}

	// 4.无限循环
	// for{
	// 	fmt.Println("haha")
	// }

	fmt.Println("========================")

	// 5.跳出循环 break跳出当前循环，return就不写了，直接跳出当前函数
	for i := 0; i < 5; i++ {
		if i==2{
			break
		}
		fmt.Println(i)
	}
}
