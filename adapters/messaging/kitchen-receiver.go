package messaging

import (
	"fmt"
	"github.com/fabianogoes/fiap-challenge/domain/entities"
	"github.com/fabianogoes/fiap-challenge/domain/ports"
	"github.com/goccy/go-json"
	"log/slog"
)

type KitchenReceiver struct {
	OrderUseCase ports.OrderUseCasePort
	config       *entities.Config
	awsSqsClient *AWSSQSClient
}

func NewKitchenReceiver(
	orderUseCase ports.OrderUseCasePort,
	config *entities.Config,
	awsSqsClient *AWSSQSClient,
) *KitchenReceiver {
	return &KitchenReceiver{
		orderUseCase,
		config,
		awsSqsClient,
	}
}

func (r *KitchenReceiver) ReceiveKitchenCallback() {
	slog.Info("running kitchen listener...")
	if messages := r.awsSqsClient.Receive(r.config.KitchenCallbackQueueUrl); messages != nil {

		for _, message := range messages.Messages {
			slog.Info("message received", "message", *message.Body)
			kitchenCallbackDTO := toKitchenCallbackDTO(*message.Body)
			order, err := r.OrderUseCase.GetOrderById(kitchenCallbackDTO.OrderId)
			if err != nil {
				slog.Error(fmt.Sprintf("kitchen error on get order id: %v - %v", kitchenCallbackDTO.OrderId, err))
				continue
			}

			if err := r.handleCallback(order, kitchenCallbackDTO.Status); err == nil {
				r.awsSqsClient.Delete(message.ReceiptHandle, r.config.KitchenCallbackQueueUrl)
			}
		}
	}
}

func (r *KitchenReceiver) handleCallback(order *entities.Order, kitchenStatus string) error {
	status := entities.ToOrderStatus(kitchenStatus)
	slog.Info(fmt.Sprintf("kitchen callback error getting order by id: %v - %v", order, status))
	if status == entities.OrderStatusKitchenReady {
		_, err := r.OrderUseCase.ReadyForDeliveryOrder(order)
		if err != nil {
			slog.Error(fmt.Sprintf("kitchen callback error getting order by id: %v - %v", order, err))
			return err
		}
	}
	return nil
}

func toKitchenCallbackDTO(jsonData string) *KitchenCallbackDTO {
	var kitchenCallbackDTO KitchenCallbackDTO
	err := json.Unmarshal([]byte(jsonData), &kitchenCallbackDTO)
	if err != nil {
		slog.Error(fmt.Sprintf("error unmarshalling json data - %v", err))
		return nil
	}

	return &kitchenCallbackDTO
}

type KitchenCallbackDTO struct {
	OrderId uint   `json:"orderId"`
	Status  string `json:"status"`
}
