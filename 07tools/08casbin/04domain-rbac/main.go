package main

import (
	"fmt"
	"github.com/casbin/casbin/v2"
)

/**
* @creator: xuwuruoshui
* @date: 2022-07-14 22:12:23
* @content: 角色下多领域
 */

func main() {

	e, err := casbin.NewEnforcer("./model.conf", "./policy.csv")
	if err != nil {
		fmt.Println(err)
	}

	// 张三是tenant1领域的admin,但在tenant2领域只是个developer
	check(e, "zhangsan", "tenant1", "data1", "read")
	check(e, "zhangsan", "tenant1", "data1", "write")
	check(e, "zhangsan", "tenant2", "data2", "read")
	check(e, "zhangsan", "tenant2", "data2", "write")

}

func check(e *casbin.Enforcer, sub, domain, obj, act string) {
	ok, _ := e.Enforce(sub, domain, obj, act)
	if ok {
		fmt.Printf("%s CAN %s %s in %s\n", sub, act, obj, domain)
	} else {
		fmt.Printf("%s CANNOT %s %s in %s\n", sub, act, obj, domain)
	}
}
