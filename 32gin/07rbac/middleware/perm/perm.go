package perm

import (
	result "07rbac/common"
	"07rbac/dao/perm"
	"07rbac/dao/role"
	"07rbac/dao/user"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func Perm() gin.HandlerFunc {
	return func(c *gin.Context) {
		if strings.HasPrefix(c.Request.RequestURI, "/auth") {
			c.Next()
			return
		}

		u, exists := c.Get("username")
		if !exists {
			c.JSON(http.StatusUnauthorized, result.NO_PERMISSION)
			c.Abort()
			return
		}
		username := u.(string)
		// 查user
		user := user.FindUserbyName(username)
		// 查role
		role := role.FindRoleByUserId(user.ID)
		// 查perm
		for _, v := range role {
			permission := perm.FindPermById(v.PermId)
			if permission.Path == c.Request.RequestURI {
				c.Next()
				return
			}
		}
		c.JSON(http.StatusUnauthorized, result.NO_PERMISSION)
		c.Abort()
		return
	}
}
