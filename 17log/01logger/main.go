package main

import "log"

/**
* @creator: xuwuruoshui
* @date: 2021-10-29 15:46:08
* @content: 日志
 */

func main() {

	// Print系列(Print|Printf|Println）
	// Fatal系列（Fatal|Fatalf|Fatalln）
	// Panic系列（Panic|Panicf|Panicln）

	// 这些日志打印会自动的提供时间信息
	log.Println("日志1~~~~~")
	str := "haha"
	log.Printf("日志2%s", str)
	// 执行此条命令,直接结束
	log.Fatal("日志3~~~~~")

	// 运行这条命令,直接抛异常
	log.Panic("日志4~~~~~")
}
