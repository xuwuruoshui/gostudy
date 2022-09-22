package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

// 优雅退出gin
func main(){
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK,map[string]string{
			"msg":"ok",
		})
	})

	go func() {
		r.Run(":8888")
	}()

	exit := make(chan os.Signal)
	signal.Notify(exit,syscall.SIGINT,syscall.SIGTERM)
	<-exit
	fmt.Println("程序结束")
}
