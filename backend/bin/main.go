package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"backend/env"
	"backend/internal"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/rdsdata"
)

var RDSClient *rdsdata.Client
var DBClusterArn string
var DBSecretArn string

func main() {
	fmt.Println("Server starting...");

	DBClusterArn = os.Getenv("DB_CLUSTER_ARN")
	DBSecretArn = os.Getenv("DB_SECRET_ARN")

	// err := env.Load();
	// if(err != nil) {
	// 	log.Fatal("failed to load .env %S", err.Error())
	// }

	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion(env.AWS_REGION.GetValue()))
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