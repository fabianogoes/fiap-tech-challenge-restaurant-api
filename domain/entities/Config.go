package entities

import (
	"fmt"
	"github.com/joho/godotenv"
	"log/slog"
	"os"
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

func NewConfig() *Config {
	loadEnvironment()

	config := &Config{
		Environment: strings.TrimRight(os.Getenv("APP_ENV"), "\n\r"),
		AppPort:     strings.TrimRight(os.Getenv("APP_PORT"), "\n\r"),
		DBHost:      strings.TrimRight(os.Getenv("DB_HOST"), "\n\r"),
		DBPort:      strings.TrimRight(os.Getenv("DB_PORT"), "\n\r"),
		DBName:      strings.TrimRight(os.Getenv("DB_DATABASE"), "\n\r"),
		DBUser:      strings.TrimRight(os.Getenv("DB_USERNAME"), "\n\r"),
		DBPassword:  strings.TrimRight(os.Getenv("DB_PASSWORD"), "\n\r"),
		APIVersion:  strings.TrimRight(os.Getenv("API_VERSION"), "\n\r"),
		TokenSecret: strings.TrimRight(os.Getenv("TOKEN_SECRET"), "\n\r"),
	}

	printConfig(config)
	return config
}

func loadEnvironment() {
	if os.Getenv("APP_ENV") == "production" {
		// Load .env file
		err := godotenv.Load()
		if err != nil {
			slog.Error("Error loading .env file", "error", err)
			os.Exit(1)
		}
	} else if os.Getenv("APP_ENV") == "development" {
		// Load .env.development file
		err := godotenv.Load(".env.development")
		if err != nil {
			slog.Error("Error loading .env file", "error", err)
			os.Exit(1)
		}
	} else {
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
