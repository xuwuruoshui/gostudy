package main

import (
	"fmt"
	"time"
)

/**
* @creator: xuwuruoshui
* @date: 2021-10-16 20:30:33
* @content: 模拟继承
 */

type Animal struct {
	name string
}

func (animal Animal) move() {
	fmt.Println(animal.name, "is moving")
}

type Dog struct {
	category string
	Animal
}

func (dog Dog) bark() {
	fmt.Println("the category of ", dog.category, "is barking who name is", dog.name)
}

func main() {
	dog1 := Dog{"土狗", Animal{"大黄"}}
	dog1.bark()
	dog1.move()
	time.Sleep(5000)
}
