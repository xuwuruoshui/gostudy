package main

import (
	"fmt"
	"reflect"
)

/**
* @creator: xuwuruoshui
* @date: 2021-10-29 17:05:30
* @content: 获取类型
 */

type myInt int64

type student struct {
	name string
}

type book struct {
	name string
}

type play interface {
}

func reflectType(x interface{}) {
	t := reflect.TypeOf(x)
	fmt.Printf("type:%v kind:%v\n", t.Name(), t.Kind())
}

func main() {
	var a *float32
	var b myInt
	var c rune
	var d student
	var e book
	var f play = "aaa"

	reflectType(a)
	reflectType(b)
	reflectType(c)
	reflectType(d)
	reflectType(e)
	// 空指针直接报错
	reflectType(f)
}
