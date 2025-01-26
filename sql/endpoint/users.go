package endpoint

import (
	"encoding/json"
	"github.com/drademann/fugo/fp"
	"go-backend-demo/sql/endpoint/dto"
	"go-backend-demo/sql/service"
	"go-backend-demo/sql/service/model"
	"net/http"
)

func getAllUsers(w http.ResponseWriter, _ *http.Request) {
	users, err := service.FindAllUsers()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(fp.Map(users, func(user model.User) dto.User {
		return dto.User{
			ID:   uint(user.ID),
			Name: user.Name,
		}
	}))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
