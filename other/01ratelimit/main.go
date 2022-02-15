package main

import (
	"net/http"

	"golang.org/x/time/rate"
)

/**
* @creator: xuwuruoshui
* @date: 2022-02-14 22:12:23
* @content: 限流器
 */
var limiter = rate.NewLimiter(1, 5)

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8000", nil)

}
func handler(w http.ResponseWriter, r *http.Request) {
	// 一秒填充1个，初始5个

	if !limiter.Allow() {
		w.Write([]byte(`{"error":"to many request"}`))
	} else {
		w.Write([]byte(`{"msg":"haha"}`))
	}
}
