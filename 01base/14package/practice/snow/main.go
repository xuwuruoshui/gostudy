package main

import (
	"fmt"
	c "../calc"
)

/**
* @creator: xuwuruoshui
* @date: 2021-10-28 10:42:37
* @content: 包的使用
 */

func main(){
	fmt.Println(c.Add(1,2))
	fmt.Println(c.Subtract(3,4))
	fmt.Println(c.Multiply(10,2))
	fmt.Println(c.Divide(100,10))
}