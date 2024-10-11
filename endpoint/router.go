package endpoint

import "net/http"

func DefaultRouter() {
	http.HandleFunc("/users", getAllUsers)
}
