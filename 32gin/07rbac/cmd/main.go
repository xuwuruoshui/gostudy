package main

import (
	"07rbac/middleware"
	"07rbac/router"
	"github.com/gin-gonic/gin"
)

func init() {

}

func main() {
	engine := gin.Default()

	middleware.Middleware(engine)

	// 路由配置
	router.RouterEntry(engine)

	engine.Run("127.0.0.1:3000")
}
