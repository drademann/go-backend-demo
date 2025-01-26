package repository

import (
	"github.com/drademann/fugo/fp"
	"go-backend-demo/repository/database"
	"go-backend-demo/repository/entity"
	"go-backend-demo/service/model"
)

func FindAllUsers() ([]model.User, error) {
	collection := database.Collection("appUser")
	var users []entity.User
	err := database.FindAll(collection, &users)
	if err != nil {
		return nil, err
	}
	return fp.Map(users, toModelUser), nil
}

func toModelUser(entity entity.User) model.User {
	return model.User{
		ID:   model.UserID(entity.ID.Hex()),
		Name: entity.Name,
	}
}
