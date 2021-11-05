package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println(fibonacci(i))
	}

}

func fibonacci(num int) int {
	if num <= 1 {
		return num
	}
	// 当前传参等于，上上次的传参和上次的传参
	return fibonacci(num-2) + fibonacci(num-1)
}
