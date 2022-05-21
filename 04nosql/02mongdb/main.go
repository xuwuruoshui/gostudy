package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main(){
	r := gin.Default()
	
	r.GET("/test", func(c *gin.Context){
		
		c.Set("data",gin.H{"haha":"11111"})
		c.Set("msg","hahahha")
	}, func(c *gin.Context) {
		data ,_:=c.Get("data")
		msg ,_:=c.Get("msg")
		res := gin.H{
			"data": data,
			"msg":  msg,
		}

		c.JSON(http.StatusOK,res)
	})
	
	r.Run("127.0.0.1:8080")
}
