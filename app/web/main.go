package main

import (
	"fmt"
	"github.com/fabianogoes/fiap-challenge/adapters/messaging"
	"github.com/fabianogoes/fiap-challenge/frameworks/scheduler"
	"github.com/fabianogoes/fiap-challenge/shared"
	"log/slog"
	"os"

	"github.com/fabianogoes/fiap-challenge/adapters/delivery"
	"github.com/fabianogoes/fiap-challenge/adapters/payment"
	"github.com/fabianogoes/fiap-challenge/domain/usecases"
	"github.com/fabianogoes/fiap-challenge/frameworks/repository"

	"github.com/fabianogoes/fiap-challenge/domain/entities"

	"github.com/fabianogoes/fiap-challenge/frameworks/rest"
)

func init() {
	fmt.Println("Initializing...")

	var logHandler *slog.JSONHandler

	config := entities.NewConfig()
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

	var err error

	config := entities.NewConfig()
	crypto := shared.NewCrypto([]byte(config.CryptoKey))
	db, err := repository.InitDB(config)
	if err != nil {
		fmt.Printf("error while initializing database %v", err)
		panic(err)
	}
	fmt.Println("DB connected successfully")

	sqsClient := messaging.NewAWSSQSClient(config)
	attendantRepository := repository.NewAttendantRepository(db, crypto)
	attendantUseCase := usecases.NewAttendantService(attendantRepository)
	attendantHandler := rest.NewAttendantHandler(attendantUseCase)

	customerRepository := repository.NewCustomerRepository(db, crypto)
	customerUseCase := usecases.NewCustomerService(customerRepository)
	customerHandler := rest.NewCustomerHandler(customerUseCase, config)

	productRepository := repository.NewProductRepository(db)
	productUseCase := usecases.NewProductService(productRepository)
	productHandler := rest.NewProductHandler(productUseCase)

	paymentClientAdapter := payment.NewPaymentClientAdapter(config)
	paymentRepository := repository.NewPaymentRepository(db)
	paymentUseCase := usecases.NewPaymentService(paymentRepository)
	orderItemRepository := repository.NewOrderItemRepository(db)
	orderRepository := repository.NewOrderRepository(db, orderItemRepository)
	deliveryClientAdapter := delivery.NewDeliveryClientAdapter()
	deliveryRepository := repository.NewDeliveryRepository(db)
	outboxRepository := repository.NewOutboxRepository(db)
	kitchenPublisher := messaging.NewKitchenPublisher(sqsClient, outboxRepository)
	paymentMessaging := messaging.NewPaymentPublisher(sqsClient, outboxRepository)
	orderUseCase := usecases.NewOrderService(
		orderRepository,
		customerRepository,
		attendantRepository,
		paymentUseCase,
		paymentClientAdapter,
		deliveryClientAdapter,
		deliveryRepository,
		kitchenPublisher,
		paymentMessaging,
	)
	orderHandler := rest.NewOrderHandler(
		orderUseCase,
		customerUseCase,
		attendantUseCase,
		productUseCase,
	)

	paymentReceiver := messaging.NewPaymentReceiver(orderUseCase, config, sqsClient)
	kitchenReceiver := messaging.NewKitchenReceiver(orderUseCase, config, sqsClient)
	outboxRetry := messaging.NewOutboxRetry(sqsClient, outboxRepository)
	cron := scheduler.InitCronScheduler(paymentReceiver, kitchenReceiver, outboxRetry)
	defer cron.Stop()

	router, err := rest.NewRouter(
		customerHandler,
		attendantHandler,
		productHandler,
		orderHandler,
	)
	if err != nil {
		fmt.Printf("error while initializing router %v", err)
		panic(err)
	}

	err = router.Run(config.AppPort)
	if err != nil {
		fmt.Printf("error while starting web server %v", err)
		panic(err)
	}
}
