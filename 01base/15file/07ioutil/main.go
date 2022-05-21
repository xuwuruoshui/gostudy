package main

import (
	"fmt"
	"io/ioutil"
)

/**
* @creator: xuwuruoshui
* @date: 2021-10-28 15:33:09
* @content: ioutil写操作
 */

func main() {

	// 直接写
	str := "abcdefg"
	err := ioutil.WriteFile("./a.txt", []byte(str), 0666)
	if err != nil {
		fmt.Println("write file failed,err", err)
		return
	}
}
