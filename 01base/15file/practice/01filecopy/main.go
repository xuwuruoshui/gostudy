package main

import (
	"fmt"
	"io/ioutil"
)

/**
* @creator: xuwuruoshui
* @date: 2021-10-28 16:24:26
* @content: 文件拷贝
 */

func main(){
	oldImg,err := ioutil.ReadFile("./bg10.jpg")
	if err!=nil{
		fmt.Println("读取出错了:",err)
	}
	
	err = ioutil.WriteFile("./haha.jpg",oldImg,0666)
	if err!=nil{
		fmt.Println("写入失败:",err)
	}
	fmt.Println("复制成功!!!")
}


