package main

import (
	"encoding/json"
	"fmt"
)

/**
* @creator: xuwuruoshui
* @date: 2021-10-16 22:57:39
* @content: struct与json互转
 */

type User struct {
	// json包中规定结构体必须大写,如果想用小写就只能在后面限定
	Id   int    `json:"id" db:"id" ini:"id"`
	Name string `json:"name" db:"name" ini:"name"`
}

func main() {
	u1 := User{1, "zhangsan"}

	// 1.结构体转json
	// data为二进制
	data, err := json.Marshal(u1)
	if err != nil {
		fmt.Println("转换错误")
	}
	fmt.Printf("%s\n", string(data))

	// 2.json转结构体
	var u2 User
	// 参数传二进制,结构体地址
	json.Unmarshal(data, &u2)
	fmt.Println(u2)

}
