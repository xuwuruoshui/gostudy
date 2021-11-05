package main

import "fmt"

/**
* @creator: xuwuruoshui
* @date: 2021-11-03 14:56:21
* @content: 多路复用
 */

func main(){
	ch := make(chan int,1)

	// select可以处理一个或多个channel的发送，接收操作
	// 如果多个case同时满足，select会随机选择一个

	for i:=0;i<10;i++{
		select{
		case x:= <-ch:
			fmt.Println(x)
		case ch<-i:
		}
	}
	close(ch)
}