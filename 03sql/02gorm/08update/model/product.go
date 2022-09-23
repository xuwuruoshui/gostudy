package model

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Code string `json:"code" gorm:""`
	Price int64 `json:"price"`
}
