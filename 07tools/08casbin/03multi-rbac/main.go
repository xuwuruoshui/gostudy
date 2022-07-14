package main

import (
	"fmt"
	"github.com/casbin/casbin/v2"
)

/**
* @creator: xuwuruoshui
* @date: 2022-07-14 22:12:23
* @content: 多RBAC,用户和资源都有角色
 */

func main() {

	e, err := casbin.NewEnforcer("./model.conf", "./policy.csv")
	if err != nil {
		fmt.Println(err)
	}

	// 这两个有权限
	check(e, "zhangsan", "prod.data", "read")
	check(e, "zhangsan", "prod.data", "write")
	check(e, "lisi", "dev.data", "read")
	check(e, "lisi", "dev.data", "write")
	check(e, "lisi", "prod.data", "write")

}

func check(e *casbin.Enforcer, sub, obj, act string) {
	ok, _ := e.Enforce(sub, obj, act)
	if ok {
		fmt.Printf("%s CAN %s %s\n", sub, act, obj)
	} else {
		fmt.Printf("%s CANNOT %s %s\n", sub, act, obj)
	}
}
