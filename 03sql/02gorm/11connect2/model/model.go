package model

import "gorm.io/gorm"

type Employer struct {
	gorm.Model
	Name string
	CompanyId int
	Company Company
	Creditcards []*Creditcard
}


type Company struct {
	gorm.Model
	Name string
}

// 信用卡
type Creditcard struct {
	gorm.Model
	Number string
	EmployerId uint
}