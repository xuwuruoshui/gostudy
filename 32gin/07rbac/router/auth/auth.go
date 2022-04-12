package auth

import (
	code "07rbac/common"
	"07rbac/entity"
	service "07rbac/service/auth"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Auth(router *gin.Engine) {

	authGroup := router.Group("/auth")
	{
		// 注册
		authGroup.POST("/register", func(context *gin.Context) {
			var user entity.User
			if err := context.ShouldBindJSON(&user); err != nil {
				context.JSON(http.StatusOK, code.UNKNOW_ERROR)
				return
			}
			context.JSON(http.StatusOK, service.Register(&user))
		})

		// 登录
		authGroup.POST("/login", func(context *gin.Context) {
			var user entity.User
			if err := context.ShouldBindJSON(&user); err != nil {
				context.JSON(http.StatusOK, code.UNKNOW_ERROR)
				return
			}
			context.JSON(http.StatusOK, service.Login(&user))
		})
	}

}
