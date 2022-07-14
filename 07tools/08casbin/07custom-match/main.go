package main

import (
	"fmt"
	"github.com/casbin/casbin/v2"
	"regexp"
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

	// 自带规则 https://casbin.org/docs/en/function
	// gin用keyMatch2

	// 客制匹配规则
	//e.AddFunction("my_func", KeyMatchFunc)

	check(e, "dajun", "/user/1", "post")
	check(e, "dajun", "/user/2", "get")
	check(e, "dajun", "/user/test/1", "get")
	check(e, "dajun", "/user/test/1", "pot")

}

func KeyMatchFunc(args ...interface{}) (interface{}, error) {
	name1 := args[0].(string)
	name2 := args[1].(string)

	return KeyMatch(name1, name2), nil
}

func KeyMatch(key1, key2 string) bool {
	res, _ := regexp.MatchString(key2, key1)
	return res
}

func check(e *casbin.Enforcer, sub, obj, act string) {
	ok, _ := e.Enforce(sub, obj, act)
	if ok {
		fmt.Printf("%s CAN %s %s\n", sub, act, obj)
	} else {
		fmt.Printf("%s CANNOT %s %s\n", sub, act, obj)
	}
}
