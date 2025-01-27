package service

import (
	"go-backend-demo/mongo/repository"
	"go-backend-demo/mongo/service/model"
)

type userService struct {
	repository *repository.UserRepository
}

func newUserService(repository *repository.UserRepository) UserService {
	return &userService{repository: repository}
}

func (s *userService) FindAllUsers() ([]model.User, error) {
	return s.repository.FindAllUsers()
}

func (s *userService) FindUserByName(name string) (model.User, error) {
	return s.repository.FindUserByName(name)
}
