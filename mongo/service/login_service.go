package service

import (
	"fmt"
	"go-backend-demo/mongo/repository"
)

type loginService struct {
	repository *repository.UserRepository
}

func newLoginService(repository *repository.UserRepository) LoginService {
	return &loginService{repository: repository}
}

func (s *loginService) Login(name string, password string) (string, error) {
	u, err := s.repository.FindByName(name)
	if err != nil {
		return "", fmt.Errorf("unknown user - %w", err)
	}
	if password != u.Password {
		return "", fmt.Errorf("wrong password")
	}
	return "jwt", nil
}
