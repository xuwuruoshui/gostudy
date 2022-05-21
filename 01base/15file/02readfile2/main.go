package main

import (
	"fmt"
	"io"
	"os"
)

/**
* @creator: xuwuruoshui
* @date: 2021-10-28 13:52:26
* @content: 循环读取文件
 */

func main(){
	// 1.打开文件
	file,err := os.Open("./a.txt")
	if err!=nil{
		fmt.Println("文件不能为空")
	}
	// 2. 文件处理后，最后关闭文件
	defer file.Close()


	// 3.定义临时文件变量，定义最终的文件变量
	var tmp = make([]byte,5)
	var content []byte

	for{
		n,err := file.Read(tmp)
		if err==io.EOF{
			fmt.Println("文件读取完成!!!")
			break
		}
		
		if err!=nil{
			fmt.Println("文件读取错误:",err)
			return
		}
		content = append(content, tmp[:n]...)
	}
	
	fmt.Println(string(content))
}