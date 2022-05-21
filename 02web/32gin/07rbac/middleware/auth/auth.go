package auth

import (
	result "07rbac/common"
	"07rbac/middleware/claims"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// auth开头跳过
		if strings.HasPrefix(c.Request.RequestURI, "/auth") {
			c.Next()
		} else {
			// 授权
			tokenValidate(c)
		}
	}
}

func tokenValidate(c *gin.Context) {
	authHeader := c.Request.Header.Get("Authorization")
	if authHeader == "" {
		c.JSON(http.StatusUnauthorized, result.NOT_FOUND_TOKEN)
		c.Abort()
		return
	}
	// 按空格分割
	parts := strings.SplitN(authHeader, " ", 2)
	if !(len(parts) == 2 && parts[0] == "Bearer") {
		c.JSON(http.StatusUnauthorized, result.TOKEN_FORMATE_ERROR)
		c.Abort()
		return
	}
	// parts[1]是获取到的tokenString，我们使用之前定义好的解析JWT的函数来解析它
	mc, err := claims.ParseToken(parts[1])
	if err != nil {
		c.JSON(http.StatusUnauthorized, result.INVALID_TOKEN)
		c.Abort()
		return
	}
	// 将当前请求的username信息保存到请求的上下文c上
	c.Set("username", mc.Username)
	c.Next() // 后续的处理函数可以用过c.Get("username")来获取当前请求的用户信
}
