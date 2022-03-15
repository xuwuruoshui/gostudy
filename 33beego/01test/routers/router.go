package routers

import (
	"01test/controllers"
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/hello", &controllers.MainController{}, "get:GetHello")
	beego.Get("/demo", controllers.GetHello)
	ns := beego.NewNamespace("/v1",
		beego.NSGet("/a", controllers.GetA),
		beego.NSPost("/b", controllers.POSTA))
	beego.AddNamespace(ns)
	beego.InsertFilter("/v1/*", beego.BeforeRouter, func(ctx *context.Context) {
		ctx.WriteString("禁止访问")
	})
	// 新语法测试
	beego.CtrlGet("/test", (*controllers.MainController).GetHelloWorld)

	// db测试,增删改查
	beego.Get("/user/:id", controllers.GetUser)
	beego.Post("/user", controllers.AddUser)
	beego.Put("/user", controllers.UpdateUser)
	beego.Delete("/user/:id", controllers.DelteUser)
	beego.Get("/user/list/:page/:size", controllers.GetUserList)
}
