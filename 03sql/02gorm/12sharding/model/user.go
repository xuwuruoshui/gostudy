package model

import (
	"gorm.io/gorm"
	"strconv"
	"time"
)

type User struct {
	Id   string `json:"id,omitempty" gorm:"primary_key"`
	Name string `json:"name"`
}

func TableOfIdLastYear(tableName string, idLast int, tm time.Time) func(db *gorm.DB) *gorm.DB {

	tms := tm.Format("20060102")
	return func(db *gorm.DB) *gorm.DB {
		tableName := tableName + "_" + strconv.Itoa(idLast) + "_" + tms
		return db.Table(tableName)
	}
}
