package main

import (
	"fmt"
)

type Manager struct {
	students map[uint]*Student
}

func (manager Manager) addStudent() {
	fmt.Println("请输入要添加的学生的学号:")
	var id uint
	fmt.Scan(&id)
	_, ok := manager.students[id]
	if ok {
		fmt.Println("已存在学生!!!")
		return
	}

	fmt.Println("请输入要添加的学生的姓名:")
	var name string
	fmt.Scan(&name)

	fmt.Println("是否添加(y,n)?")
	var isAdd string
	fmt.Scan(&isAdd)
	if isAdd == "y" {
		manager.students[id] = &Student{id, name}
		fmt.Println("添加成功!!!")
	} else {
		fmt.Println("已取消添加!!!")
	}
}

func (manager Manager) editStudent() {
	fmt.Println("请输入要修改的学生的学号:")
	var id uint
	fmt.Scan(&id)
	student, ok := manager.students[id]
	if !ok {
		fmt.Println("不存在该学生!!!")
		return
	}

	fmt.Println("请输入要修改的学生的姓名:")
	var name string
	fmt.Scan(&name)

	fmt.Println("是否修改(y,n)?")
	var isEdit string
	fmt.Scan(&isEdit)
	if isEdit == "y" {
		student.name = name
		fmt.Println("修改成功!!!")
	} else {
		fmt.Println("已取消修改!!!")
	}
}

func (manager Manager) delStudent() {
	fmt.Println("请输入要删除学生的学号:")
	var id uint
	fmt.Scan(&id)

	var isDel string
	fmt.Println("是否删除(y,n)?")
	fmt.Scan(&isDel)
	if isDel == "y" {
		delete(manager.students, id)
		fmt.Println("删除成功!!!")
	} else {
		fmt.Println("已取消修改!!!")
	}
}

func (manager Manager) searchStudent() {
	fmt.Println("请输入要查询的学生的学号:")
	var id uint
	fmt.Scan(&id)
	student, ok := manager.students[id]
	if ok {
		fmt.Println(student)
		fmt.Println("查询成功!!!")
	} else {
		fmt.Println("查无此人!!!")
	}

}
