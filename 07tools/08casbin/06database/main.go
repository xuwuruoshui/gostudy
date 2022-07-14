package main

import (
	"fmt"
	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	_ "github.com/go-sql-driver/mysql"
)

/**
* @creator: xuwuruoshui
* @date: 2022-07-14 22:12:23
* @content: RBAC保存到数据库
 */

func check(e *casbin.Enforcer, sub, obj, act string) {
	ok, _ := e.Enforce(sub, obj, act)
	if ok {
		fmt.Printf("%s CAN %s %s\n", sub, act, obj)
	} else {
		fmt.Printf("%s CANNOT %s %s\n", sub, act, obj)
	}
}

func main() {
	a, _ := gormadapter.NewAdapter("mysql", "root:123456@tcp(127.0.0.1:3306)/")
	e, _ := casbin.NewEnforcer("./model.conf", a)

	e.LoadPolicy()

	// 数据库中设置了zhangsan为admin,lisi为student
	check(e, "zhangsan", "/user", "POST")
	check(e, "zhangsan", "/user", "GET")
	check(e, "lisi", "/user", "POST")
	check(e, "lisi", "/user", "GET")

	// 1.为角色添加访问权限
	e.AddPolicy("student", "/user", "POST")
	check(e, "lisi", "/user", "POST")

	// 2.修改权限
	e.UpdatePolicy([]string{"student", "/user", "POST"}, []string{"student", "/user", "PUT"})
	check(e, "lisi", "/user", "POST")
	check(e, "lisi", "/user", "PUT")

	// 3.删除权限
	e.RemovePolicy("student", "/user", "PUT")
	check(e, "lisi", "/user", "PUT")
}
