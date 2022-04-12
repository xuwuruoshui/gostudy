package entity

import "time"

type Role struct {
	Id       int       `json:"int"`
	UserId   int       `json:"userId" db:"userId"`
	PermId   int       `json:"permId" db:"permId"`
	CreateAt time.Time `json:"createAt"`
	UpdateAt time.Time `json:"updateAt"`
	DeleteAt time.Time
}
