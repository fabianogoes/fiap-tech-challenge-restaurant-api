package repository

import (
	"context"
	"fmt"
	"github.com/fabianogoes/fiap-challenge/domain/entities"
	dbo2 "github.com/fabianogoes/fiap-challenge/frameworks/repository/dbo"
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB(ctx context.Context, config *entities.Config) (*gorm.DB, error) {
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

	fmt.Printf("DB_HOST = %s\n", os.Getenv("DB_HOST"))

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Error connecting to database", err)
	}

	db.AutoMigrate(
		&dbo2.Category{},
		&dbo2.Product{},
		&dbo2.Customer{},
		&dbo2.Attendant{},
		&dbo2.Order{},
		&dbo2.OrderItem{},
		&dbo2.Payment{},
		&dbo2.Delivery{},
	)

	db.Create(&dbo2.Attendant{Name: "Miguel"})
	db.Create(&dbo2.Attendant{Name: "Sophia"})
	db.Create(&dbo2.Attendant{Name: "Alice"})
	db.Create(&dbo2.Attendant{Name: "Pedro"})
	db.Create(&dbo2.Attendant{Name: "Manuela"})

	db.Create(&dbo2.Customer{Name: "Bernardo", Email: "bernardo@gmail.com", CPF: "29381510040"})
	db.Create(&dbo2.Customer{Name: "Laura", Email: "laura@hotmail.com", CPF: "15204180001"})
	db.Create(&dbo2.Customer{Name: "Lucas", Email: "lucas@gmail.com", CPF: "43300921074"})
	db.Create(&dbo2.Customer{Name: "Maria Eduarda", Email: "meduarda@uol.com.br", CPF: "85752055016"})
	db.Create(&dbo2.Customer{Name: "Guilherme", Email: "guilherme@microsoft.com", CPF: "17148604001"})

	db.Create(&dbo2.Category{Name: "Sanduíches"})
	db.Create(&dbo2.Category{Name: "Bebidas Frias"})
	db.Create(&dbo2.Category{Name: "Bebidas Quentes"})
	db.Create(&dbo2.Category{Name: "Combos"})
	db.Create(&dbo2.Category{Name: "Sobremesas"})
	db.Create(&dbo2.Category{Name: "Acompanhamentos"})
	db.Create(&dbo2.Category{Name: "Café da Manhã"})

	db.Create(&dbo2.Product{Name: "Big Lanche", Price: 29.90, CategoryID: 1})
	db.Create(&dbo2.Product{Name: "X-Burguer", Price: 19.90, CategoryID: 1})
	db.Create(&dbo2.Product{Name: "X-Salada", Price: 21.90, CategoryID: 1})
	db.Create(&dbo2.Product{Name: "X-Bacon", Price: 23.90, CategoryID: 1})
	db.Create(&dbo2.Product{Name: "X-Tudo", Price: 27.90, CategoryID: 1})
	db.Create(&dbo2.Product{Name: "Coca-Cola", Price: 5.90, CategoryID: 2})
	db.Create(&dbo2.Product{Name: "Guaraná", Price: 5.90, CategoryID: 2})
	db.Create(&dbo2.Product{Name: "Fanta", Price: 5.90, CategoryID: 2})
	db.Create(&dbo2.Product{Name: "Suco de Laranja", Price: 5.90, CategoryID: 2})
	db.Create(&dbo2.Product{Name: "Suco de Uva", Price: 5.90, CategoryID: 2})
	db.Create(&dbo2.Product{Name: "Café", Price: 3.90, CategoryID: 3})
	db.Create(&dbo2.Product{Name: "Cappuccino", Price: 4.90, CategoryID: 3})
	db.Create(&dbo2.Product{Name: "Chocolate Quente", Price: 4.90, CategoryID: 3})
	db.Create(&dbo2.Product{Name: "Misto Quente + Fritas", Price: 9.90, CategoryID: 4})
	db.Create(&dbo2.Product{Name: "X-Burguer + Fritas + Coca-Cola", Price: 29.90, CategoryID: 4})
	db.Create(&dbo2.Product{Name: "X-Salada + Fritas + Coca-Cola", Price: 31.90, CategoryID: 4})
	db.Create(&dbo2.Product{Name: "X-Bacon + Fritas + Coca-Cola", Price: 33.90, CategoryID: 4})
	db.Create(&dbo2.Product{Name: "X-Tudo + Fritas + Coca-Cola", Price: 37.90, CategoryID: 4})
	db.Create(&dbo2.Product{Name: "Sorvete", Price: 7.90, CategoryID: 5})
	db.Create(&dbo2.Product{Name: "Sundae", Price: 9.90, CategoryID: 5})
	db.Create(&dbo2.Product{Name: "Açaí", Price: 11.90, CategoryID: 5})
	db.Create(&dbo2.Product{Name: "Batata Frita", Price: 9.90, CategoryID: 6})
	db.Create(&dbo2.Product{Name: "Batata Frita com Cheddar", Price: 11.90, CategoryID: 6})
	db.Create(&dbo2.Product{Name: "Batata Frita com Cheddar e Bacon", Price: 13.90, CategoryID: 6})
	db.Create(&dbo2.Product{Name: "Batata Frita com Cheddar e Calabresa", Price: 13.90, CategoryID: 6})
	db.Create(&dbo2.Product{Name: "Batata Frita com Cheddar e Frango", Price: 13.90, CategoryID: 6})

	return db, nil
}
