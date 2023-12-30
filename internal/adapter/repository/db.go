package repository

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

/**
 * DB is a wrapper for PostgreSQL database connection
 * that uses pgxpool as database driver
 */
type DB struct {
	*gorm.DB
}

// NewDB creates a new PostgreSQL database instance
func InitDB(ctx context.Context) (*DB, error) {
	loc, _ := time.LoadLocation("UTC")

	dsn := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=disable TimeZone=%s",
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_DATABASE"),
		loc,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Error connecting to database", err)
	}

	db.AutoMigrate(
		&Category{},
		&Product{},
		&Customer{},
		&Attendant{},
		&Order{},
		&OrderItem{},
		&Payment{},
		&Delivery{},
	)

	return &DB{
		db,
	}, nil
}
