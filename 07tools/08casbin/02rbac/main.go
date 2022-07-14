package main

import (
	"fmt"
	"github.com/casbin/casbin/v2"
)

/**
* @creator: xuwuruoshui
* @date: 2022-07-14 22:12:23
* @content: RBAC角色的访问控制
 */

func main() {

	e, err := casbin.NewEnforcer("./model.conf", "./policy.csv")
	if err != nil {
		fmt.Println(err)
	}

	check(e, "zhangsan", "data", "read")
	check(e, "zhangsan", "data", "write")
	check(e, "lisi", "data", "read")
	check(e, "lisi", "data", "write")
	// rooter是一个角色, 他还有一个角色是admin,所以也有admin下的所有角色
	check(e, "wangwu", "data", "read")
	check(e, "wangwu", "data", "write")

}

func check(e *casbin.Enforcer, sub, obj, act string) {
	ok, _ := e.Enforce(sub, obj, act)
	if ok {
		fmt.Printf("%s CAN %s %s\n", sub, act, obj)
	} else {
		fmt.Printf("%s CANNOT %s %s\n", sub, act, obj)
	}
}
