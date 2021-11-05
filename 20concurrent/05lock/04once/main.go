package main

import (
	"fmt"
	"sync"
)

/**
* @creator: xuwuruoshui
* @date: 2021-11-03 15:55:57
* @content: 单例
 */

type singleton struct{}

var instance *singleton
var once sync.Once

func GetInstance() *singleton{
	// 只会执行一次
	once.Do(func() {
		instance = &singleton{}
	})
	return instance
}

func main(){
	a := GetInstance()
	b := GetInstance()
	c := GetInstance()

	fmt.Println(a==b)
	fmt.Println(b==c)
	fmt.Printf("%T\n", a)
}