package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

/**
* @creator: xuwuruoshui
* @date: 2021-10-28 14:10:07
* @content: bufio读取文件
 */

func main() {
	// 1.基本文件读取操作
	file,err := os.Open("./a.txt")
	if err!=nil{
		fmt.Println("open file faild,err:",err)
		return
	}
	defer file.Close()

	// 2.开启bufio
	reader := bufio.NewReader(file)

	for{
		// 3.读取一行，\n此处为字符
		line,err := reader.ReadString('\n')
		if err==io.EOF{
			if len(line)!=0{
				fmt.Println(line)
			}
			fmt.Println("文件读完了")
			break
		}

		if err!=nil{
			fmt.Println("读取文件失败:",err)
			return
		}
		fmt.Println(line)
	}
}
