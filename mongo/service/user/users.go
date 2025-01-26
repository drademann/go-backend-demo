package user

import (
	"go-backend-demo/mongo/repository/user"
	"go-backend-demo/mongo/service/model"
)

func FindAll() ([]model.User, error) {
	return user.FindAll()
}

func FindByName(name string) (model.User, error) {
	return user.FindByName(name)
}
