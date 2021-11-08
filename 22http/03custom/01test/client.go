package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

/**
* @creator: xuwuruoshui
* @date: 2021-11-08 10:48:40
* @content: 自定义client
 */

func main(){
	client := &http.Client{}
	req,err := http.NewRequest("Get","https://www.bilibili.com/",nil)
	if err!=nil{
		fmt.Println("转换失败:",err)
	}
	req.Header.Add("User-Agent","Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/93.0.4577.63 Safari/537.36")

	resp,err := client.Do(req)
	if err!=nil{
		fmt.Println("响应错误:",err)
	}
	content,err := ioutil.ReadAll(resp.Body)
	if err!=nil{
		fmt.Println("读取数据失败:",err)
	}
	fmt.Println(string(content))
}