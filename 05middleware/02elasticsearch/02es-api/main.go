package main

import (
	"bytes"
	"context"
	"encoding/json"
	"es/model"
	"fmt"
	"github.com/elastic/go-elasticsearch/v7"
	"github.com/elastic/go-elasticsearch/v7/esapi"
	"log"
	"net/http"
)

type User struct {
	Id  string `json:"_id"`
	Age int    `json:"age"`
}

var EsClient *elasticsearch.Client

func init() {

	config := elasticsearch.Config{
		Addresses: []string{
			"http://localhost:9200",
		},
	}

	es, err := elasticsearch.NewClient(config)
	if err != nil {
		log.Fatalf("Error creating the client: %s", err)
	}

	_, err = es.Info()
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}

	EsClient = es

}

func main() {
	// 1. 添加/修改
	// Index("user", "2", map[string]interface{}{"name": "ffff", "age": 12})
	// IndexEspi("user", "3", map[string]interface{}{"name": "hahahah", "age": 25, "sex": 2})

	// 2.查询
	// 条件查询
	//Search("user", model.Es{"query": model.Es{
	//	"match": model.Es{
	//		"_id": "3",
	//	},
	//}})
	// 查询全部
	//Search("user", model.Es{})

	// 3.删除
	//Delete("user", "2")

	// 4.通过条件删除
	//DeleteByQuery([]string{"user"}, model.Es{
	//	"query": model.Es{
	//		"range": model.Es{"age": model.Es{
	//			"lte": 16, "gt": 12,
	//		}},
	//	},
	//})

	// 5. sql
	SearchEsapiSql(model.Es{"query": "select * from user where age<=16"})
}

func Index(index, idx string, doc map[string]interface{}) {
	//index:="my_index_name_v1"
	res, err := EsClient.Info()
	fmt.Println(res, err)
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}
	var buf bytes.Buffer

	if err = json.NewEncoder(&buf).Encode(doc); err != nil {
		fmt.Println(err, "Error encoding doc")
		return
	}
	res, err = EsClient.Index(
		index,                              // Index name
		&buf,                               // Document body
		EsClient.Index.WithDocumentID(idx), // Document ID
		// Document ID
		EsClient.Index.WithRefresh("true"), // Refresh
	)
	//res, err = EsClient.Create(index, idx, &buf)
	if err != nil {
		fmt.Println(err, "Error create response")
	}
	defer res.Body.Close()
	fmt.Println(res.String())
	log.Println(res)
}

// 添加 POST
// struct 类型允许使用更实际的方法，您可以在其中创建一个新结构，将请求配置作为字段，并使用上下文和客户端作为参数调用 Do() 方法：

func IndexEspi(index, idx string, doc map[string]interface{}) {
	//index:="my_index_name_v1"
	res, err := EsClient.Info()
	fmt.Println(res, err)
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}
	var buf bytes.Buffer
	if err = json.NewEncoder(&buf).Encode(doc); err != nil {
		fmt.Println(err, "Error encoding doc")
		return
	}

	req := esapi.IndexRequest{
		Index:      index,  // Index name
		Body:       &buf,   // Document body
		DocumentID: idx,    // Document ID
		Refresh:    "true", // Refresh
	}

	res, err = req.Do(context.Background(), EsClient)
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}
	defer res.Body.Close()
	fmt.Println(res.String())
	log.Println(res)
}

// 查询

func Search(index string, query map[string]interface{}) {
	res, err := EsClient.Info()
	if err != nil {
		fmt.Println(err, "Error getting response")
	}
	//fmt.Println(res.String())
	// search - highlight
	var buf bytes.Buffer
	//query := map[string]interface{}{
	//	"query": map[string]interface{}{
	//		"match": map[string]interface{}{
	//			"title": title,
	//		},
	//	},
	//	"highlight": map[string]interface{}{
	//		"pre_tags":  []string{"<font color='red'>"},
	//		"post_tags": []string{"</font>"},
	//		"fields": map[string]interface{}{
	//			"title": map[string]interface{}{},
	//		},
	//	},
	//}
	if err := json.NewEncoder(&buf).Encode(query); err != nil {
		fmt.Println(err, "Error encoding query")
	}
	// Perform the search request.
	res, err = EsClient.Search(
		EsClient.Search.WithContext(context.Background()),
		EsClient.Search.WithIndex(index),
		EsClient.Search.WithBody(&buf),
		EsClient.Search.WithTrackTotalHits(true),
		EsClient.Search.WithFrom(0),
		EsClient.Search.WithSize(10),
		//EsClient.Search.WithSort("time:desc"),
		EsClient.Search.WithPretty(),
	)
	if err != nil {
		fmt.Println(err, "Error getting response")
	}
	defer res.Body.Close()
	fmt.Println(res.String())
}

// 删除
func Delete(index, idx string) {
	//index:="my_index_name_v1"
	res, err := EsClient.Info()
	fmt.Println(res, err)
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}
	res, err = EsClient.Delete(
		index, // Index name
		idx,   // Document ID
		EsClient.Delete.WithRefresh("true"),
	)
	if err != nil {
		fmt.Println(err, "Error create response")
	}
	defer res.Body.Close()
	fmt.Println(res.String())
	log.Println(res)
}

// 查询删除
func DeleteByQuery(index []string, query map[string]interface{}) {
	res, err := EsClient.Info()
	if err != nil {
		fmt.Println(err, "Error getting response")
	}
	//fmt.Println(res.String())
	// search - highlight
	var buf bytes.Buffer
	//query := map[string]interface{}{
	//	"query": map[string]interface{}{
	//		"match": map[string]interface{}{
	//			"title": title,
	//		},
	//	},
	//	},
	//}
	if err := json.NewEncoder(&buf).Encode(query); err != nil {
		fmt.Println(err, "Error encoding query")
	}
	// Perform the search request.
	res, err = EsClient.DeleteByQuery(
		index,
		&buf,
	)
	if err != nil {
		fmt.Println(err, "Error getting response")
	}
	defer res.Body.Close()
	fmt.Println(res.String())

}

// es api
func SearchEsapiSql(query map[string]interface{}) {
	jsonBody, _ := json.Marshal(query)
	req := esapi.SQLQueryRequest{Body: bytes.NewReader(jsonBody)}
	res, _ := req.Do(context.Background(), EsClient)
	defer res.Body.Close()
	fmt.Println(res.String())
}

// http方式
func SearchHttp(method, url string, query map[string]interface{}) {
	jsonBody, _ := json.Marshal(query)
	req, _ := http.NewRequest(method, url, bytes.NewReader(jsonBody))
	req.Header.Add("Content-type", "application/json")
	res, err := EsClient.Perform(req)
	if err != nil {
		return
	}
	defer res.Body.Close()
	buf := new(bytes.Buffer)
	buf.ReadFrom(res.Body)
	fmt.Println(buf.String())
}
