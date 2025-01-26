package endpoint

import "net/http"

func DefaultRouter() {
	http.HandleFunc("/login", login)
	http.HandleFunc("/users", getAllUsers)
}
