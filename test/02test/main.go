package main

import "fmt"

type Abc interface {
	play(name string)
}

type Person struct {
	name string
}

func (p *Person) play(name string)  {
	fmt.Println(p.name)
	p.name = name
	fmt.Println(p.name)
}

func main() {
	person :=Person{
		name: "aaa",
	}
	var ggg Abc = &person
	fmt.Printf("%T",ggg)
	ggg.play("hhah")
 }
