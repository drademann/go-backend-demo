package database

import (
	"go-backend-demo/repository/entity"
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

func Transactional(f func(tx *gorm.DB) error) error {
	return db.Transaction(f)
}

func Find(users *[]entity.User) *gorm.DB {
	return db.Find(users)
}
