package middleware

import (
	auth "07rbac/middleware/auth"
	"07rbac/middleware/perm"
	"github.com/gin-gonic/gin"
)

func Middleware(engine *gin.Engine) {
	// 鉴权
	engine.Use(auth.Auth())
	engine.Use(perm.Perm())
}
