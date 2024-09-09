package ports

import "github.com/fabianogoes/fiap-challenge/domain/entities"

type KitchenPublisherPort interface {
	PublishKitchen(order *entities.Order) error
}

type KitchenReceiverPort interface {
	ReceiveKitchenCallback()
}
