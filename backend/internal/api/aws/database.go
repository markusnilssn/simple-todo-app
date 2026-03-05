package aws

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
	_ "github.com/lib/pq"
)

type DBSecret struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

var SQLDatabase *sql.DB = nil

func InitStorage() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("failed to load .env %S", err.Error())
		return
	}

	databaseName := os.Getenv("AWS_DATABASE_NAME")
	databaseHost := os.Getenv("AWS_HOST")
	databasePort, err := strconv.Atoi(os.Getenv("AWS_PORT"))
	if err != nil {
		log.Printf("failed to convert aws port %s", err.Error())
		return
	}

	secretName := os.Getenv("AWS_SECRET_NAME")
	region := os.Getenv("AWS_REGION")

	config, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion(region))
	if err != nil {
		log.Fatalf("failed to load .env %s", err.Error())
	}

	svc := secretsmanager.NewFromConfig(config)

	input := &secretsmanager.GetSecretValueInput{
		SecretId:     aws.String(secretName),
		VersionStage: aws.String("AWSCURRENT"), // VersionStage defaults to AWSCURRENT if unspecified
	}

	result, err := svc.GetSecretValue(context.TODO(), input)
	if err != nil {
		log.Fatalf("failed to load .env %s", err.Error())
		return
	}

	var secret DBSecret
	if err := json.Unmarshal([]byte(*result.SecretString), &secret); err != nil {
		log.Fatalf("failed to parse secret JSON: %s", err)
	}

	dsn := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=require",
		databaseHost,
		databasePort,
		secret.Username,
		secret.Password,
		databaseName,
	)
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Println(err)
		return
	}

	err = db.Ping()
	if err != nil {
		log.Println(err)
		return
	}

	SQLDatabase = db

	log.Print("successfully connected to the database!")
}

func InitTable() {
	query := `
CREATE TABLE IF NOT EXISTS todos (
    id SERIAL PRIMARY KEY,
    title TEXT NOT NULL,
    description TEXT,
    priority INTEGER,
    completed BOOLEAN DEFAULT FALSE
)
`
	_, err := SQLDatabase.Exec(query)
	if err != nil {
		log.Fatal(err)
	}

}
