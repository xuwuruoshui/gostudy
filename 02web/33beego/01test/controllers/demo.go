package controllers

import (
	"github.com/beego/beego/v2/server/web/context"
)

// 函数式风格
func GetHello(ctx *context.Context) {
	ctx.WriteString("HelloWorld1")
}

func GetA(ctx *context.Context) {
	ctx.WriteString("GET-A")
}

func POSTA(ctx *context.Context) {
	ctx.WriteString("POST-B")
}
