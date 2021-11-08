package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

/**
* @creator: xuwuruoshui
* @date: 2021-11-08 10:20:30
* @content: post服务端
 */

func main() {
	// json()
	form()
}

func form(){
	// 1.application/x-www-form-urlencoded
	addr := "http://127.0.0.1:9090/post"
	data := url.Values{}
	data.Set("name","张三")
	data.Set("age","18")
	resp,err := http.PostForm(addr,data)
	if err != nil {
		fmt.Println("发送请求失败:", err)
	}
	defer resp.Body.Close()

	// 3.读取响应数据
	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("读取数据失败:", err)
	}
	fmt.Println(string(content))
}

func json() {
	// 1.url、contentType、data
	addr := "http://127.0.0.1:9090/post"
	contentType := "application/json"
	data := `{"name":"zhangsan","age":10}`

	// 2.发送请求、关闭
	resp, err := http.Post(addr, contentType, strings.NewReader(data))
	if err != nil {
		fmt.Println("发送请求失败:", err)
	}
	defer resp.Body.Close()

	// 3.读取响应数据
	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("读取数据失败:", err)
	}
	fmt.Println(string(content))
}
