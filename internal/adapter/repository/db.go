package repository

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/fiap/challenge-gofood/internal/adapter/repository/dbo"
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

	db.AutoMigrate(
		&dbo.Category{},
		&dbo.Product{},
		&dbo.Customer{},
		&dbo.Attendant{},
		&dbo.Order{},
		&dbo.OrderItem{},
		&dbo.Payment{},
		&dbo.Delivery{},
	)

	db.Create(&dbo.Attendant{Name: "Miguel"})
	db.Create(&dbo.Attendant{Name: "Sophia"})
	db.Create(&dbo.Attendant{Name: "Alice"})
	db.Create(&dbo.Attendant{Name: "Pedro"})
	db.Create(&dbo.Attendant{Name: "Manuela"})

	db.Create(&dbo.Customer{Name: "Bernardo", Email: "bernardo@gmail.com", CPF: "29381510040"})
	db.Create(&dbo.Customer{Name: "Laura", Email: "laura@hotmail.com", CPF: "15204180001"})
	db.Create(&dbo.Customer{Name: "Lucas", Email: "lucas@gmail.com", CPF: "43300921074"})
	db.Create(&dbo.Customer{Name: "Maria Eduarda", Email: "meduarda@uol.com.br", CPF: "85752055016"})
	db.Create(&dbo.Customer{Name: "Guilherme", Email: "guilherme@microsoft.com", CPF: "17148604001"})

	db.Create(&dbo.Category{Name: "Sanduíches"})
	db.Create(&dbo.Category{Name: "Bebidas Frias"})
	db.Create(&dbo.Category{Name: "Bebidas Quentes"})
	db.Create(&dbo.Category{Name: "Combos"})
	db.Create(&dbo.Category{Name: "Sobremesas"})
	db.Create(&dbo.Category{Name: "Acompanhamentos"})
	db.Create(&dbo.Category{Name: "Café da Manhã"})

	db.Create(&dbo.Product{Name: "Big Lanche", Price: 29.90, CategoryID: 1})
	db.Create(&dbo.Product{Name: "X-Burguer", Price: 19.90, CategoryID: 1})
	db.Create(&dbo.Product{Name: "X-Salada", Price: 21.90, CategoryID: 1})
	db.Create(&dbo.Product{Name: "X-Bacon", Price: 23.90, CategoryID: 1})
	db.Create(&dbo.Product{Name: "X-Tudo", Price: 27.90, CategoryID: 1})
	db.Create(&dbo.Product{Name: "Coca-Cola", Price: 5.90, CategoryID: 2})
	db.Create(&dbo.Product{Name: "Guaraná", Price: 5.90, CategoryID: 2})
	db.Create(&dbo.Product{Name: "Fanta", Price: 5.90, CategoryID: 2})
	db.Create(&dbo.Product{Name: "Suco de Laranja", Price: 5.90, CategoryID: 2})
	db.Create(&dbo.Product{Name: "Suco de Uva", Price: 5.90, CategoryID: 2})
	db.Create(&dbo.Product{Name: "Café", Price: 3.90, CategoryID: 3})
	db.Create(&dbo.Product{Name: "Cappuccino", Price: 4.90, CategoryID: 3})
	db.Create(&dbo.Product{Name: "Chocolate Quente", Price: 4.90, CategoryID: 3})
	db.Create(&dbo.Product{Name: "Misto Quente + Fritas", Price: 9.90, CategoryID: 4})
	db.Create(&dbo.Product{Name: "X-Burguer + Fritas + Coca-Cola", Price: 29.90, CategoryID: 4})
	db.Create(&dbo.Product{Name: "X-Salada + Fritas + Coca-Cola", Price: 31.90, CategoryID: 4})
	db.Create(&dbo.Product{Name: "X-Bacon + Fritas + Coca-Cola", Price: 33.90, CategoryID: 4})
	db.Create(&dbo.Product{Name: "X-Tudo + Fritas + Coca-Cola", Price: 37.90, CategoryID: 4})
	db.Create(&dbo.Product{Name: "Sorvete", Price: 7.90, CategoryID: 5})
	db.Create(&dbo.Product{Name: "Sundae", Price: 9.90, CategoryID: 5})
	db.Create(&dbo.Product{Name: "Açaí", Price: 11.90, CategoryID: 5})
	db.Create(&dbo.Product{Name: "Batata Frita", Price: 9.90, CategoryID: 6})
	db.Create(&dbo.Product{Name: "Batata Frita com Cheddar", Price: 11.90, CategoryID: 6})
	db.Create(&dbo.Product{Name: "Batata Frita com Cheddar e Bacon", Price: 13.90, CategoryID: 6})
	db.Create(&dbo.Product{Name: "Batata Frita com Cheddar e Calabresa", Price: 13.90, CategoryID: 6})
	db.Create(&dbo.Product{Name: "Batata Frita com Cheddar e Frango", Price: 13.90, CategoryID: 6})

	return db, nil
}
