package controllers

import (
	beego "github.com/beego/beego/v2/server/web"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.Data["IsEmail"] = 1

	pages := []struct {
		Num int
	}{{10}, {20}, {30}}
	c.Data["Numbers"] = pages

	c.TplName = "index.html"
}

func (c *MainController) GetHello() {
	c.Ctx.WriteString("Hello World!!!")
}

func (c *MainController) GetHelloWorld() {
	c.Ctx.WriteString("新版路由")
}
