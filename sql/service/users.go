package service

import (
	"go-backend-demo/sql/repository"
	"go-backend-demo/sql/service/model"
)

func FindAllUsers() ([]model.User, error) {
	return repository.FindAllUsers()
}
