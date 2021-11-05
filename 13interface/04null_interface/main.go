package main

import "fmt"

/**
* @creator: xuwuruoshui
* @date: 2021-10-27 00:02:10
* @content: 空接口
 */

func main() {
	i1 := make(map[string]interface{})

	// 1.空接口可以存储任意类型的值
	i1["name"] = "zhangsan"
	i1["age"] = 12
	i1["hobbies"] = [...]string{"a", "b", "c"}
	fmt.Println(i1)


	// 2. 类型断言1 判断空接口的值
	method1("aaa")
	method2(true)

}

func method1(x interface{}){
	v,ok := x.(string)
	if ok{
		fmt.Printf("%s时%T类型\n",v,v)
	}else{
		fmt.Println("断言类型失败")
	}
}


func method2(x interface{}){
	switch v:=x.(type){
	case string:
		fmt.Printf("%T\n",v)
	case int:
		fmt.Printf("%T\n",v)
	case bool:
		fmt.Printf("%T\n",v)
	default:
		fmt.Println("unknown type")
	}
}