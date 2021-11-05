package main

import (
	"fmt"
	"reflect"
)

/**
* @creator: xuwuruoshui
* @date: 2021-10-29 17:19:37
* @content: 普通类型和reflect.Value和互相转换
 */

func main() {
	var a float32 = 3.14
	var b int64 = 100

	reflectValue(a)
	reflectValue(b)

	// int类型转reflect.ValueOf
	c := reflect.ValueOf(10)
	fmt.Printf("type c:%T\n", c)
}


func reflectValue(x interface{}){
	v := reflect.ValueOf(x)

	k := v.Kind()

	switch k {
	case reflect.Int64:
		fmt.Printf("%d\n",int64(v.Int()))
	case reflect.Float32:
		fmt.Printf("%f\n",float32(v.Float()))
	}
}