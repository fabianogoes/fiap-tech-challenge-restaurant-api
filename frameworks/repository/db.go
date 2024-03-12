package repository

import (
	"context"
	"fmt"
	"github.com/fabianogoes/fiap-challenge/frameworks/repository/dbo"
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB(ctx context.Context) (*gorm.DB, error) {
	loc, _ := time.LoadLocation("UTC")

	dsn := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=disable TimeZone=%s",
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_DATABASE"),
		loc,
	)

	fmt.Printf("DB_HOST = %s\n", os.Getenv("DB_HOST"))
	fmt.Printf("DB_PORT = %s\n", os.Getenv("DB_PORT"))

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
	); err != nil {
		log.Fatal("AutoMigrate error", err)
		return nil, err
	}

	InitialDataAttendants(db)
	InitialDataCustomers(db)
	InitialDataProducts(db)

	return db, nil
}
