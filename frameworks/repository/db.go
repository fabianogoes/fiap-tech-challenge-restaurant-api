package repository

import (
	"fmt"
	"github.com/fabianogoes/fiap-challenge/domain/entities"
	"github.com/fabianogoes/fiap-challenge/frameworks/repository/dbo"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"time"
)

func InitDB(config *entities.Config) (*gorm.DB, error) {
	loc, _ := time.LoadLocation("UTC")

	var dsnTemplate string
	if config.Environment == "production" {
		dsnTemplate = "user=%s password=%s host=%s port=%s dbname=%s TimeZone=%s"
	} else {
		dsnTemplate = "user=%s password=%s host=%s port=%s dbname=%s sslmode=disable TimeZone=%s"
	}

	dsn := fmt.Sprintf(dsnTemplate,
		config.DBUser,
		config.DBPassword,
		config.DBHost,
		config.DBPort,
		config.DBName,
		loc,
	)
	fmt.Printf("dsn = %s\n", dsn)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
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
		&dbo.Outbox{},
	); err != nil {
		log.Fatal("AutoMigrate error", err)
		return nil, err
	}

	InitialDataAttendants(db)
	InitialDataCustomers(db)
	InitialDataProducts(db)

	return db, nil
}
