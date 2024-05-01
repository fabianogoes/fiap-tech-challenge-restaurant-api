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

	appEnv := strings.Replace(os.Getenv("APP_ENV"), "\r", "", -1)
	appPort := strings.Replace(os.Getenv("APP_PORT"), "\r", "", -1)
	dbHost := strings.Replace(strings.Replace(os.Getenv("DB_HOST"), "\r", "", -1), "\n", "", -1)
	dbPort := strings.Replace(os.Getenv("DB_PORT"), "\r", "", -1)
	dbName := strings.Replace(strings.Replace(os.Getenv("DB_DATABASE"), "\r", "", -1), "\n", "", -1)
	dbUser := strings.Replace(strings.Replace(os.Getenv("DB_USERNAME"), "\r", "", -1), "\n", "", -1)
	dbPassword := strings.Replace(strings.Replace(os.Getenv("DB_PASSWORD"), "\r", "", -1), "\n", "", -1)
	apiVersion := strings.Replace(os.Getenv("API_VERSION"), "\r", "", -1)
	tokenSecret := strings.Replace(strings.Replace(os.Getenv("TOKEN_SECRET"), "\r", "", -1), "\n", "", -1)

	fmt.Println("APP_ENV=" + appEnv)
	fmt.Println("APP_PORT=" + appPort)
	fmt.Println("DB_HOST=" + dbHost)
	fmt.Println("DB_PORT=" + dbPort)
	fmt.Println("DB_DATABASE=" + dbName)
	fmt.Println("DB_USERNAME=" + dbUser)
	fmt.Println("DB_PASSWORD=" + dbPassword)
	fmt.Println("API_VERSION=" + apiVersion)
	fmt.Println("TOKEN_SECRET=" + tokenSecret)

	return &Config{
		Environment: appEnv,
		AppPort:     appPort,
		DBHost:      dbHost,
		DBPort:      dbPort,
		DBName:      dbName,
		DBUser:      dbUser,
		DBPassword:  dbPassword,
		APIVersion:  apiVersion,
		TokenSecret: tokenSecret,
	}, nil
}
