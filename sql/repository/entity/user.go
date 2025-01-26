package entity

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name string
}

func (User) TableName() string {
	return "app_users"
}
