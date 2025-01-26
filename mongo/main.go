package main

import (
	"go-backend-demo/mongo/endpoint"
	"go-backend-demo/mongo/repository/database"
	"log"
	"net/http"
)

func main() {
	err := database.Open("demo")
	if err != nil {
		log.Fatalf("failed to open database: %v", err)
	}
	defer database.Close()

	endpoint.DefaultRouter()
	log.Println("starting server at http://localhost:8000")
	log.Fatalln(http.ListenAndServe(":8000", nil))
}
