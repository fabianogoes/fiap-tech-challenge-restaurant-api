package scheduler

import (
	"github.com/fabianogoes/fiap-challenge/domain/ports"
	"github.com/robfig/cron"
	"log/slog"
)

func InitCronScheduler(paymentReceiver ports.PaymentReceiverPort) *cron.Cron {
	job := cron.New()

	_ = job.AddFunc("*/5 * * * *", paymentReceiver.ReceiveCallback)

	job.Start()
	slog.Info("cron scheduler initialized")

	return job
}
