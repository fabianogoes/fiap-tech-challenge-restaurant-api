package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"

	"github.com/fiap/challenge-gofood/internal/adapter/delivery"
	"github.com/fiap/challenge-gofood/internal/adapter/handler"
	"github.com/fiap/challenge-gofood/internal/adapter/payment"
	"github.com/fiap/challenge-gofood/internal/adapter/repository"
	"github.com/fiap/challenge-gofood/internal/domain/service"
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
	var err error

	db, err := repository.InitDB(ctx)
	if err != nil {
		panic(err)
	}

	attendantRepository := repository.NewAttendantRepository(db)
	attendantUseCase := service.NewAttendantService(attendantRepository)
	attendantHandler := handler.NewAttendantHandler(attendantUseCase)

	customerRepository := repository.NewCustomerRepository(db)
	customerUseCase := service.NewCustomerService(customerRepository)
	customerHandler := handler.NewCustomerHandler(customerUseCase)

	productRepository := repository.NewProductRepository(db)
	productUseCase := service.NewProductService(productRepository)
	productHandler := handler.NewProductHandler(productUseCase)

	deliveryClientAdapter := delivery.NewDeliveryClientAdapter()
	paymentClientAdapter := payment.NewPaymentClientAdapter()
	paymentRepository := repository.NewPaymentRepository(db)
	paymentUseCase := service.NewPaymentService(paymentRepository)
	orderRepository := repository.NewOrderRepository(db)
	orderUseCase := service.NewOrderService(
		orderRepository,
		customerRepository,
		attendantRepository,
		paymentUseCase,
		paymentClientAdapter,
		deliveryClientAdapter,
	)
	orderHandler := handler.NewOrderHandler(
		orderUseCase,
		customerUseCase,
		attendantUseCase,
		productUseCase,
	)

	router, err := handler.NewRouter(
		customerHandler,
		attendantHandler,
		productHandler,
		orderHandler,
	)
	if err != nil {
		panic(err)
	}

	fmt.Println("DB connected")
	fmt.Println(db)

	router.Run(os.Getenv("HTTP_PORT"))
}
