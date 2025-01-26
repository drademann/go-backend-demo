package database

import (
	"go-backend-demo/sql/repository/entity"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func Open() error {
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

func Find(users *[]entity.User) *gorm.DB {
	return db.Find(users)
}
