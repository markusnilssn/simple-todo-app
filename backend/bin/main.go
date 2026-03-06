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
	aws.InitTable()

	mutex := http.NewServeMux()
	mutex.HandleFunc("/todos", enableCORS(aws.Handle))
	mutex.HandleFunc("/todos/", enableCORS(aws.HandleByID))

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	address := "0.0.0.0:" + port
	log.Printf("Server starting on %s", address)
	server := &http.Server{
		Addr:    address,
		Handler: mutex,
	}

	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Listen: %s\n", err)
	}
}
