package main

import (
	"context"
	"fmt"
	"github.com/fabianogoes/fiap-challenge/domain/usecases"
	repository2 "github.com/fabianogoes/fiap-challenge/frameworks/repository"
	controllers2 "github.com/fabianogoes/fiap-challenge/frameworks/rest"
	"log/slog"
	"os"

	"github.com/fabianogoes/fiap-challenge/adapters/delivery"
	"github.com/fabianogoes/fiap-challenge/adapters/payment"
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

	db, err := repository2.InitDB(ctx)
	if err != nil {
		panic(err)
	}

	attendantRepository := repository2.NewAttendantRepository(db)
	attendantUseCase := usecases.NewAttendantService(attendantRepository)
	attendantHandler := controllers2.NewAttendantHandler(attendantUseCase)

	customerRepository := repository2.NewCustomerRepository(db)
	customerUseCase := usecases.NewCustomerService(customerRepository)
	customerHandler := controllers2.NewCustomerHandler(customerUseCase)

	productRepository := repository2.NewProductRepository(db)
	productUseCase := usecases.NewProductService(productRepository)
	productHandler := controllers2.NewProductHandler(productUseCase)

	paymentClientAdapter := payment.NewPaymentClientAdapter()
	paymentRepository := repository2.NewPaymentRepository(db)
	paymentUseCase := usecases.NewPaymentService(paymentRepository)
	orderItemRepository := repository2.NewOrderItemRepository(db)
	orderRepository := repository2.NewOrderRepository(db, orderItemRepository)
	deliveryClientAdapter := delivery.NewDeliveryClientAdapter()
	deliveryRepository := repository2.NewDeliveryRepository(db)
	orderUseCase := usecases.NewOrderService(
		orderRepository,
		customerRepository,
		attendantRepository,
		paymentUseCase,
		paymentClientAdapter,
		deliveryClientAdapter,
		deliveryRepository,
	)
	orderHandler := controllers2.NewOrderHandler(
		orderUseCase,
		customerUseCase,
		attendantUseCase,
		productUseCase,
	)

	router, err := controllers2.NewRouter(
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

	err = router.Run(os.Getenv("HTTP_PORT"))
	if err != nil {
		panic(err)
	}
}
