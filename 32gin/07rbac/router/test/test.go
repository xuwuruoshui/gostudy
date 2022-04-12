package test

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Test(router *gin.Engine) {
	router.GET("/test", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{"test": "ok"})
	})

	router.GET("/test2", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{"test": "ok"})
	})
}
