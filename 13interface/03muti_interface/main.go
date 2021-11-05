package main

import "fmt"

/**
* @creator: xuwuruoshui
* @date: 2021-10-24 23:54:31
* @content: 同一结构体实现多个接口
 */

type music interface{
	sing()
}

type sport interface{
	run()
}

type human struct{

}

func (h human)sing(){
	fmt.Println("唱歌...")
}

func (h human)run(){
	fmt.Println("跑步...")
}


// 两个接口合成一个
type skill interface{
	music
	sport
}

func main(){
	h1 := human{}

	var m1 music = h1
	m1.sing()

	var s1 sport = h1
	s1.run()

	var sk1 skill = h1
	sk1.run()
	sk1.sing()
}