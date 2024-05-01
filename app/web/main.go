package main

import (
	"context"
	"fmt"
	"github.com/fabianogoes/fiap-challenge/adapters/delivery"
	"github.com/fabianogoes/fiap-challenge/adapters/payment"
	"github.com/fabianogoes/fiap-challenge/domain/usecases"
	"github.com/fabianogoes/fiap-challenge/frameworks/repository"
	"log/slog"
	"os"

	"github.com/fabianogoes/fiap-challenge/domain/entities"

	"github.com/fabianogoes/fiap-challenge/frameworks/rest"
)

func init() {
	fmt.Println("Initializing...")

	var logHandler *slog.JSONHandler

	config, _ := entities.NewConfig()
	if config.Environment == "production" {
		logHandler = slog.NewJSONHandler(os.Stdout, nil)
	} else {
		logHandler = slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
			Level: slog.LevelDebug,
		})
	}

	logger := slog.New(logHandler)
	slog.SetDefault(logger)
}

func main() {
	fmt.Println("Starting web server...")

	ctx := context.Background()
	var err error

	config, err := entities.NewConfig()
	if err != nil {
		panic(err)
	}
	db, err := repository.InitDB(ctx, config)
	if err != nil {
		panic(err)
	}

	attendantRepository := repository.NewAttendantRepository(db)
	attendantUseCase := usecases.NewAttendantService(attendantRepository)
	attendantHandler := rest.NewAttendantHandler(attendantUseCase)

	customerRepository := repository.NewCustomerRepository(db)
	customerUseCase := usecases.NewCustomerService(customerRepository)
	customerHandler := rest.NewCustomerHandler(customerUseCase, config)

	productRepository := repository.NewProductRepository(db)
	productUseCase := usecases.NewProductService(productRepository)
	productHandler := rest.NewProductHandler(productUseCase)

	paymentClientAdapter := payment.NewPaymentClientAdapter()
	paymentRepository := repository.NewPaymentRepository(db)
	paymentUseCase := usecases.NewPaymentService(paymentRepository)
	orderItemRepository := repository.NewOrderItemRepository(db)
	orderRepository := repository.NewOrderRepository(db, orderItemRepository)
	deliveryClientAdapter := delivery.NewDeliveryClientAdapter()
	deliveryRepository := repository.NewDeliveryRepository(db)
	orderUseCase := usecases.NewOrderService(
		orderRepository,
		customerRepository,
		attendantRepository,
		paymentUseCase,
		paymentClientAdapter,
		deliveryClientAdapter,
		deliveryRepository,
	)
	orderHandler := rest.NewOrderHandler(
		orderUseCase,
		customerUseCase,
		attendantUseCase,
		productUseCase,
	)

	router, err := rest.NewRouter(
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

	err = router.Run(config.AppPort)
	if err != nil {
		panic(err)
	}
}
