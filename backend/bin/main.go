package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"backend/internal"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/rdsdata"
)

var RDSClient *rdsdata.Client
var DBClusterArn string
var DBSecretArn string
var DatabaseName string

func main() {
	fmt.Println("Server starting...");

	DBClusterArn = os.Getenv("DB_CLUSTER_ARN")
	DBSecretArn = os.Getenv("DB_SECRET_ARN")
	DatabaseName = os.Getenv("DATABASE_NAME") 

	if DBClusterArn == "" || DBSecretArn == "" {
		log.Fatal("DB_CLUSTER_ARN and DB_SECRET_ARN environment variables are required")
	}

	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion(Region))
	if err != nil {
        log.Fatal(err)
	}

	RDSClient = rdsdata.NewFromConfig(cfg)
	
	http.HandleFunc("/todos", internal.Handle)
	http.HandleFunc("/todos/", internal.HandleByID)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Println("Server running on: " + port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}