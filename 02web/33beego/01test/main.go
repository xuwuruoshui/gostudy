package main

import (
	_ "01test/routers"
	"github.com/beego/beego/v2/client/orm"
	beego "github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, _ := beego.AppConfig.String("defaultdb")
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", db)
	beego.Run()
}
