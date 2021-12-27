package main

import "fmt"



func main() {
		slice := make([]int,0,5)
		test(&slice)
		fmt.Println(slice)
}

 func test(s *[]int){
	*s = append(*s,10)
 }
