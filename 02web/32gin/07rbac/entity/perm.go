package entity

import "time"

type Perm struct {
	Id       int       `json:"id"`
	Path     string    `json:"path"`
	CreateAt time.Time `json:"createAt"`
	UpdateAt time.Time `json:"updateAt"`
	DeleteAt time.Time
}
