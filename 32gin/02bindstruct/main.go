package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type User struct {
	UserName string `form:"username" json:"username" uri:"username" xml:"username" binding:"required"`
	Password string `form:"password" json:"password" uri:"password" xml:"password" binding:"required"`
}

func main(){
	r := gin.Default()

	r.POST("login", func(c *gin.Context) {

		var data User

		// 1.请求体json,解析到结构体中
		//if err:=c.ShouldBindJSON(&data);err!=nil{
		//	// 返回JSON数据
		//	c.JSON(http.StatusBadRequest,gin.H{"error":"数据转换异常"})
		//	panic(err)
		//}

		// 2.Bind根据请求头的content-type自动推断，默认x-www-form-urlencoded
		if err:=c.Bind(&data);err!=nil{
			// 返回JSON数据
			c.JSON(http.StatusBadRequest,gin.H{"error":"数据转换异常"})
			panic(err)
		}

		// 判断用户名密码是否正确
		if data.UserName!="root" || data.Password!="root"{
			c.JSON(http.StatusBadRequest,gin.H{"msg":"用户名或密码错误","status":302})
			return
		}

		c.JSON(http.StatusOK,gin.H{"msg":"登陆成功","status":200})
	})

	// 3.uri
	r.POST("login1/:username/:password", func(c *gin.Context) {

		var data User

		if err:=c.ShouldBindUri(&data);err!=nil{
			c.JSON(http.StatusBadRequest,gin.H{"error":"数据转换异常"})
			panic(err)
		}

		// 判断用户名密码是否正确
		if data.UserName!="root" || data.Password!="root"{
			c.JSON(http.StatusBadRequest,gin.H{"msg":"用户名或密码错误","status":302})
			return
		}

		c.JSON(http.StatusOK,gin.H{"msg":"登陆成功","status":200})
	})
	r.Run(":8080")
}
