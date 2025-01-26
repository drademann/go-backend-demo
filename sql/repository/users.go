package repository

import (
	"github.com/drademann/fugo/fp"
	"go-backend-demo/sql/repository/database"
	"go-backend-demo/sql/repository/entity"
	"go-backend-demo/sql/service/model"
)

func FindAllUsers() ([]model.User, error) {
	var users []entity.User
	rs := database.Find(&users)
	if rs.Error != nil {
		return nil, rs.Error
	}
	return fp.Map(users, toUserModel), rs.Error
}

func toUserModel(entity entity.User) model.User {
	return model.User{
		ID:   model.UserID(entity.ID),
		Name: entity.Name,
	}
}
