package main

import (
	"35elasticsearch/01normal/es"
	"encoding/json"
	"fmt"
)

func main() {

	// 添加mapping
	//addMapping()

	// 添加
	//add()

	// 修改
	//update()

	// 查询
	search()

	//删除
	//delete()

}

type ResData struct {
	Title string
	Id    int
}

func delete() {
	es.Delete("demo", "3")
}

func search() {
	query := es.E{
		"match": es.E{
			"title": "日照",
		},
	}
	res, _ := es.Search("demo", query, 0, 10, []map[string]string{})
	var resData []ResData

	for _, v := range res.Hits.Hits {
		var data ResData
		json.Unmarshal(v.Source, &data)
		resData = append(resData, data)
	}
	fmt.Println("总数: ", res.Hits.Total.Value)
	fmt.Println(resData)
}

func update() {
	editContent := es.E{
		"content": "日照香炉生紫烟",
	}
	es.Update("demo", "1", editContent)
}

func add() {
	body := es.E{
		"id":    4,
		"title": "日出江花红胜火",
	}
	es.Add("demo", "4", body)
}

func addMapping() {
	mapping := es.E{
		"properties": es.E{
			"id": es.E{
				"type": "integer",
			},
			"title": es.E{
				"type":            "text",
				"analyzer":        "ik_max_word",
				"search_analyzer": "ik_max_word",
			},
		},
	}
	es.AddMapping("demo", mapping)
}
