package messaging

import (
	"fmt"
	"github.com/fabianogoes/fiap-challenge/domain/entities"
	"github.com/fabianogoes/fiap-challenge/domain/ports"
	"github.com/goccy/go-json"
	"log/slog"
)

type PaymentReceiver struct {
	OrderUseCase ports.OrderUseCasePort
	config       *entities.Config
	awsSqsClient *AWSSQSClient
}

func NewPaymentReceiver(
	orderUseCase ports.OrderUseCasePort,
	config *entities.Config,
	awsSqsClient *AWSSQSClient,
) *PaymentReceiver {
	return &PaymentReceiver{
		orderUseCase,
		config,
		awsSqsClient,
	}
}

func (r *PaymentReceiver) ReceivePaymentCallback() {
	slog.Info("running payment listener...")
	if messages := r.awsSqsClient.Receive(r.config.PaymentCallbackQueueUrl); messages != nil {

		for _, message := range messages.Messages {
			slog.Info("message received", "message", *message.Body)
			paymentCallbackDTO := toPaymentCallbackDTO(*message.Body)
			order, err := r.OrderUseCase.GetOrderById(paymentCallbackDTO.OrderId)
			if err != nil {
				slog.Error(fmt.Sprintf("payment callback error getting order by id: %v - %v", paymentCallbackDTO.OrderId, err))
				continue
			}

			callbackStatus := entities.ToPaymentStatus(paymentCallbackDTO.Status)
			switch callbackStatus {
			case entities.PaymentStatusPaid:
				r.handleCallbackPaid(order)
			case entities.PaymentStatusError:
				r.handleCallbackError(order, callbackStatus)
			default:
				slog.Info(fmt.Sprintf("order %v payment %v callback with status %v\n", paymentCallbackDTO.OrderId, paymentCallbackDTO.PaymentId, callbackStatus.ToString()))
			}

			r.awsSqsClient.Delete(message.ReceiptHandle, r.config.PaymentCallbackQueueUrl)
		}
	}
}

func (r *PaymentReceiver) handleCallbackPaid(order *entities.Order) {
	orderConfirmed, err := r.OrderUseCase.PaymentOrderConfirmed(order, order.Payment.Method.ToString())
	if err != nil {
		slog.Error(fmt.Sprintf("payment callback orderID %v error to confirmed - %v", order.ID, err))
	}

	if _, err := r.OrderUseCase.InPreparationOrder(orderConfirmed); err != nil {
		slog.Error(fmt.Sprintf("order %v preparation error - %v", order.ID, err))
	}
}

func (r *PaymentReceiver) handleCallbackError(order *entities.Order, callbackStatus entities.PaymentStatus) {
	_, err := r.OrderUseCase.PaymentOrderError(order, order.Payment.Method.ToString(), callbackStatus.ToString())
	if err != nil {
		slog.Error(fmt.Sprintf("payment callback orderID %v error to error - %v", order.ID, err))
	}
}

func toPaymentCallbackDTO(jsonData string) *PaymentCallbackDTO {
	var paymentCallbackDTO PaymentCallbackDTO
	err := json.Unmarshal([]byte(jsonData), &paymentCallbackDTO)
	if err != nil {
		slog.Error(fmt.Sprintf("error unmarshalling json data - %v", err))
		return nil
	}

	return &paymentCallbackDTO
}

type PaymentCallbackDTO struct {
	OrderId   uint   `json:"orderId"`
	PaymentId string `json:"paymentId"`
	Status    string `json:"status"`
}
