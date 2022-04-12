package role

import (
	"07rbac/config"
	"07rbac/entity"
	"fmt"
)

func FindRoleByUserId(id int) []entity.Role {
	sqlStr := `
	SELECT
	r.id,
	r.name
FROM
	user u
	LEFT JOIN user_role ur ON ur.user_id = u.id
	LEFT JOIN role r ON ur.role_id = r.id
	where u.id=?
	GROUP BY r.id;
`

	var r []entity.Role
	err := config.DB.Select(&r, sqlStr, id)
	if err != nil {
		fmt.Printf("get failed, err:%v\n", err)
		return nil
	}
	return r
}
