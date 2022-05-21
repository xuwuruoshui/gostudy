package main

import "fmt"

func main() {
	num := 3
	fmt.Println(factorial((num)))
}

func factorial(num int) int {
	if(num==1){
		return num
	}
	return num * factorial(num-1)
}