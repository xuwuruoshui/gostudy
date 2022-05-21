package main

import "fmt"

/**
* @creator: xuwuruoshui
* @date: 2021-10-16 17:34:49
* @content: 结构体中内嵌结构体
 */

type User struct {
	id   int
	name string
	Address
	heal Health
}

type Address struct {
	area     string
	city     string
	province string
}

type Health struct {
	weight int
	heigh  int
}

func newUse(id int, name string, address Address,heal Health) *User {
	return &User{
		id,
		name,
		address,
		heal}
}

func newAddress(city string, province string, area string) *Address {
	return &Address{
		city,
		province,
		area}
}

func newHealth(weight int, heigh int) *Health {
	return &Health{weight, heigh}
}

func main() {
	heal := newHealth(120,180)
	addr := newAddress("成都", "四川", "高新区")
	u1 := newUse(1, "haha", *addr,*heal)
	fmt.Println(u1)
	// 匿名的方式可以直接访问,也可以使用全名的方式
	// 但是如果嵌套的其他结构体中有相同字段, 也会出问题
	fmt.Println(u1.id, u1.Address.city, u1.province)
	fmt.Println(u1.heal.heigh)
}
