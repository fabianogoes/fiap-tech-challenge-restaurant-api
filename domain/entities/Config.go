package entities

import (
	"fmt"
	"log/slog"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

type Config struct {
	AppName                 string
	Environment             string
	AppPort                 string
	DBUser                  string
	DBPassword              string
	DBHost                  string
	DBPort                  string
	DBName                  string
	APIVersion              string
	TokenSecret             string
	PaymentApiUrl           string
	KitchenApiUrl           string
	AwsRegion               string
	AwsEndpoint             string
	PaymentQueueUrl         string
	PaymentCallbackQueueUrl string
}

func NewConfig() *Config {
	loadEnvironment()

	config := &Config{
		Environment:             strings.TrimRight(os.Getenv("APP_ENV"), "\n\r"),
		AppName:                 strings.TrimRight(os.Getenv("APP_NAME"), "\n\r"),
		AppPort:                 strings.TrimRight(os.Getenv("APP_PORT"), "\n\r"),
		DBHost:                  strings.TrimRight(os.Getenv("DB_HOST"), "\n\r"),
		DBPort:                  strings.TrimRight(os.Getenv("DB_PORT"), "\n\r"),
		DBName:                  strings.TrimRight(os.Getenv("DB_DATABASE"), "\n\r"),
		DBUser:                  strings.TrimRight(os.Getenv("DB_USERNAME"), "\n\r"),
		DBPassword:              strings.TrimRight(os.Getenv("DB_PASSWORD"), "\n\r"),
		APIVersion:              strings.TrimRight(os.Getenv("API_VERSION"), "\n\r"),
		TokenSecret:             strings.TrimRight(os.Getenv("TOKEN_SECRET"), "\n\r"),
		PaymentApiUrl:           strings.TrimRight(os.Getenv("PAYMENT_API_URL"), "\n\r"),
		KitchenApiUrl:           strings.TrimRight(os.Getenv("KITCHEN_API_URL"), "\n\r"),
		AwsRegion:               strings.TrimRight(os.Getenv("AWS_REGION"), "\n\r"),
		AwsEndpoint:             strings.TrimRight(os.Getenv("AWS_ENDPOINT"), "\n\r"),
		PaymentQueueUrl:         strings.TrimRight(os.Getenv("PAYMENT_QUEUE_URL"), "\n\r"),
		PaymentCallbackQueueUrl: strings.TrimRight(os.Getenv("PAYMENT_CALLBACK_QUEUE_URL"), "\n\r"),
	}

	printConfig(config)
	return config
}

func loadEnvironment() {
	switch os.Getenv("APP_ENV") {
	case "production":
		loadProductionEnv()
	case "development":
		loadDevelopmentEnv()
	default:
		loadDefaultEnv()
	}
}

func loadDefaultEnv() {
	_ = os.Setenv("APP_ENV", "default")
	_ = os.Setenv("APP_NAME", "restaurant-api")
	_ = os.Setenv("APP_PORT", ":8080")
	_ = os.Setenv("DB_HOST", "localhost")
	_ = os.Setenv("DB_PORT", "5432")
	_ = os.Setenv("DB_DATABASE", "restaurant_db")
	_ = os.Setenv("DB_USERNAME", "tech_challenge_usr")
	_ = os.Setenv("DB_PASSWORD", "tech_challenge_pwd")
	_ = os.Setenv("API_VERSION", "1.0")
	_ = os.Setenv("TOKEN_SECRET", "123")
	_ = os.Setenv("PAYMENT_API_URL", "http://localhost:8010")
	_ = os.Setenv("KITCHEN_API_URL", "http://localhost:8020")
	_ = os.Setenv("AWS_REGION", "us-east-1")
	_ = os.Setenv("AWS_ENDPOINT", "http://localhost:4566")
	_ = os.Setenv("PAYMENT_QUEUE_URL", "https://localhost.localstack.cloud:4566/000000000000/order-payment-queue")
	_ = os.Setenv("PAYMENT_CALLBACK_QUEUE_URL", "https://localhost.localstack.cloud:4566/000000000000/order-payment-callback-queue")
}

func loadProductionEnv() {
	err := godotenv.Load() // Load .env file
	if err != nil {
		slog.Info("loading .env file not found")
	}
}

func loadDevelopmentEnv() {
	err := godotenv.Load(".env.development")
	if err != nil {
		slog.Error("Error loading .env file", "error", err)
		os.Exit(1)
	}
}

func printConfig(config *Config) {
	fmt.Println("*** Environments ***")
	fmt.Printf("Environment: %s\n", config.Environment)
	fmt.Printf("App Name: %s\n", config.AppName)
	fmt.Printf("App Port: %s\n", config.AppPort)
	fmt.Printf("DB Host: %s\n", config.DBHost)
	fmt.Printf("DB Port: %s\n", config.DBPort)
	fmt.Printf("DB Name: %s\n", config.DBName)
	fmt.Printf("DB User: %s\n", config.DBUser)
	fmt.Printf("DB Password: %s\n", config.DBPassword)
	fmt.Printf("API version: %s\n", config.APIVersion)
	fmt.Printf("Token Secret: %s\n", config.TokenSecret)
	fmt.Printf("Payment API URL: %s\n", config.PaymentApiUrl)
	fmt.Printf("Kitchen API URL: %s\n", config.KitchenApiUrl)
	fmt.Printf("AWS Region: %s\n", config.AwsRegion)
	fmt.Printf("AWS EndPoint: %s\n", config.AwsEndpoint)
	fmt.Printf("Payment Queue URL: %s\n", config.PaymentQueueUrl)
	fmt.Printf("Payment Callback Queue URL: %s\n", config.PaymentCallbackQueueUrl)
}
