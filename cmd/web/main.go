package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"

	"github.com/fiap/challenge-gofood/internal/adapter/repository"
	"github.com/joho/godotenv"
)

func init() {
	fmt.Println("Initializing...")

	var logHandler *slog.JSONHandler

	env := getAppEnv()

	if env == "production" {
		logHandler = slog.NewJSONHandler(os.Stdout, nil)
	} else {
		logHandler = slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
			Level: slog.LevelDebug,
		})

		// Load .env file
		err := godotenv.Load()
		if err != nil {
			slog.Error("Error loading .env file", "error", err)
			os.Exit(1)
		}
	}

	fmt.Printf("env = %s\n", env)

	logger := slog.New(logHandler)
	slog.SetDefault(logger)
}

func getAppEnv() string {
	env := os.Getenv("APP_ENV")

	if env == "" {
		env = "development"
	}

	return env
}

func main() {
	fmt.Println("Starting web server...")

	ctx := context.Background()

	db, err := repository.InitDB(ctx)
	if err != nil {
		panic(err)
	}
	fmt.Println("DB connected")
	fmt.Println(db)
}
