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

func (r *PaymentReceiver) ReceiveCallback() {
	slog.Info("running payment listener...")
	if messages := r.awsSqsClient.Receive(r.config.PaymentCallbackQueueUrl); messages != nil {

		for _, message := range messages.Messages {
			slog.Info("message received", "message", *message.Body)
			paymentCallbackDTO, err := toPaymentCallbackDTO(*message.Body)
			if err != nil {
				slog.Error(fmt.Sprintf("error converting message to PaymentCallbackDTO - %v", err))
				continue
			}

			order, err := r.OrderUseCase.GetOrderById(paymentCallbackDTO.OrderId)
			if err != nil {
				slog.Error(fmt.Sprintf("payment callback error getting order by id: %v - %v", paymentCallbackDTO.OrderId, err))
				continue
			}

			callbackStatus := entities.ToPaymentStatus(paymentCallbackDTO.Status)
			paymentMethod := order.Payment.Method.ToString()
			if callbackStatus == entities.PaymentStatusPaid {
				if _, err = r.OrderUseCase.PaymentOrderConfirmed(order, paymentMethod); err != nil {
					slog.Error(fmt.Sprintf("payment callback orderID %v error to confirmed - %v", order.ID, err))
					continue
				} else {
					r.awsSqsClient.Delete(message.ReceiptHandle, r.config.PaymentCallbackQueueUrl)
				}
			} else {
				_, err := r.OrderUseCase.PaymentOrderError(order, paymentMethod, callbackStatus.ToString())
				if err != nil {
					slog.Error(fmt.Sprintf("payment callback orderID %v error to error - %v", order.ID, err))
					continue
				}
			}

		}
	}
}

func toPaymentCallbackDTO(jsonData string) (*PaymentCallbackDTO, error) {
	var paymentCallbackDTO PaymentCallbackDTO
	err := json.Unmarshal([]byte(jsonData), &paymentCallbackDTO)
	if err != nil {
		return nil, err
	}
	return &paymentCallbackDTO, nil
}

type PaymentCallbackDTO struct {
	OrderId   uint   `json:"orderId"`
	PaymentId string `json:"paymentId"`
	Status    string `json:"status"`
}
