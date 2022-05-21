package main

import (
	"os"
	"text/template"
)

/**
* @creator: xuwuruoshui
* @date: 2022-05-21 23:44:51
* @content: 模版语法
 */
type Content struct{
	Title string
	Paragraphs string
	Href string
	Hrefname string
}


func main(){
	data := &Content{
		Title: "haha",
		Paragraphs: "yyy",
		Href: "https://www.baidu.com",
		Hrefname: "百度",
	}

	t, _ := template.ParseFiles("./test.html")

	

	f, _ := os.OpenFile("./test1.html", os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0666)
	t.Execute(f,data)


	
}