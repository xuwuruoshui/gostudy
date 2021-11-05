package main

import "fmt"

/**
* @creator: xuwuruoshui
* @date: 2021-10-06 22:32:27
* @content: 结构体,指针
*/

func main() {

	// 1.结构体是值类型
	user := User{"haha", 20, 1}
	printUser(user)
	fmt.Println(user)
	printUser1(&user)
	fmt.Println(user)

	// 2.new关键字
	user1 := new(User)
	fmt.Printf("%T\n",user1)
	fmt.Println(user1)
	
}

func printUser1(user *User) {
	// (*user)和user都能写
	(*user).age = 25
	// 自动根据指针找到对应的变量
	user.name="aaa"
}

func printUser(user User) {
	user.age = 25
}

type User struct {
	name string
	age  int8
	id   int
}
