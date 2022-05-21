package main

import (
	"fmt"
	"reflect"
)

/**
* @creator: xuwuruoshui
* @date: 2021-10-30 11:44:49
* @content: isNil()和isValid()
 */

type user struct {
	age int
}

func (u user) play() {
	fmt.Println("play")
}

func main() {
	// isNil()用于通道、函数、接口、map、切片，其他则panic
	// isValid()用于IsValid、String、Kind，其他则panic

	var a *int
	fmt.Println("var a *int IsNil:", reflect.ValueOf(a).IsNil())

	// nil值
	fmt.Println("nil IsValid:", reflect.ValueOf(nil).IsValid())

	b := user{10}
	// 通过反射从结构体中获取abc字段
	fmt.Println("不存在的结构体成员", reflect.ValueOf(b).FieldByName("abc").IsValid())
	fmt.Println("存在的结构体成员", reflect.ValueOf(b).FieldByName("age").IsValid())

	// 通过反射从结构体中获取abc字段
	fmt.Println("不存在的结构体成员", reflect.ValueOf(b).MethodByName("abc").IsValid())
	// TODO ???
	fmt.Println("存在的结构体方法", reflect.ValueOf(b).MethodByName("play").IsValid())


	c := map[string]int{}
	// 从map中获取一个不存在的键
	fmt.Println("map中不存在的键", reflect.ValueOf(c).MapIndex(reflect.ValueOf("a")).IsValid())

}
