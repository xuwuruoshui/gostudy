package main

import "fmt"

/**
* @creator: xuwuruoshui
* @date: 2021-10-06 22:32:27
* @content: 回文数判断
 */

func main(){
	
	str := "abbcbbaa"
	for i := 0; i < (len(str)-1)/2; i++ {
		if str[i]!=str[len(str)-1-i]{
			fmt.Println("不是回文数")
			return
		}
	}
	fmt.Println("是回文数")
}