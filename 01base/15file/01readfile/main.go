package main

import (
	"fmt"
	"io"
	"os"
)

/**
* @creator: xuwuruoshui
* @date: 2021-10-28 13:38:49
* @content: 文件读取
 */

func main() {

	// 1.打开文件
	file, err := os.Open("./a.txt")
	if err != nil {
		fmt.Println("文件打开错误:", err)
	}

	// 2.文件最后关闭流
	defer file.Close()

	// 3.以6字节读取
	var tmp = make([]byte, 128)
	n, err := file.Read(tmp)
	if err == io.EOF {
		fmt.Println("文件为空!!!")
		return
	}

	if err != nil {
		fmt.Println("读取失败!!!")
	}

	fmt.Printf("读取字节数据:%d\n", n)
	fmt.Println(string(tmp[:n]))
}
