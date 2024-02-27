package main

import (
	"context"
	"fmt"
	service2 "github.com/fiap/challenge-gofood/usecases"
	"log/slog"
	"os"

	"github.com/fiap/challenge-gofood/adapters/delivery"
	"github.com/fiap/challenge-gofood/adapters/handler"
	"github.com/fiap/challenge-gofood/adapters/payment"
	"github.com/fiap/challenge-gofood/adapters/repository"
	"github.com/joho/godotenv"
)

func init() {
	fmt.Println("Initializing...")

	var logHandler *slog.JSONHandler

	env := getAppEnv()

	if env == "production" {
		logHandler = slog.NewJSONHandler(os.Stdout, nil)

		// Load .env file
		err := godotenv.Load()
		if err != nil {
			slog.Error("Error loading .env file", "error", err)
			os.Exit(1)
		}
	} else {
		logHandler = slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
			Level: slog.LevelDebug,
		})

		// Load .env.development file
		err := godotenv.Load(".env.development")
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
	attendantUseCase := service2.NewAttendantService(attendantRepository)
	attendantHandler := handler.NewAttendantHandler(attendantUseCase)

	customerRepository := repository.NewCustomerRepository(db)
	customerUseCase := service2.NewCustomerService(customerRepository)
	customerHandler := handler.NewCustomerHandler(customerUseCase)

	productRepository := repository.NewProductRepository(db)
	productUseCase := service2.NewProductService(productRepository)
	productHandler := handler.NewProductHandler(productUseCase)

	paymentClientAdapter := payment.NewPaymentClientAdapter()
	paymentRepository := repository.NewPaymentRepository(db)
	paymentUseCase := service2.NewPaymentService(paymentRepository)
	orderItemRepository := repository.NewOrderItemRepository(db)
	orderRepository := repository.NewOrderRepository(db, orderItemRepository)
	deliveryClientAdapter := delivery.NewDeliveryClientAdapter()
	deliveryRepository := repository.NewDeliveryRepository(db)
	orderUseCase := service2.NewOrderService(
		orderRepository,
		customerRepository,
		attendantRepository,
		paymentUseCase,
		paymentClientAdapter,
		deliveryClientAdapter,
		deliveryRepository,
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
