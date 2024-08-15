package repository

import (
	"fmt"
	"github.com/fabianogoes/fiap-challenge/domain/entities"
	"github.com/fabianogoes/fiap-challenge/frameworks/repository/dbo"

	"gorm.io/gorm"
)

type OutboxRepository struct {
	db *gorm.DB
}

func NewOutboxRepository(db *gorm.DB) *OutboxRepository {
	return &OutboxRepository{db}
}

func (or *OutboxRepository) GetOutboxById(id uint) (*entities.Outbox, error) {
	var outbox dbo.Outbox

	if err := or.db.Where("id = ?", id).First(&outbox).Error; err != nil {
		return nil, fmt.Errorf("error to find outbox with id %or - %v\n", id, err)
	}

	return outbox.ToEntity(), nil
}

func (or *OutboxRepository) CreateOutbox(orderID uint, messageBody string, queueUrl string) (*entities.Outbox, error) {
	outbox := dbo.ToOutboxDBO(orderID, messageBody, queueUrl)

	if err := or.db.Create(&outbox).Error; err != nil {
		return nil, err
	}

	return outbox.ToEntity(), nil
}

func (or *OutboxRepository) DeleteOutbox(id uint) error {
	if err := or.db.Delete(&dbo.Outbox{}, id).Error; err != nil {
		return err
	}

	return nil
}

func (or *OutboxRepository) GetAll() ([]*entities.Outbox, error) {
	var results []*dbo.Outbox
	if err := or.db.Find(&results).Error; err != nil {
		return nil, err
	}

	var list []*entities.Outbox
	for _, result := range results {
		list = append(list, result.ToEntity())
	}

	return list, nil
}
