package main

import "fmt"

/**
* @creator: xuwuruoshui
* @date: 2021-12-18 16:12:52
* @content: 选择模式
* 传参可以随意传，添加新字段时,原有构造函数不需要改变,只需要多添加一个函数就可以了
 */

type Person struct{
	name string
	age uint8
	sex uint8
	phone string
	email string

}

type PersonFuc func(*Person)

// 通常写法
func NewPerson(params ... func(*Person))*Person{
	p := &Person{}
	for _,fn :=range params{
		fn(p)
	}
	return p
}

func setName(name string) func(*Person){
	return func(p *Person) {
		p.name = name
	}
}

func setAge(age uint8) func(*Person){
	return func(p *Person) {
		p.age = age
	}
}


func setSex(sex uint8) func(*Person){
	return func(p *Person) {
		p.sex = sex
	}
}

func setPhone(phone string) func(*Person){
	return func(p *Person) {
		p.phone = phone
	}
}

func setEmail(email string) func(*Person){
	return func(p *Person) {
		p.email = email
	}
}




func main(){
	person := NewPerson(setName("haha"),setAge(12))
	fmt.Println(person)
}

