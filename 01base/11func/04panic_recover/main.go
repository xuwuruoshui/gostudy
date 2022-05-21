package main

import "fmt"

/**
* @creator: xuwuruoshui
* @date: 2021-10-06 22:32:27
* @content: panic/recover
 */

func A(){
	fmt.Println("func A")
}

func B(){
	panic("func B")
}

func B1(){
	// 使用defer捕获异常
	defer func ()  {
		if nil!=recover(){
			fmt.Println("recover in B")
		}
	}()
	panic("func B")
}

func C(){
	fmt.Println("func C")
}

func main(){
	A()
	// 抛出异常
	// B()
	
	// 捕获异常
	B1()
	C()
}