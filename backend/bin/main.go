package main

import (
	"backend/internal/api/aws"
	"fmt"
	"log"
	"net/http"
	"os"
)

const (
	PORT string = "PORT"
)

func enableCORS(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next(w, r)
	}
}

func main() {
	fmt.Println("Server starting...")

	aws.InitStorage()

	http.HandleFunc("/todos", enableCORS(aws.Handle))
	http.HandleFunc("/todos/", enableCORS(aws.HandleByID))

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Println("Server running on: " + port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
