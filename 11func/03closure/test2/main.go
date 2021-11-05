package main

import "fmt"

/**
* @creator: xuwuruoshui
* @date: 2021-10-06 22:32:27
* @content: 闭包应用
 */

// 其他人写的函数
func sum(a,b int)int{
	return a+b
}


// 自己写的函数
func tempFunc(a,b int,f func(a,b int)int) func()int{
	return func () int{
		return f(a,b)
	}
}

func mysum(f func()int) int{
	return f()
}


func main(){
	test := tempFunc(10,20,sum)
	fmt.Println(mysum(test))
}