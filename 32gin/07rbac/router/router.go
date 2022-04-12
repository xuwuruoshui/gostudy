package router

import (
	"07rbac/router/auth"
	"07rbac/router/test"
	"github.com/gin-gonic/gin"
)

func RouterEntry(router *gin.Engine) {

	// 授权
	auth.Auth(router)
	test.Test(router)
}
