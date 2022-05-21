package main

import (
	"fmt"
	"time"
)

/**
* @creator: xuwuruoshui
* @date: 2021-10-06 22:32:27
* @content: 时间的操作
 */

func main() {
	// 1.Add当前时间加一个小时
	now := time.Now()
	fmt.Println(now.Add(time.Hour))

	// 2.Sub求两个时间的差
	cusTime := time.Date(2021, time.October, 5, 0, 0, 0, 0, time.Local)
	fmt.Println(now.Sub(cusTime).Seconds())

	// 3.Equal判断两个时间是否相同,时区也在判断的范围内
	fmt.Println(time.Now().Equal(now))

	// 4.Before/After 某个时间 在 另外一个时间 之前和之后
	fmt.Println(now.Before(time.Now()))
	fmt.Println(now.After(time.Now()))

	// 5.定时器 本质是一个通道
	// 每5秒执行一次
	// ticker := time.Tick(time.Second*5)
	// for i := range ticker {
	// 	fmt.Println(i)
	// }

	// 6.格式化
	// 口诀2006 1 2 3 4 5
	// 24小时制
	fmt.Println(now.Format("2006-01-02 15:04:05.000 Mon Jan"))
	// 12小时制
	fmt.Println(now.Format("2006-01-02 03:04:05.000 PM Mon Jan"))
	fmt.Println(now.Format("2006/01/02 15:04"))
	fmt.Println(now.Format("15:04 2006/01/02"))
	fmt.Println(now.Format("2006/01/02"))

	// 7.解析字符串格式时间
	// 加载时区
	local, err := time.LoadLocation("Asia/Shanghai")
	if nil != err {
		fmt.Println(err)
		return
	}

	// 8.按照指定格式指定时区解析字符串时间
	timeObj, err := time.ParseInLocation("2006-01-02 15:04:05", "2021-10-01 11:11:11", local)
	if nil != err {
		fmt.Println(err)
		return
	}

	fmt.Println(timeObj)

}
