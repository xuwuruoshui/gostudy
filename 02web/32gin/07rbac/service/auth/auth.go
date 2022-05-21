package service

import (
	result "07rbac/common"
	"07rbac/dao/user"
	"07rbac/entity"
	"07rbac/middleware/claims"
	"github.com/gin-gonic/gin"
)

func Register(u *entity.User) *result.Result {

	// 2、插入
	if err := user.Insert(u); err != nil {
		return result.UNKNOW_ERROR
	}

	return result.SUCCESS
}

func Login(u *entity.User) *result.Result {
	user := user.FindUserbyName(u.Username)
	if user == nil || user.Password != u.Password {
		return result.PASSWORD_ERROR
	}

	token, err := claims.GenToken(u.Username)
	if err != nil {
		return result.UNKNOW_ERROR
	}

	res := result.SUCCESS
	res.Data = gin.H{
		"token": token,
	}
	return res
}
