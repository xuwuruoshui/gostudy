package main

import (
	"fmt"
	"os"
)

/**
* @creator: xuwuruoshui
* @date: 2021-10-28 15:01:27
* @content: 写文件
 */

func main() {

	// 模式	含义
	// os.O_WRONLY	只写
	// os.O_CREATE	创建文件
	// os.O_RDONLY	只读
	// os.O_RDWR	读写
	// os.O_TRUNC	清空(写文件时直接清空)
	// os.O_APPEND	追加(在原来的基础上追加)
	// 最后一个参数和linux文件系统一样,4读,2写,1可执行

	// 1.创建文件,并设置相关权限参数
	file, err := os.OpenFile("a.txt", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 666)
	if err != nil {
		fmt.Println("open file failed", err)
		return
	}

	defer file.Close()

	str := "Hello Golang!!!\n"
	// 写入字节切片
	file.Write([]byte(str))
	// 写入字符串
	file.WriteString("Let's go!!!\n")

}
