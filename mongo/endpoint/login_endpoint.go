package endpoint

import (
	"encoding/json"
	"go-backend-demo/mongo/service"
	"net/http"
)

func login(w http.ResponseWriter, r *http.Request) {
	type LoginBody struct {
		Name     string `json:"name"`
		Password string `json:"password"`
	}
	var body LoginBody
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	jwt, err := service.NewLoginService().Login(body.Name, body.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusForbidden)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	type LoginResponse struct {
		JWT string `json:"jwt"`
	}
	err = json.NewEncoder(w).Encode(LoginResponse{jwt})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
