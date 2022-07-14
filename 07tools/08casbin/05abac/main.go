package main

import (
	"fmt"
	"github.com/casbin/casbin/v2"
	"log"
)

/**
* @creator: xuwuruoshui
* @date: 2022-07-14 22:12:23
* @content: ABAC属性的权限控制
 */

type Object struct {
	Name  string
	Owner string
}

type Subject struct {
	Name string
	Hour int
}

func check(e *casbin.Enforcer, sub Subject, obj Object, act string) {
	ok, _ := e.Enforce(sub, obj, act)
	if ok {
		fmt.Printf("%s CAN %s %s at %d:00\n", sub.Name, act, obj.Name, sub.Hour)
	} else {
		fmt.Printf("%s CANNOT %s %s at %d:00\n", sub.Name, act, obj.Name, sub.Hour)
	}
}

func main() {
	e, err := casbin.NewEnforcer("./model.conf", "./policy.csv")
	if err != nil {
		log.Fatalf("NewEnforecer failed:%v\n", err)
	}

	// 数据,及其所有者
	// model.conf中设置了非所有者只能9点到18点之间能访问,数据所有者任何时间都能访问
	o := Object{"data", "zhangsan"}

	// lisi访问的时间
	s1 := Subject{"lisi", 10}
	check(e, s1, o, "read")

	s2 := Subject{"zhangsan", 10}
	check(e, s2, o, "read")

	s4 := Subject{"lisi", 20}
	check(e, s4, o, "read")

	s3 := Subject{"zhangsan", 20}
	check(e, s3, o, "read")
}
