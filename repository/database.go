package repository

import (
	"go-backend-demo/repository/entity"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func OpenDatabase() error {
	var err error
	db, err = gorm.Open(sqlite.Open("demo.db"), &gorm.Config{})
	if err != nil {
		return err
	}
	err = db.AutoMigrate(&entity.User{})
	if err != nil {
		return err
	}
	return nil
}
