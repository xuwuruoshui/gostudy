package main

import (
	"fmt"
	"io"
	"os"
)

/**
* @creator: xuwuruoshui
* @date: 2021-11-13 17:09:42
* @content: 文件中间插入
 */

func main() {

	// 1.读取文件
	file, err := os.OpenFile("./a.txt", os.O_RDWR, 0664)
	if err != nil {
		fmt.Println("open file error:", err)
	}

	// 2.读3个字节,a、\r、\n共三个字节
	tempFile, err := os.OpenFile("./a.tmp", os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0664)
	if err != nil {
		fmt.Println("create file error", err)
	}

	value1 := make([]byte, 3)
	n, _ := file.Read(value1)
	tempFile.Write(value1[:n])

	// 3.偏移量设置为3,
	_, err = file.Seek(3, io.SeekStart)
	if err != nil {
		fmt.Println("read file error", err)
	}

	// 3.c写进去
	tempFile.Write([]byte("c\r\n"))

	for {
		value2 := make([]byte, 1024)
		n, err := file.Read(value2)
		if err == io.EOF {
			tempFile.Write(value2[:n])
			break
		}
		if err != nil {
			fmt.Println("file read error:", err)
		}
		tempFile.Write(value2[:n])
	}

	file.Close()
	tempFile.Close()

	// 4.重命名覆盖
	err = os.Rename("./a.tmp", "./a.txt")
	if err != nil {
		fmt.Println("rename error:", err)
	}
}
