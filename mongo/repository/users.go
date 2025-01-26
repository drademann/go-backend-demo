package repository

import (
	"github.com/drademann/fugo/fp"
	"go-backend-demo/mongo/repository/database"
	"go-backend-demo/mongo/repository/entity"
	"go-backend-demo/mongo/service/model"
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
