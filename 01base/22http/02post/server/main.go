package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

/**
* @creator: xuwuruoshui
* @date: 2021-11-08 10:20:26
* @content: post客户端
 */

func greet(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	// 1. 请求类型是application/x-www-form-urlencode时解析form数据
	r.ParseForm()
	// 可以直接遍历map
	fmt.Println(r.PostForm)
	fmt.Println(r.PostForm.Get("name"), r.PostForm.Get("age"))

	// 2. 请求类型是application/json时从r.Body读取数据
	content, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println("读取数据失败", err)
	}
	fmt.Println(string(content))

	w.Write([]byte(`{"status","ok"}`))
}

func main() {
	http.HandleFunc("/post", greet)
	http.ListenAndServe("127.0.0.1:9090", nil)
}
