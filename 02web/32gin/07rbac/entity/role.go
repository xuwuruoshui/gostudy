package entity

import "time"

type Role struct {
	Id       int       `json:"id"`
	Name     string    `json:"name"`
	CreateAt time.Time `json:"createAt"`
	UpdateAt time.Time `json:"updateAt"`
	DeleteAt time.Time
}
