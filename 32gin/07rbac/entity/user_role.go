package entity

import "time"

type UserRole struct {
	Id       int       `json:"id"`
	UserId   int       `json:"userId"`
	RoleId   int       `json:"roleId"`
	CreateAt time.Time `json:"createAt"`
	UpdateAt time.Time `json:"updateAt"`
	DeleteAt time.Time
}
