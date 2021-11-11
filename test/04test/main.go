package main

import (
	"encoding/json"
	"fmt"
)

/**
* @creator: xuwuruoshui
* @date: 2021-11-04 16:39:10
* @content: service
 */

 type User struct {
	Id       int `json:"id,omitempty"`
	Age      int `json:"age,omitempty"`
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
}

func main() {

	content,err := json.Marshal(User{Username: "zhangsna",Password: "root"})
	if err!=nil{
		fmt.Println(err)
	}

	fmt.Println(string(content))
}
