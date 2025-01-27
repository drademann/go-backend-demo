package service

import (
	"go-backend-demo/mongo/repository"
	"go-backend-demo/mongo/service/model"
)

type LoginService interface {
	Login(name string, password string) (string, error)
}

func NewLoginService() LoginService {
	return newLoginService(repository.NewUserRepository())
}

type UserService interface {
	FindAll() ([]model.User, error)
	FindByName(name string) (model.User, error)
}

func NewUserService() UserService {
	return newUserService(repository.NewUserRepository())
}
