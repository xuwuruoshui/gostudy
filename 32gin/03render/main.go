package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/testdata/protoexample"
	"net/http"
)

func main(){
	r := gin.Default()

	// gin.H本质为 map[string]interface{}
	r.GET("/someJson", func(c *gin.Context) {
		c.JSON(http.StatusOK,gin.H{"message":"hey","status":http.StatusOK})
	})

	r.GET("/moreJson", func(c *gin.Context) {
		var msg struct{
			Name string `json:"user"`
			Message string
			Number int
		}
		msg.Name = "zhangsan"
		msg.Message = "nb"
		msg.Number = 666

		// 返回是Name字段会变为user
		c.JSON(http.StatusOK,msg)
	})

	r.GET("/someXML", func(c *gin.Context) {
		c.XML(http.StatusOK,gin.H{"message":"haha","status":http.StatusOK})
	})

	r.GET("/someYAML", func(c *gin.Context) {
		c.YAML(http.StatusOK,gin.H{"message":"hey","status":http.StatusOK})
	})

	r.GET("/someProtoBuf", func(c *gin.Context) {
		reps := []int64{int64(1),int64(2)}
		lable := "test"
		data := &protoexample.Test{
			Label: &lable,
			Reps: reps,
		}
		c.ProtoBuf(http.StatusOK,data)
	})

	// html模板渲染
	r.LoadHTMLGlob("template/*")
	//r.LoadHTMLFiles("template/index.tmpl")
	r.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK,"index.tmpl",gin.H{
			"title":"Main website",
		})
	})
	r.Run(":8080")
}
