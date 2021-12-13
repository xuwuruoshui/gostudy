package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

func main() {
	r := gin.Default()

	// 1.自定义全局中间件
	r.Use(Logger())

	{
		r.GET("/test", func(c *gin.Context) {
			// 如果值不存在就panic
			example := c.MustGet("example").(string)
			log.Println("请求中:",example)
		})

		// 2.局部中间件,第二参数
		// 全局与局部中间件顺序: 全局前 局部前 请求中 局部后 全局后
		r.GET("/test1",Logger1(), func(c *gin.Context) {
			example := c.MustGet("test")
			log.Println("请求中:",example)
		})
	}

	// 5.同步与异步
	//syncandasync(r)



	r.Run(":8080")
}

func Logger() gin.HandlerFunc{
	return func(c *gin.Context) {

		c.Set("example","123456")

		// 请求前
		log.Println("请求前")

		// 遍历所有的handle
		c.Next()

		// 请求后
		log.Println("请求后")

		// 返回http当前响应的状态码
		status := c.Writer.Status()
		log.Println("请求后",status)
	}
}

func Logger1() gin.HandlerFunc{
	return func(c *gin.Context) {
		c.Set("test","123456")

		log.Println("请求前(局部)")
		c.Next()
		log.Println("请求后(局部)")

		// 返回http当前响应的状态码
		status := c.Writer.Status()
		log.Println("请求后(局部)",status)
	}
}

func syncandasync(r *gin.Engine) {
	// 1.同步与异步
	// 执行的效果，同步浏览器要转5秒的圈，异步则是浏览器访问直接完成。打印都是5秒后
	r.GET("/long_sync", func(c *gin.Context) {
		time.Sleep(5 * time.Second)
		log.Println("Done! in path" + c.Request.URL.Path)
	})
	r.GET("/long_async", func(c *gin.Context) {
		// 不能直接用原始的context,要用copy过的
		cCp := c.Copy()
		go func() {
			time.Sleep(5 * time.Second)
			log.Println("Done! in path" + cCp.Request.URL.Path)
		}()
	})
}

