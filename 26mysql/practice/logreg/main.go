package main

import (
	"gostudy/25mysql/04practice/logreg/controller"
	"net/http"
)

/**
* @creator: xuwuruoshui
* @date: 2021-11-10 10:46:06
* @content: 简单的登录注册功能
 */

func main() {
	http.HandleFunc("/index", filter(controller.Index))
	http.HandleFunc("/login", filter(controller.Login))
	http.HandleFunc("/reg", filter(controller.Reg))
	http.ListenAndServe(":8080", nil)
}

func filter(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5500")                   //允许访问所有域
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")                           //header的类型
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, PUT, OPTIONS, DELETE, PATCH") //header的类型
		w.Header().Set("Access-Control-Max-Age", "3600")                                         //header的类型
		w.Header().Set("Access-Control-Allow-Headers", "*")                                      //header的类型

		// options检测跨域和header,一旦跨域访问r中什么都得不到
		if r.Method == "OPTIONS" {
			return
		}
		h(w, r)
	}
}
