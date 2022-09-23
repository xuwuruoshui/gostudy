package model

type Food struct {
	FoodId int `gorm:"primarykey"`
	Name string `gorm:"column:food_name;type:varchar(32);index:idx_food_name,unique"`
}
