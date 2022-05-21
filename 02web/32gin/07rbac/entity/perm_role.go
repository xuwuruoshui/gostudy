package entity

import "time"

type PermRole struct {
	Id       int       `json:"id"`
	PermId   int       `json:"permId"`
	RoleId   int       `json:"roleId"`
	CreateAt time.Time `json:"createAt"`
	UpdateAt time.Time `json:"updateAt"`
	DeleteAt time.Time
}
