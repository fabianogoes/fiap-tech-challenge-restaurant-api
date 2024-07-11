package entities

import (
	"fmt"
	"github.com/joho/godotenv"
	"log/slog"
	"os"
)

type Config struct {
	AppName     string
	Environment string
	AppPort     string
	DBUser      string
	DBPassword  string
	DBHost      string
	DBPort      string
	DBName      string
	APIVersion  string
	TokenSecret string
}

func NewConfig() *Config {
	loadEnvironment()

	config := &Config{
		Environment: os.Getenv("APP_ENV"),
		AppPort:     os.Getenv("APP_PORT"),
		DBHost:      os.Getenv("DB_HOST"),
		DBPort:      os.Getenv("DB_PORT"),
		DBName:      os.Getenv("DB_DATABASE"),
		DBUser:      os.Getenv("DB_USERNAME"),
		DBPassword:  os.Getenv("DB_PASSWORD"),
		APIVersion:  os.Getenv("API_VERSION"),
		TokenSecret: os.Getenv("TOKEN_SECRET"),
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
	_ = os.Setenv("APP_PORT", ":8020")
	_ = os.Setenv("DB_HOST", "localhost")
	_ = os.Setenv("DB_PORT", "5432")
	_ = os.Setenv("DB_DATABASE", "restaurant_db")
	_ = os.Setenv("DB_USERNAME", "usr")
	_ = os.Setenv("DB_PASSWORD", "pwd")
	_ = os.Setenv("API_VERSION", "1.0")
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
	fmt.Printf("Environment: %s\n", config.Environment)
	fmt.Printf("App Port: %s\n", config.AppPort)
	fmt.Printf("DB Host: %s\n", config.DBHost)
	fmt.Printf("DB Port: %s\n", config.DBPort)
	fmt.Printf("DB Name: %s\n", config.DBName)
	fmt.Printf("DB User: %s\n", config.DBUser)
	fmt.Printf("DB Password: %s\n", config.DBPassword)
	fmt.Printf("API version: %s\n", config.APIVersion)
	fmt.Printf("Token Secret: %s\n", config.TokenSecret)
}
