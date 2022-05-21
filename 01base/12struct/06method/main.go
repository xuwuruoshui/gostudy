package main

import "fmt"

/**
* @creator: xuwuruoshui
* @date: 2021-10-06 22:32:27
* @content: 方法
 */

// Go语言中如果标识符首字母是大写，就表示对外部包可见(暴露的，公有的)
// Dog 狗
type Dog struct {
	name string
}

func newDog(name string) *Dog {
	return &Dog{name: name}
}

// Dog下的方法
func (d Dog) bark() {
	fmt.Println(d.name, ":汪汪汪~~")
}

type Person struct {
	name string
	age  int
}

func newPerson(name string, age int) *Person {
	return &Person{name: name, age: age}
}

func (p Person) happyNewYear() {
	p.age++
}

func (p *Person) happyNewYear2() {
	p.age++
}

// 自定义类型
type myInt int

func (myint myInt) test(){
	fmt.Println("Custom method")
}


func main() {
	yellowDog := newDog("大黄")
	yellowDog.bark()

	thePerson := newPerson("zhansan", 18)

	// 1.值传递
	thePerson.happyNewYear()
	fmt.Println(thePerson.age)

	// 2.引用传递,happyNewYear2才能修改+1,happyNewYear是拷贝了一份新的Person
	thePerson.happyNewYear2()
	fmt.Println(thePerson.age)

	// 3.自定义类型方法
	// 基本数据类型不能使用方法
	var yourInt myInt = 10
	yourInt.test()

	
}
