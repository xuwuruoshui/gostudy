package main

import (
	"fmt"
	"reflect"
)

/**
* @creator: xuwuruoshui
* @date: 2021-10-31 19:01:27
* @content: 结构体反射
 */

type student struct {
	Name  string `json:"name"`
	Score int    `json:"score"`
}

// 给student添加两个方法 Study和Sleep(注意首字母大写)
func (s student) Study() string {
	msg := "好好学习，天天向上。"
	fmt.Println(msg)
	return msg
}

func (s student) Sleep() string {
	msg := "好好睡觉，快快长大。"
	fmt.Println(msg)
	return msg
}

func main() {
	stu1 := student{
		Name:  "张三",
		Score: 90,
	}

	t := reflect.TypeOf(stu1)
	fmt.Println(t.Name(), t.Kind())

	// 1.通过for循环遍历结构体的所有字段信息
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		fmt.Printf("name:%s index:%d type:%v json tag:%v\n", field.Name, field.Index, field.Type, field.Tag.Get("json"))
	}

	// 2.通过字段名获取指定结构体字段信息
	if scoreField, ok := t.FieldByName("Score"); ok {
		fmt.Printf("name:%s index:%d type:%v json tag:%v\n", scoreField.Name, scoreField.Index, scoreField.Type, scoreField.Tag.Get("json"))
	}

	// 3.打印stu1包含的方法
	printMethod(stu1)
}

func printMethod(x interface{}) {
	// 获取字段信息
	t := reflect.TypeOf(x)
	// 修改字段、调用方法
	v := reflect.ValueOf(x)

	fmt.Println(t.NumMethod())
	for i := 0; i < v.NumMethod(); i++ {
		methodType := v.Method(i).Type()
		fmt.Printf("method name:%s\n", t.Method(i).Name)
		fmt.Printf("method:%s\n", methodType)

		// 通过反射调用方法传递参数必须是[]reflect.Value类型
		var args = []reflect.Value{}
		v.Method(i).Call(args)
	}
}
