package main

import (
	"fmt"
	"reflect"
)

/**
* @creator: xuwuruoshui
* @date: 2021-10-29 16:58:54
* @content: 获取任意值的类型
 */

func main() {
	var a float32 = 3.14
	reflectType(a)
	var b int64 = 100
	reflectType(b)
}

func reflectType(x interface{}){
	v := reflect.TypeOf(x)
	fmt.Printf("type:%v\n", v)
}