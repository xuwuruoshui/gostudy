package main

import (
	"fmt"
	"github.com/casbin/casbin/v2"
)

/**
* @creator: xuwuruoshui
* @date: 2022-07-14 22:12:23
* @content: ACL用户的权限管理模型
 */

func main() {

	e, err := casbin.NewEnforcer("./model.conf", "./policy.csv")
	if err != nil {
		fmt.Println(err)
	}

	// 这两个有权限
	check(e, "zhangsan", "data1", "read")
	check(e, "lisi", "data2", "write")

	// 这三个没有
	check(e, "wangwu", "data1", "write")
	check(e, "zhaoliu", "data2", "read")
	check(e, "zhangsan", "data2", "read")
}

func check(e *casbin.Enforcer, sub, obj, act string) {
	ok, _ := e.Enforce(sub, obj, act)
	if ok {
		fmt.Printf("%s CAN %s %s\n", sub, act, obj)
	} else {
		fmt.Printf("%s CANNOT %s %s\n", sub, act, obj)
	}
}
