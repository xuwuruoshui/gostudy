package main

import (
	"reflect"
)

func GetValue() int {
	return 1
}

func main() {
	i := GetValue()
	switch reflect.TypeOf(i).Kind() {
	case reflect.Int:
			println("int")
	case reflect.String:
		 println("string")
	case reflect.Interface:
		 println("interface")
	default:
		 println("unknown")
 }
}
