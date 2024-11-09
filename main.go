package main

import (
	"fmt"
	"go-backend-demo/endpoint"
	"go-backend-demo/repository/database"
	"log"
	"net/http"
)

func main() {
	err := database.Open()
	if err != nil {
		log.Fatalf("failed to open database: %v", err)
	}

	endpoint.DefaultRouter()
	fmt.Println("starting server at http://localhost:8000")
	log.Fatalln(http.ListenAndServe(":8000", nil))
}
