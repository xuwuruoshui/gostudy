package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func main(){
	// 1.创建路由(defualt自带了两个中间件Logger和Recovery)
	// gin.New()也能创建，但不带中间件
	router := gin.Default()

	// 2.路径映射对应的方法
	router.GET("/", func(context *gin.Context) {
		context.String(http.StatusOK,"Hello World")
	})

	// 3.路由匹配规则

	// 只能匹配到 /user/xxx
	router.GET("/user/:username", func(context *gin.Context) {
		username := context.Param("username")
		context.String(http.StatusOK,username)
	})

	// 既可以匹配 /user/xxx (默认转为/user/xxx/),也可以匹配到 /user/xxx/xxx
	router.GET("/user/:username/*action", func(context *gin.Context) {
		username := context.Param("username")
		action := context.Param("action")
		context.String(http.StatusOK,username+" is "+action)
	})

	// 参数 /welcome?firstname="xxx"&lastname="xxx" firstname不写则默认为Guest
	router.GET("/welcome", func(context *gin.Context) {
		// 有默认值
		firstname := context.DefaultQuery("firstname","Guest")
		lastname := context.Query("lastname")
		context.String(http.StatusOK,"Hello "+firstname+" "+lastname)
	})

	// 4.form表单
	// 提交参数
	router.POST("/login", func(context *gin.Context) {
		context.DefaultPostForm("type","alert")
		username := context.PostForm("username")
		password := context.PostForm("password")

		// 接受表单内的数组
		hobby := context.PostFormArray("hobby")
		context.String(http.StatusOK,fmt.Sprintf("username:%s\n password:%s\n hobby:%s\n",username,password,strings.Join(hobby,",")))
	})

	// 上传图片
	// 图片最大为8M
	router.MaxMultipartMemory = 8<<20
	router.POST("/upload", func(context *gin.Context) {

		// 单文件上传
		//file, _ :=context.FormFile("img")
		//// file,文件路径
		//context.SaveUploadedFile(file,"C:\\Users\\10852\\Desktop\\aaa\\"+file.Filename)
		//context.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))

		// 多文件上传
		form,err := context.MultipartForm()
		if err!=nil{
			context.String(http.StatusBadRequest,"存储图片失败")
			panic(err)
		}
		files := form.File["imgs"]

		//遍历所有图片
		for _, file := range files {
			if err= context.SaveUploadedFile(file,file.Filename);err!=nil{
				context.String(http.StatusBadRequest,"存储图片失败")
				panic(err)
			}
		}

		context.String(http.StatusOK,"存储图片成功")
	})

	// 5.路由组
	v1 := router.Group("/v1")
	{
		v1.POST("/login", login)
		v1.POST("/submit", submit)
	}

	v2 := router.Group("/v2")
	{
		v2.POST("/login", login)
		v2.POST("/submit", submit)
	}

	router.Run("localhost:8000")
}

func login(context *gin.Context){
	context.String(http.StatusOK,"login")
}
func submit(context *gin.Context){
	context.String(http.StatusOK,"submit")
}
