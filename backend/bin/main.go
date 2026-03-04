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

	"github.com/joho/godotenv"
)

// global variables

var RDSClient *rdsdata.Client
var AWSAccessKeyID string
var AWSSecretAccessKey string

const (
	PORT string = "PORT"
	AWS_REGION string = "AWS_REGION"
	DATABASE_NAME string = "DATABASE_NAME"
	AWS_ACCESS_KEY_ID string = "AWS_ACCESS_KEY_ID"
	AWS_SECRET_ACCESS_KEY string = "AWS_SECRET_ACCESS_KEY"
)

func main() {
	fmt.Println("Server starting...");

	err := godotenv.Load(".env")
	if(err != nil) {
		log.Fatal("failed to load .env %S", err.Error())
	}
	
	AWSAccessKeyID = os.Getenv(AWS_ACCESS_KEY_ID)
	AWSSecretAccessKey = os.Getenv(AWS_SECRET_ACCESS_KEY)

	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion(os.Getenv(AWS_REGION)))
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