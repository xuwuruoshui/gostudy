package models

import (
	"fmt"
	"github.com/beego/beego/v2/client/orm"
)

type User struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

func init() {
	orm.RegisterModel(new(User))
}

// 获取用户信息
func UserInfo(id int) (*User, error) {
	o := orm.NewOrm()
	u := &User{Id: id}
	err := o.Read(u)
	return u, err
}

func Save(u *User) error {
	o := orm.NewOrm()
	id, err := o.Insert(u)
	fmt.Println(id)
	return err
}

func Update(u *User) error {
	o := orm.NewOrm()
	id, err := o.Update(u)
	fmt.Println(id)
	return err
}

func Delete(id int) error {
	o := orm.NewOrm()
	u := User{Id: id}
	i, err := o.Delete(&u)
	fmt.Println(i)
	return err
}

func List(page, size int) (*[]User, error) {
	o := orm.NewOrm()
	var u []User
	id, err := o.QueryTable("user").
		Filter("id__gt", 2).
		Offset(page).
		Limit(size).
		OrderBy("-id").All(&u)
	fmt.Println(id)
	return &u, err
}

// 原生SQL
// o := orm.NewOrm()
// 查询
// o.Raw("select * from user where id = ?",id).QueryRow(&user)
// 插入
// o.Raw("insert into user(username,password,email) values(?,?,?)",username,password,email).Exec()
// 更新
// o.Raw("update user set name = ? where id = ?",name,id).Exec()
// 删除
// o.Raw("DELETE FROM user where id=?",id).Exec()
// 批量查询
//  o.Raw("select * from user where id > ? order by id desc",id).QueryRow(&users)
