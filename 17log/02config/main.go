package main

import (
	"fmt"
	"log"
	"os"
)

/**
* @creator: xuwuruoshui
* @date: 2021-10-29 16:07:56
* @content: 配置日志
 */

func main(){
	// 具体配置看源码
	// 1.打印一个具体日期到微秒，长文件的日志
	log.SetFlags(log.Llongfile | log.Lmicroseconds | log.Ldate)
	log.Println("这是一条很普通的日志")

	// 2.打印一个短文件
	log.SetFlags(log.Lshortfile | log.Lmicroseconds | log.Ldate)
	log.Println("短一点的日志")

	// 3.加一个前缀，放日志最前面
	log.SetPrefix("[666]")
	log.Println("普通的日志")

	// 4.前缀移到具体信息前
	log.SetFlags(log.Lmsgprefix|log.Lshortfile | log.Lmicroseconds | log.Ldate)
	log.Println("普通的日志")

	// 5.日志设置输出位置(配置日志输出位置，通常放在init函数中初始化)
	logFile,err := os.OpenFile("./a.log",os.O_CREATE | os.O_WRONLY | os.O_APPEND,0644)
	if err!=nil{
		fmt.Println("文件创建失败")
	}

	log.SetOutput(logFile)
	log.SetFlags(log.Llongfile | log.Lmicroseconds | log.Ldate)
	log.Println("===============启动link start==============")
	log.SetPrefix("xuwuruoshui:")
	log.Println("===================获取数据=================")
	log.Println("==============获取数据完毕=================")


	// 6.
}