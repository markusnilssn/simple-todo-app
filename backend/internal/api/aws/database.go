package aws

import (
	"context"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/rdsdata"
	"github.com/joho/godotenv"
)

const (
	AWS_REGION            string = "AWS_REGION"
	DATABASE_NAME         string = "DATABASE_NAME"
	AWS_ACCESS_KEY_ID     string = "AWS_ACCESS_KEY_ID"
	AWS_SECRET_ACCESS_KEY string = "AWS_SECRET_ACCESS_KEY"
)

var RDSClient *rdsdata.Client
var AWSAccessKeyID string
var AWSSecretAccessKey string

func InitStorage() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("failed to load .env %S", err.Error())
	}

	AWSAccessKeyID = os.Getenv(AWS_ACCESS_KEY_ID)
	AWSSecretAccessKey = os.Getenv(AWS_SECRET_ACCESS_KEY)

	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion(os.Getenv(AWS_REGION)))
	if err != nil {
		log.Fatal(err)
	}

	RDSClient = rdsdata.NewFromConfig(cfg)
	if RDSClient == nil {
		log.Println("rsd client is null")
	}

}
