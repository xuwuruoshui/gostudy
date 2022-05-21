package main

import "fmt"

/**
* @creator: xuwuruoshui
* @date: 2021-10-24 23:41:54
* @content: 接口传指针
 */

type animal interface{
	move()
}

type dog struct{}
// func (d dog)move(){
// 	fmt.Println("狗移动...")
// }
func (d *dog)move(){
	fmt.Println("狗移动...")
}
type cat struct{}
// func (c cat)move(){
// 	fmt.Println("猫移动...")
// }
func (c *cat)move(){
	fmt.Println("猫移动...")
}

func main(){
	var a1 animal

	// 1. 接口使用值接收，结构体和结构体的指针都能用
	// a1 = dog{}
	// a1.move()
	// a1 = &cat{}
	// a1.move()

	// 2. 接口使用指针接收，只能结构体的指针使用
	a1 = &dog{}
	a1.move()
	a1 = &cat{}
	a1.move()
}