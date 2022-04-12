package role

import (
	"07rbac/config"
	"07rbac/entity"
	"fmt"
)

func FindRoleByUserId(id int) []entity.Role {
	sqlStr := "select id,userId,permId from role where userId=?"

	fmt.Println(config.DB)
	var r []entity.Role
	err := config.DB.Select(&r, sqlStr, id)
	if err != nil {
		fmt.Printf("get failed, err:%v\n", err)
		return nil
	}
	return r
}
