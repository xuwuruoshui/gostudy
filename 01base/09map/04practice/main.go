package main

import (
	"fmt"
	"strings"
)

/**
* @creator: xuwuruoshui
* @date: 2021-10-06 22:32:27
* @content: 统计一个字符串每个单词出现的次数
 */

func main(){
	str := "how do you do"
	strSlice := strings.Split(str," ")
	wordsNum := make(map[string]int)

	for _,value := range strSlice {
		wordsNum[value]++
	}

	for key,value := range wordsNum {
		fmt.Println(key,value)
	}
	
}