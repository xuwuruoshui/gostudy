package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

/**
* @creator: xuwuruoshui
* @date: 2021-11-08 09:47:24
* @content: get请求
 */

func main(){
	// no_param_req()
	param_req()
}

func param_req(){
	apiUrl :="http://127.0.0.1:9090/get"
	// 1.设置参数
	data := url.Values{}
	data.Set("name","张三")
	data.Set("age","18")

	// 2.字符串转url
	u, err := url.ParseRequestURI(apiUrl)
	if err!=nil{
		fmt.Println("转参数失败:",err)
	}
	u.RawQuery = data.Encode()

	// 转换后的路径
	fmt.Println(u.String())

	// 3.发送请求
	resp,err := http.Get(u.String())
	if err!=nil{
		fmt.Println("请求失败:",err)
		return
	}
	defer resp.Body.Close()

	b,err := ioutil.ReadAll(resp.Body)
	if err!=nil{
		fmt.Println("读取失败:",err)
	}
	fmt.Println(string(b))
}

func no_param_req(){
	resp,err := http.Get("https://www.baidu.com/")
	
	if err!=nil{
		fmt.Println("获取数据失败",err)
		return
	}
	defer resp.Body.Close()

	body,err := ioutil.ReadAll(resp.Body)
	if err!=nil{
		fmt.Println("读取数据失败",err)
	}
	fmt.Println(string(body))
}