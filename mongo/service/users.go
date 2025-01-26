package service

import (
	"go-backend-demo/mongo/repository"
	"go-backend-demo/mongo/service/model"
)

func FindAllUsers() ([]model.User, error) {
	return repository.FindAllUsers()
}

func FindUserByName(name string) (model.User, error) {
	return repository.FindUserByName(name)
}
