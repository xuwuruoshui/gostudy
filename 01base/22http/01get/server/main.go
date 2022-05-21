package main

import (
	"fmt"
	"net/http"
)

func get(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	data := r.URL.Query()
	fmt.Println(data.Get("name"))
	fmt.Println(data.Get("age"))
	answer := `{"status","ok"}`
	// 往客户端写数据
	w.Write([]byte(answer))
}

func main() {
	http.HandleFunc("/get", get)
	http.ListenAndServe("127.0.0.1:9090", nil)
}
