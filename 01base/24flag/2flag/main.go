package main

import (
	"flag"
	"fmt"
	"time"
)

/**
* @creator: xuwuruoshui
* @date: 2021-11-08 17:45:31
* @content:
 */

func main() {
	var name string
	var age int
	var married bool
	var delay time.Duration

	flag.StringVar(&name, "name", "张三", "姓名")
	flag.IntVar(&age, "age", 18, "姓名")
	flag.BoolVar(&married, "married", false, "婚否")
	flag.DurationVar(&delay, "delay", 1, "延时间隔")

	fmt.Println(name)
	fmt.Println(age)
	fmt.Println(married)
	fmt.Println(delay)
	// 布尔值不支持 -param value, 只支持 -param="true"
	flag.Parse()
	fmt.Println("====================")
	fmt.Println(name, age, married, delay)

	// 运行 go run .\some_test.go -name zhangsan -age 12 -married=true -delay 3s aaaa bbbb
	// 返回命令行参数后的其他参数
	fmt.Println(flag.Args())
	// 返回命令行参数后的其他参数个数
	fmt.Println(flag.NArg())
	// 返回使用命令行参数个数
	fmt.Println(flag.NFlag())
}
