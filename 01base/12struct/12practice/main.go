package main

import (
	"fmt"
	"os"
)

/**
* @creator: xuwuruoshui
* @date: 2021-10-16 23:18:17
* @content: 学生管理系统
* @tip: 运行时需要一起运行,执行命令 `go run some_test.go manager.go student.go`
 */

func main() {

	m1 := Manager{make(map[uint]*Student, 0)}

	var option uint8
	for {
		fmt.Println(`
=====欢迎进入学生管理系统=====
1.添加学生
2.删除学生
3.修改学生
4.查询学生
0.退出
请输入要操作的选项:
		`)
		fmt.Scan(&option)
		switch option {
		case 1:
			m1.addStudent()
		case 2:
			m1.delStudent()
		case 3:
			m1.editStudent()
		case 4:
			m1.searchStudent()
		case 0:
			os.Exit(0)
		default:
			fmt.Println("输入错误")
		}
	}
}
