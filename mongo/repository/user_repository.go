package repository

import (
	"github.com/drademann/fugo/fp"
	"go-backend-demo/mongo/repository/database"
	"go-backend-demo/mongo/repository/entity"
	"go-backend-demo/mongo/service/model"
	"go.mongodb.org/mongo-driver/bson"
)

type UserRepository struct{}

func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

func (r *UserRepository) FindAllUsers() ([]model.User, error) {
	collection := database.Collection("appUser")
	var users []entity.User
	err := database.FindAll(collection, &users)
	if err != nil {
		return nil, err
	}
	return fp.Map(users, toUserModel), nil
}

func (r *UserRepository) FindUserByName(name string) (model.User, error) {
	collection := database.Collection("appUser")
	var user entity.User
	err := database.FindOne(collection, bson.M{"name": name}, &user)
	if err != nil {
		return model.User{}, err
	}
	return toUserModel(user), nil
}

func toUserModel(entity entity.User) model.User {
	return model.User{
		ID:       model.UserID(entity.ID.Hex()),
		Name:     entity.Name,
		Password: entity.Password,
	}
}
