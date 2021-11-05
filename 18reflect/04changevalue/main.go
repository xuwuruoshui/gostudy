package main

import (
	"fmt"
	"reflect"
)

/**
* @creator: xuwuruoshui
* @date: 2021-10-30 11:30:13
* @content: 反射设置变量
 */

func main() {
	var a int64 = 100
	//reflectSetValue1(a)
	// 函数传参是值拷贝，需要传变量地址，然后通过Elem()方法获取指针对应的值
	reflectSetValue2(&a)
	fmt.Println(a)
}

func reflectSetValue1(x interface{}) {
	v := reflect.ValueOf(x)
	if v.Kind() == reflect.Int64 {
		v.SetInt(200) //修改的是副本(如果x接受的是地址，不会报错)，直接panic引发异常
	}
}

func reflectSetValue2(x interface{}) {
	v := reflect.ValueOf(x)
	// Elem()获取指针对应的值
	if v.Elem().Kind() == reflect.Int64 {
		v.Elem().SetInt(200)
	}
}
