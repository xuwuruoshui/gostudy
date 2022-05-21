package perm

import (
	"07rbac/config"
	"07rbac/entity"
	"fmt"
	"github.com/jmoiron/sqlx"
)

func FindPermByRoles(roleIds []int) []entity.Perm {
	sqlStr := `
	SELECT
	p.id,
	p.path
FROM
	role r
	LEFT JOIN perm_role pr ON  pr.role_id = r.id
	LEFT JOIN perm p ON pr.perm_id = p.id
	where r.id in (?)
	GROUP BY p.id;`

	query, args, _ := sqlx.In(sqlStr, roleIds)
	var r []entity.Perm
	query = config.DB.Rebind(query)
	err := config.DB.Select(&r, query, args...)
	if err != nil {
		fmt.Printf("get failed, err:%v\n", err)
		return nil
	}
	return r
}
