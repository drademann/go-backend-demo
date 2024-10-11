package main

import (
	"fmt"
	"go-backend-demo/endpoint"
	"go-backend-demo/repository"
	"log"
	"net/http"
)

func main() {
	err := repository.OpenDatabase()
	if err != nil {
		log.Fatalf("failed to open database: %v", err)
	}

	endpoint.DefaultRouter()
	fmt.Println("starting server at http://localhost:8000")
	log.Fatalln(http.ListenAndServe(":8000", nil))
}
