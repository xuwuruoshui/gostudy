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

		username := c.GetString("username")
		if username == "" {
			c.JSON(http.StatusUnauthorized, result.NO_PERMISSION)
			c.Abort()
			return
		}

		// 查user
		user := user.FindUserbyName(username)

		// TODO 后续role和perm可以保存到redis里面
		// 查role列表
		roles := role.FindRoleByUserId(user.ID)
		roleIds := make([]int, len(roles))
		for i, v := range roles {
			roleIds[i] = v.Id
		}

		// 查perm列表
		permissions := perm.FindPermByRoles(roleIds)

		for _, permission := range permissions {
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
