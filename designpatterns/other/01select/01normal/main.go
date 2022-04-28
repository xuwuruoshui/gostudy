package main

import "fmt"

/**
* @creator: xuwuruoshui
* @date: 2021-12-18 16:12:52
* @content: 普通写法
 */

type Person struct{
	name string
	age uint8
	sex uint8
	phone string
	email string

}

// 通常写法
func NewPerson(name,phone,email string,age,sex uint8)*Person{
	return &Person{
		name:name,
		age:age,
		sex:sex,
		phone:phone,
		email:email,
	}
}

func main(){
	person := NewPerson("张三","13234567890","abc@fff.com",12,13)
	fmt.Println(person)
	
}

