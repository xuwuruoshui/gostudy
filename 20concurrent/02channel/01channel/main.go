package main

import "fmt"

/**
* @creator: xuwuruoshui
* @date: 2021-11-03 10:10:51
* @content: channel: 使两个goroutine进行交互
 */

func main() {

	// 1.声明一个整形的通道
	var ch chan int
	// 通道引用类型, 所以为nil
	fmt.Println(ch)

	// 2. 初始化, 10为缓冲区大小(非必填)
	ch = make(chan int,10)

	// 3. 基本操作
	// 发送
	ch <- 10

	// 接收
	// 接收并赋值
	x := <- ch
	fmt.Println(x)
	// 接收忽略结果
	//<-ch

	// 关闭
	close(ch)
	// tip:
	// 1.对关闭的通道再发送值，会引发panic
	// 2.对关闭的通道进行接受,会一直获取直到通道为空
	// 3.对关闭且没有值的通道接收，得到对应的类型为零值
	// 4.关闭一个已经关闭的通道会导致panic




}
