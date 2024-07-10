package repository

import (
	"fmt"
	"github.com/fabianogoes/fiap-challenge/domain/entities"
	"github.com/fabianogoes/fiap-challenge/frameworks/repository/dbo"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func InitDB(config *entities.Config) (*gorm.DB, error) {
	fmt.Printf("DB_CONNECTION = %s\n", config.DBConnection)

	db, err := gorm.Open(postgres.Open(config.DBConnection), &gorm.Config{})
	if err != nil {
		log.Fatal("Error connecting to database", err)
	}

	if err = db.AutoMigrate(
		&dbo.Category{},
		&dbo.Product{},
		&dbo.Customer{},
		&dbo.Attendant{},
		&dbo.Order{},
		&dbo.OrderItem{},
		&dbo.Payment{},
		&dbo.Delivery{},
	); err != nil {
		log.Fatal("AutoMigrate error", err)
		return nil, err
	}

	InitialDataAttendants(db)
	InitialDataCustomers(db)
	InitialDataProducts(db)

	return db, nil
}
