package env

// import (
// 	"os"

// 	"github.com/joho/godotenv"
// )

// type EnvKey string

// const (
//  AWS_REGION EnvKey = "AWS_REGION"
//  DB_CLUSTER_ARN EnvKey = "DB_CLUSTER_ARN"
//  DB_SECRET_ARN = "DB_SECRET_ARN"
//  DATABASE_NAME EnvKey = "DATABASE_NAME"
// )

// func Load() error {
//  return godotenv.Load(".env")
// }

// func (key EnvKey) GetValue() string {
//  return os.Getenv(string(key))
// }