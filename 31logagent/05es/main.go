package main

import (
	"context"
	"fmt"
	"github.com/olivere/elastic/v7"
)

type Person struct {
	Name string `json:"name"`
	Age int `json:"age"`
	Married bool `json:"married"`
}

var client *elastic.Client

func init(){
	var err error
	// 关闭嗅探集群
	client,err = elastic.NewClient(elastic.SetURL("http://192.168.0.110:9200"), elastic.SetSniff(false))
	if err!=nil{
		panic(err)
	}
	fmt.Println("连接成功")
}

func main(){
	s1 := Person{
		Name: "张三",
		Age: 12,
		Married: true,
	}

	indexRes,err := client.Index().Index("person").BodyJson(s1).Do(context.Background())
	if err!=nil{
		panic(err)
	}
	fmt.Println(indexRes.Index,indexRes.Id,indexRes.Result,indexRes.SeqNo)
}
