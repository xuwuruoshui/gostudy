package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main(){
	r := gin.Default()
	r.Use(Auth())

	r.GET("/login/:username/:password", func(c *gin.Context) {
		username := c.Param("username")
		password := c.Param("password")

		if username=="admin" && password=="admin"{
			c.SetCookie("username",username,60,"/","localhost",false,true)
		}else{
			c.JSON(http.StatusUnauthorized,gin.H{"status":http.StatusUnauthorized,"msg":"用户名密码错误"})
		}

	})

	r.GET("/home", func(c *gin.Context) {
		name ,err :=c.Cookie("username")
		if err!=nil{
			panic(err)
		}
		c.JSON(http.StatusOK,gin.H{"status":http.StatusOK,"msg":"welcome","username":name})
	})

	r.Run(":8080")
}

func Auth()  gin.HandlerFunc{
	return func(c *gin.Context) {
		if c.FullPath()== "/login/:username/:password"{
			return
		}

		name ,err :=c.Cookie("username")
		if err!=nil || name!="admin"{
			c.JSON(http.StatusUnauthorized,gin.H{"status":http.StatusUnauthorized,"msg":"未登录"})
			c.Abort()
		}
	}
}