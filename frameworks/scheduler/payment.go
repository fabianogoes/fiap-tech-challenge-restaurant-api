package scheduler

import (
	"github.com/fabianogoes/fiap-challenge/adapters/messaging"
	"github.com/fabianogoes/fiap-challenge/domain/ports"
	"github.com/robfig/cron"
	"log/slog"
)

func InitCronScheduler(paymentReceiver ports.PaymentReceiverPort, kitchenReceiver *messaging.KitchenReceiver) *cron.Cron {
	job := cron.New()

	_ = job.AddFunc("*/5 * * * *", paymentReceiver.ReceivePaymentCallback)
	_ = job.AddFunc("*/5 * * * *", kitchenReceiver.ReceiveKitchenCallback)

	job.Start()
	slog.Info("cron scheduler initialized")

	return job
}
