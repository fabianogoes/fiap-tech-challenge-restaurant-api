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

	config := &Config{
		Environment: getEnv("APP_ENV", "development"),
		AppPort:     getEnv("APP_PORT", ":8080"),
		DBHost:      getEnv("DB_HOST", "localhost"),
		DBPort:      getEnv("DB_PORT", "5432"),
		DBName:      getEnv("DB_DATABASE", "db"),
		DBUser:      getEnv("DB_USERNAME", "usr"),
		DBPassword:  getEnv("DB_PASSWORD", "pwd"),
		APIVersion:  getEnv("API_VERSION", "v1"),
		TokenSecret: getEnv("TOKEN_SECRET", "123"),
	}

	fmt.Println(config)
	return config, nil
}

func getEnv(key, fallback string) string {
	if envValue, ok := os.LookupEnv(key); ok {
		return strings.TrimRight(envValue, "\n\r")
	}

	return fallback
}
