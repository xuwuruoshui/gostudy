package main

import "fmt"

func main() {

	ch := make(chan int,2)
	ch <- 1
	ch <- 2

	x,y :=<-ch,<-ch
	fmt.Println(x)
	fmt.Println(y)
}
