package service

import (
	"go-backend-demo/repository"
	"go-backend-demo/service/model"
)

func FindAllUsers() ([]model.User, error) {
	return repository.FindAllUsers()
}
