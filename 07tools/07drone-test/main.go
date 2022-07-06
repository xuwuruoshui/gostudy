package main

import (
	"fmt"
	"net/http"
)

/**
* @creator: xuwuruoshui
* @date: 2022-07-06 15:51:58
* @content:
 */

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":4567", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("ff")
	w.Write([]byte("ffff"))
}
