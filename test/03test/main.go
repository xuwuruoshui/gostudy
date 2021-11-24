package main

import "fmt"



func main() {
	value:= int64(7416324704784844908)
	var result int64
	for{
		result +=value%10
		if value/10<10{
			fmt.Println(result)
			break
		}
		value = value/10		
	}
}
