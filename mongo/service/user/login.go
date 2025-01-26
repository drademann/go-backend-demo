package user

import (
	"fmt"
	"go-backend-demo/mongo/repository/user"
)

func Login(name string, password string) (string, error) {
	u, err := user.FindByName(name)
	if err != nil {
		return "", fmt.Errorf("unknown user - %w", err)
	}
	if password != u.Password {
		return "", fmt.Errorf("wrong password")
	}
	return "jwt", nil
}
