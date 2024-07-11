package entities

import (
	"fmt"
	"log/slog"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

type Config struct {
	AppName      string
	Environment  string
	AppPort      string
	DBConnection string
	APIVersion   string
	TokenSecret  string
}

func NewConfig() *Config {
	loadEnvironment()

	config := &Config{
		AppName:      strings.TrimRight(os.Getenv("APP_NAME"), "\n\r"),
		Environment:  strings.TrimRight(os.Getenv("APP_ENV"), "\n\r"),
		AppPort:      strings.TrimRight(os.Getenv("APP_PORT"), "\n\r"),
		DBConnection: strings.TrimRight(os.Getenv("DB_CONNECTION"), "\n\r"),
		APIVersion:   strings.TrimRight(os.Getenv("API_VERSION"), "\n\r"),
		TokenSecret:  strings.TrimRight(os.Getenv("TOKEN_SECRET"), "\n\r"),
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
	_ = os.Setenv("APP_NAME", "restaurant-api")
	_ = os.Setenv("APP_ENV", "default")
	_ = os.Setenv("APP_PORT", ":8080")
	_ = os.Setenv("DB_CONNECTION", "postgres://tech_challenge_usr:tech_challenge_pwd@localhost:5432/restaurant_db?sslmode=disable")
	_ = os.Setenv("API_VERSION", "4.0")
	_ = os.Setenv("TOKEN_SECRET", "123")
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
	fmt.Printf("App Name: %s\n", config.AppName)
	fmt.Printf("Environment: %s\n", config.Environment)
	fmt.Printf("App Port: %s\n", config.AppPort)
	fmt.Printf("DB Connection: %s\n", config.DBConnection)
	fmt.Printf("API version: %s\n", config.APIVersion)
	fmt.Printf("Token Secret: %s\n", config.TokenSecret)
}
