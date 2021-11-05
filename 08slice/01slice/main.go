package main

import "fmt"

// =====================
// 切片
//======================

func main() {

	// 1.切片赋值/访问

	// 使用make函数构建, 参数分别为:类型 长度 容量
	var slice1 = make([]int, 5, 5)

	slice1[0] = 1
	slice1[1] = 2

	fmt.Println("slice1:", slice1)
	fmt.Println("slice1 cap:", cap(slice1))
	slice1[2] = 3
	fmt.Println("slice1:", slice1)
	fmt.Println("slice1 cap:", cap(slice1))

	// 2.切片的容量默认为初始容量
	slice2 := []string{"a", "b", "c", "d"}
	fmt.Println(cap(slice2))
	slice2[0] = "aa"
	slice2[1] = "bb"

	// 3.下标访问切片,使用append拼接元素,扩容时默认长度翻一倍
	fmt.Println("slice2[1:2]", slice2[1:2])
	slice2 = append(slice2, "e")

	// 扩容1倍
	fmt.Println("slice2", slice2)
	fmt.Println("slice2 cap:", cap(slice2))
	// 扩容机制
	// 1.如果增长的容量长度超过了原来的容量长度的两倍，则将新开的容量作为当前容量
	// 2.如果增长的容量长度小于原来的两倍
	// 		2.1 容量小于1024就翻倍
	//		2.2 大于1024,就for循环提高原来的1/4,直到大于当前增长后的容量
	// tip: 使用append扩容后，是将原来数组复制到了另一个数组，指针所指向的位置也发生了改变

	// 4.拷贝切片元素
	var slice3 []string = make([]string, len(slice2), cap(slice2)*2)
	copy(slice3, slice2)
	fmt.Println("slice3", slice3)
	fmt.Println("slice3 cap", cap(slice3))

	// 5.从数组中拷贝元素
	arr := [2]string{"f", "g"}
	// 拷贝所有元素
	copy(slice3, arr[:])
	fmt.Println("slice3: ", slice3)

	// 6.数组赋值给切片,容量为从开始下标到数组的最后一位,长度就为当前元素个数
	arr1 := [4]string{"h", "i", "j", "k"}
	slice4 := arr1[1:3]
	fmt.Println("slice4:", slice4)
	fmt.Println("slice4: cap", cap(slice4))

	// 6.切片解构，拼接元素，容量为slice3的容量
	slice5 := append(slice3, slice2...)
	fmt.Println("slice3+slice2: ", slice5)
	fmt.Println("slice5: cap", cap(slice5))

}
