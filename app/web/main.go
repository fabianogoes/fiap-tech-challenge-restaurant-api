package main

import (
	"fmt"
	"github.com/fabianogoes/fiap-challenge/domain/entities"
	"log/slog"
	"os"

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

	// ctx := context.Background()
	var err error

	config, err := entities.NewConfig()
	if err != nil {
		panic(err)
	}
	// db, err := repository2.InitDB(ctx)
	// if err != nil {
	// 	panic(err)
	// }

	// attendantRepository := repository2.NewAttendantRepository(db)
	// attendantUseCase := usecases.NewAttendantService(attendantRepository)
	// attendantHandler := controllers2.NewAttendantHandler(attendantUseCase)

	// customerRepository := repository2.NewCustomerRepository(db)
	// customerUseCase := usecases.NewCustomerService(customerRepository)
	// customerHandler := controllers2.NewCustomerHandler(customerUseCase)

	// productRepository := repository2.NewProductRepository(db)
	// productUseCase := usecases.NewProductService(productRepository)
	// productHandler := controllers2.NewProductHandler(productUseCase)

	// paymentClientAdapter := payment.NewPaymentClientAdapter()
	// paymentRepository := repository2.NewPaymentRepository(db)
	// paymentUseCase := usecases.NewPaymentService(paymentRepository)
	// orderItemRepository := repository2.NewOrderItemRepository(db)
	// orderRepository := repository2.NewOrderRepository(db, orderItemRepository)
	// deliveryClientAdapter := delivery.NewDeliveryClientAdapter()
	// deliveryRepository := repository2.NewDeliveryRepository(db)
	// orderUseCase := usecases.NewOrderService(
	// 	orderRepository,
	// 	customerRepository,
	// 	attendantRepository,
	// 	paymentUseCase,
	// 	paymentClientAdapter,
	// 	deliveryClientAdapter,
	// 	deliveryRepository,
	// )
	// orderHandler := controllers2.NewOrderHandler(
	// 	orderUseCase,
	// 	customerUseCase,
	// 	attendantUseCase,
	// 	productUseCase,
	// )

	router, err := rest.NewRouter(
		// customerHandler,
		// attendantHandler,
		// productHandler,
		// orderHandler,
	)
	if err != nil {
		panic(err)
	}

	// fmt.Println("DB connected")
	// fmt.Println(db)

	err = router.Run(config.AppPort)
	if err != nil {
		panic(err)
	}
}
