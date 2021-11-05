package main

import "fmt"

/**
* @creator: xuwuruoshui
* @date: 2021-10-06 22:32:27
* @content: defer
 */

func main() {
	// deferTest()
	// fmt.Println("=============")
	// deferTest2()
	// fmt.Println("=============")
	// fmt.Println(deferTest3())
	// fmt.Println("=============")
	// fmt.Println(deferTest4())
	fmt.Println(deferTest5())
}

// 1.被defer标明的语句在其他语句执行完后，最后执行
func deferTest() {
	fmt.Println("start")
	defer fmt.Println("haha")
	fmt.Println("end")
}

// 2.多个被defer修饰的语句，执行顺序像栈，先进后出
func deferTest2() {
	fmt.Println("1")
	defer fmt.Println("2")
	defer fmt.Println("3")
	defer fmt.Println("4")
	fmt.Println("5")
}

// 3.return操作并非原子操作
// defer修饰:
//		1.返回值赋值后
// 		2.再调用defer后的语句
// 		3.返回
// 虽然修改了a，但返回值已经赋值过了,所以最后返回为1
func deferTest3() int {
	a := 1
	defer func() {
		a++
	}()
	return a
}

// 4. 把返回参数命名为a,操作的仍然是返回值a,所以最后得到的a是3
func deferTest4() (a int) {
	defer func() {
		a++
	}()
	return 2
}


// 4. 这里内部函数其实++的是a的一个副本,所以还是2
func deferTest5() (a int) {
	defer func(a int) {
		a++
	}(a)
	return 2
}