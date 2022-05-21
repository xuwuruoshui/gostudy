package entity

import "time"

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	CreateAt time.Time
	UpdateAt time.Time
	DeleteAt time.Time
}
