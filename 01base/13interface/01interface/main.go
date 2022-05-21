package main

import "fmt"

/**
* @creator: xuwuruoshui
* @date: 2021-10-23 13:04:19
* @content: 接口
 */

 // 结构体实现接口,接口下所有的方法都必须实现
type speaker interface {
	speak()
}

type dog struct {
}

type cat struct {
}

type human struct {
}

func (d dog) speak() {
	fmt.Println("汪汪汪~~~")
}

func (d cat) speak() {
	fmt.Println("喵喵喵~~~")
}

func (d human) speak() {
	fmt.Println("卧槽~~~")
}

func main() {
	var sp speaker
	fmt.Printf("%T\n",sp)
	fmt.Printf("%v\n",sp)
	
	sp = dog{}
	sp.speak()
	sp = cat{}
	sp.speak()
	sp = human{}
	sp.speak()

	// 动态类型和值
	fmt.Printf("%T\n", sp)
}
