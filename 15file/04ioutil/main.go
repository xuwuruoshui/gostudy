package main

import (
	"fmt"
	"io/ioutil"
)

/**
* @creator: xuwuruoshui
* @date: 2021-10-28 14:56:25
* @content: ioutil读操作
 */

func main(){
	// 一步直接读取文件内容
	content,err := ioutil.ReadFile("./a.txt")
	if err!=nil{
			fmt.Println("read file failed,err:",err)
			return
	}
	fmt.Println(string(content))
}