package main

import (
	"fmt"
	"time"
)

/**
* @creator: xuwuruoshui
* @date: 2021-10-06 22:32:27
* @content: 获取当前时间，格式化输出为2017/06/19 20:30:05格式
 */

func main(){
	now := time.Now()
	fmt.Printf("%T\n",now.Format("2006/01/02 15:04:05"))
	fmt.Println(now.Format("2006/01/02 15:04:05"))

}