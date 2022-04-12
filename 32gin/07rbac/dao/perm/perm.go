package perm

import (
	"07rbac/config"
	"07rbac/entity"
	"fmt"
)

func FindPermById(id int) *entity.Perm {
	sqlStr := "select path from perm where Id=?"

	fmt.Println(config.DB)
	var r entity.Perm
	err := config.DB.Get(&r, sqlStr, id)
	if err != nil {
		fmt.Printf("get failed, err:%v\n", err)
		return nil
	}
	return &r
}
