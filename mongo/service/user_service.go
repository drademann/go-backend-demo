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

func (s *userService) FindAll() ([]model.User, error) {
	return s.repository.FindAll()
}

func (s *userService) FindByName(name string) (model.User, error) {
	return s.repository.FindByName(name)
}
