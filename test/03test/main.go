package main

import (
	"io"
	"log"
	"os"
)

func main() {
	bg,err := os.Open("./bg10.jpg")
	if err!=nil{
		log.Panicln("文件打开错误",err)
	}
	defer bg.Close()

	bg1,err := os.OpenFile("./test.jpg",os.O_CREATE|os.O_TRUNC|os.O_WRONLY,0666)

	if err!=nil{
		log.Panicln("文件创建失败",err)
	}
	io.Copy(bg1,bg)
}
