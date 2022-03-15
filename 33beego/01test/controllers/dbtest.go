package controllers

import (
	"01test/models"
	"fmt"
	"github.com/beego/beego/v2/server/web/context"
	"strconv"
)

func GetUser(ctx *context.Context) {
	idStr := ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	info, _ := models.UserInfo(id)
	ctx.JSONResp(*info)
}

func AddUser(ctx *context.Context) {
	u := models.User{}
	err := ctx.BindJSON(&u)
	if err != nil {
		fmt.Println(err)
		ctx.JSONResp(map[string]string{"status": "500"})
		return
	}
	models.Save(&u)
	ctx.JSONResp(map[string]string{"status": "200"})
}

func UpdateUser(ctx *context.Context) {
	var u models.User
	err := ctx.BindJSON(&u)
	if err != nil {
		fmt.Println(err)
		ctx.JSONResp(map[string]string{"status": "500"})
		return
	}
	models.Update(&u)
	ctx.JSONResp(map[string]string{"status": "200"})
}

func DelteUser(ctx *context.Context) {
	idStr := ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	err := models.Delete(id)
	if err != nil {
		fmt.Println(err)
		ctx.JSONResp(map[string]string{"status": "500"})
		return
	}
	models.Delete(id)
	ctx.JSONResp(map[string]string{"status": "200"})
}

// 分页查询
func GetUserList(ctx *context.Context) {
	pageStr := ctx.Input.Param(":page")
	sizeStr := ctx.Input.Param(":size")
	page, _ := strconv.Atoi(pageStr)
	size, _ := strconv.Atoi(sizeStr)
	list, err := models.List(page, size)
	if err != nil {
		ctx.JSONResp(map[string]string{"status": "500"})
		return
	}
	ctx.JSONResp(list)
}
