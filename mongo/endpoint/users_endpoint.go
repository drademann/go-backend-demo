package endpoint

import (
	"encoding/json"
	"github.com/drademann/fugo/fp"
	"go-backend-demo/mongo/endpoint/dto"
	"go-backend-demo/mongo/service"
	"go-backend-demo/mongo/service/model"
	"net/http"
)

func getAllUsers(w http.ResponseWriter, _ *http.Request) {
	users, err := service.NewUserService().FindAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(fp.Map(users, func(user model.User) dto.User {
		return dto.User{
			ID:   string(user.ID),
			Name: user.Name,
		}
	}))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
