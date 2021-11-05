package main

import (
	"fmt"
	"time"
)

/**
* @creator: xuwuruoshui
* @date: 2021-10-06 22:32:27
* @content: time包
 */

func main() {

	// 1.获取当前时间
	now := time.Now()
	fmt.Println(now)
	year := now.Year()
	month := now.Month()
	day := now.Day()
	hour := now.Hour()
	minute := now.Minute()
	second := now.Second()
	// 0:高位用0进行补位,2:宽度为2
	fmt.Printf("%d-%02d-%02d %02d:%02d:%02d\n", year, month, day, hour, minute, second)

	// 2.时间戳
	timestamp1 := now.Unix()
	// 纳秒级
	timestamp2 := now.UnixNano()
	fmt.Printf("current timestamp:%v\n", timestamp1)
	fmt.Printf("current timestamp:%v\n", timestamp2)

	// 3.时间戳转时间
	time1 := time.Unix(timestamp1, 0)
	fmt.Println(time1 == now)
	fmt.Println(time1.Unix() == now.Unix())
	time1 = time.Unix(0, timestamp2)
	fmt.Println(time1.UnixNano() == now.UnixNano())

	// 4.时间间隔单位Duration
	fmt.Println(time.Hour)
	fmt.Println(time.Minute)
	fmt.Println(time.Second)
	fmt.Println(time.Millisecond)
	fmt.Println(time.Microsecond)
	fmt.Println(time.Nanosecond)


}
