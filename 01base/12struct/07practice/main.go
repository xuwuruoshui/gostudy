package main

import (
	"fmt"
	"os"
)

/**
* @creator: xuwuruoshui
* @date:2021-10-16 17:34:49
* @content: 学生管理系统
 */

var arr []*Student = make([]*Student, 0)

func init() {
	for i := 1; i <= 10; i++ {
		arr = append(arr, newStudent(i, fmt.Sprintf("张三%d", i)))
	}
}

type Student struct {
	id   int
	name string
	// 0删除	1正常
	status int8
}

func add() {

	var name string
	for {
		fmt.Println("请输入要添加学生的姓名,退回主菜单请按0:")
		fmt.Scan(&name)
		if name == "0" {
			return
		}
		arr = append(arr, newStudent(len(arr)+1, name))
	}
}

func del() {

	var id int
	for {
		fmt.Println("请输入要删除学生的id,退回主菜单请按0:")
		fmt.Scan(&id)
		if id == 0 {
			return
		}

		if arr[id-1] != nil {
			arr[id-1].status = 0
		}
	}
}

func edit() {

	var id int
	var name string

	for {
		fmt.Println("请输入要添加学生的id和姓名,退回主菜单请按0 0:")
		fmt.Scan(&id, &name)
		if id == 0 && name == "0" {
			return
		}
		if arr[id-1] != nil {
			arr[id-1].name = name
		}
	}
}

func search() {

	var id int
	for {
		fmt.Println("请输入要搜索学生的id,退回主菜单请按0:")
		fmt.Scan(&id)

		if id == 0 {
			return
		}

		if id > 0 && arr[id-1] != nil {
			if arr[id-1].status == 0 {
				fmt.Println("该学生已经被删除！！！")
			} else {
				fmt.Println(arr[id-1])
			}
		} else {
			fmt.Println("不存在该学生！！！")
		}
	}

}

func searchAll() {
	for _, v := range arr {
		if v.status != 0 {
			fmt.Println(v)
		}
	}
}

func newStudent(id int, name string) *Student {
	return &Student{id: id, name: name, status: 1}
}

func main() {

	var choose int
	for {
		fmt.Println("=================")
		fmt.Println("===学生管理系统===")
		fmt.Println("=================")
		fmt.Println("1.添加学生")
		fmt.Println("2.删除学生")
		fmt.Println("3.修改学生")
		fmt.Println("4.查询学生")
		fmt.Println("5.查询所有学生")
		fmt.Println("0.退出")
		fmt.Println("请输入选项:")
		fmt.Scan(&choose)
		switch choose {
		case 1:
			add()
		case 2:
			del()
		case 3:
			edit()
		case 4:
			search()
		case 5:
			searchAll()
		case 0:
			os.Exit(0)
		}
	}
}
