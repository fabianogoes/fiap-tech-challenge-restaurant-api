package entities

import (
	"fmt"
	"github.com/joho/godotenv"
	"log/slog"
	"os"
	"strings"
)

type Config struct {
	Environment string
	AppPort     string
	DBHost      string
	DBPort      string
	DBName      string
	DBUser      string
	DBPassword  string
	APIVersion  string
	TokenSecret string
}

func NewConfig() (*Config, error) {
	loadEnvironment()

	config := &Config{
		Environment: getEnv("APP_ENV", "development"),
		AppPort:     getEnv("APP_PORT", ":8080"),
		DBHost:      getEnv("DB_HOST", "localhost"),
		DBPort:      getEnv("DB_PORT", "5432"),
		DBName:      getEnv("DB_DATABASE", "db"),
		DBUser:      getEnv("DB_USERNAME", "usr"),
		DBPassword:  getEnv("DB_PASSWORD", "pwd"),
		APIVersion:  getEnv("API_VERSION", "2024.5.8.1"),
		TokenSecret: getEnv("TOKEN_SECRET", "123"),
	}

	printConfig(config)
	return config, nil
}

func loadEnvironment() {
	if os.Getenv("APP_ENV") == "production" {
		// Load .env file
		err := godotenv.Load()
		if err != nil {
			slog.Error("Error loading .env file", "error", err)
			os.Exit(1)
		}
	} else {
		// Load .env.development file
		err := godotenv.Load(".env.development")
		if err != nil {
			slog.Error("Error loading .env file", "error", err)
			os.Exit(1)
		}
	}
}

func getEnv(key, fallback string) string {
	if envValue, ok := os.LookupEnv(key); ok {
		return strings.TrimRight(envValue, "\n\r")
	}

	return fallback
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
