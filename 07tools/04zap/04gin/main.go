package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
)

func main(){

	r := gin.Default()
	port := 9096
	pro, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}
	// 并发安全
	zap.ReplaceGlobals(pro)
	s := zap.S()
	defer s.Sync()
	r.GET("/", func(c *gin.Context) {
		s.Info("调用成功")
		c.JSON(http.StatusOK, map[string]string{"test":"666"})
	})

	err = r.Run(fmt.Sprintf(":%d", port))
	if err != nil {
		s.Panic(err)
	}
}
