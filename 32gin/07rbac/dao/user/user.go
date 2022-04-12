package user

import (
	result "07rbac/common"
	"07rbac/config"
	"07rbac/entity"
	"fmt"
	"time"
)

// 是否存在
func FindUserbyName(username string) *entity.User {
	sqlStr := "select id,username,password from user where username=?"

	fmt.Println(config.DB)
	var u entity.User
	err := config.DB.Get(&u, sqlStr, username)
	if err != nil {
		fmt.Printf("get failed, err:%v\n", err)
		return nil
	}

	return &u
}

// 插入
func Insert(user *entity.User) *result.Result {

	tx, err := config.DB.Beginx() // 开启事务
	if err != nil {
		fmt.Printf("begin trans failed, err:%v\n", err)
		return result.UNKNOW_ERROR
	}
	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p) // re-throw panic after Rollback
		} else if err != nil {
			fmt.Println("rollback")
			tx.Rollback() // err is non-nil; don't change it
		} else {
			err = tx.Commit() // err is nil; if Commit returns error update err
			fmt.Println("commit")
		}
	}()

	u := FindUserbyName(user.Username)
	if u != nil {
		return result.USER_IS_EXISTED
	}

	sqlStr := "insert into user(username, password,createAt,updateAt) values (?,?,?,?)"
	ret, err := config.DB.Exec(sqlStr, user.Username, user.Password, time.Now(), time.Now())
	if err != nil {
		fmt.Printf("insert failed, err:%v\n", err)
		return result.UNKNOW_ERROR
	}
	affected, err := ret.RowsAffected()
	if affected > 0 {
		return nil
	}
	return result.UNKNOW_ERROR
}
