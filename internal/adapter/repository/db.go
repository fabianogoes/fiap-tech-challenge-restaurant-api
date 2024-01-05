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

	db.Create(&Attendant{Name: "Miguel"})
	db.Create(&Attendant{Name: "Sophia"})
	db.Create(&Attendant{Name: "Alice"})
	db.Create(&Attendant{Name: "Pedro"})
	db.Create(&Attendant{Name: "Manuela"})

	db.Create(&Customer{Name: "Bernardo", Email: "bernardo@gmail.com", CPF: "29381510040"})
	db.Create(&Customer{Name: "Laura", Email: "laura@hotmail.com", CPF: "15204180001"})
	db.Create(&Customer{Name: "Lucas", Email: "lucas@gmail.com", CPF: "43300921074"})
	db.Create(&Customer{Name: "Maria Eduarda", Email: "meduarda@uol.com.br", CPF: "85752055016"})
	db.Create(&Customer{Name: "Guilherme", Email: "guilherme@microsoft.com", CPF: "17148604001"})

	db.Create(&Category{Name: "Sanduíches"})
	db.Create(&Category{Name: "Bebidas Frias"})
	db.Create(&Category{Name: "Bebidas Quentes"})
	db.Create(&Category{Name: "Combos"})
	db.Create(&Category{Name: "Sobremesas"})
	db.Create(&Category{Name: "Acompanhamentos"})
	db.Create(&Category{Name: "Café da Manhã"})

	db.Create(&Product{Name: "Big Lanche", Price: 29.90, CategoryID: 1})
	db.Create(&Product{Name: "X-Burguer", Price: 19.90, CategoryID: 1})
	db.Create(&Product{Name: "X-Salada", Price: 21.90, CategoryID: 1})
	db.Create(&Product{Name: "X-Bacon", Price: 23.90, CategoryID: 1})
	db.Create(&Product{Name: "X-Tudo", Price: 27.90, CategoryID: 1})
	db.Create(&Product{Name: "Coca-Cola", Price: 5.90, CategoryID: 2})
	db.Create(&Product{Name: "Guaraná", Price: 5.90, CategoryID: 2})
	db.Create(&Product{Name: "Fanta", Price: 5.90, CategoryID: 2})
	db.Create(&Product{Name: "Suco de Laranja", Price: 5.90, CategoryID: 2})
	db.Create(&Product{Name: "Suco de Uva", Price: 5.90, CategoryID: 2})
	db.Create(&Product{Name: "Café", Price: 3.90, CategoryID: 3})
	db.Create(&Product{Name: "Cappuccino", Price: 4.90, CategoryID: 3})
	db.Create(&Product{Name: "Chocolate Quente", Price: 4.90, CategoryID: 3})
	db.Create(&Product{Name: "Misto Quente + Fritas", Price: 9.90, CategoryID: 4})
	db.Create(&Product{Name: "X-Burguer + Fritas + Coca-Cola", Price: 29.90, CategoryID: 4})
	db.Create(&Product{Name: "X-Salada + Fritas + Coca-Cola", Price: 31.90, CategoryID: 4})
	db.Create(&Product{Name: "X-Bacon + Fritas + Coca-Cola", Price: 33.90, CategoryID: 4})
	db.Create(&Product{Name: "X-Tudo + Fritas + Coca-Cola", Price: 37.90, CategoryID: 4})
	db.Create(&Product{Name: "Sorvete", Price: 7.90, CategoryID: 5})
	db.Create(&Product{Name: "Sundae", Price: 9.90, CategoryID: 5})
	db.Create(&Product{Name: "Açaí", Price: 11.90, CategoryID: 5})
	db.Create(&Product{Name: "Batata Frita", Price: 9.90, CategoryID: 6})
	db.Create(&Product{Name: "Batata Frita com Cheddar", Price: 11.90, CategoryID: 6})
	db.Create(&Product{Name: "Batata Frita com Cheddar e Bacon", Price: 13.90, CategoryID: 6})
	db.Create(&Product{Name: "Batata Frita com Cheddar e Calabresa", Price: 13.90, CategoryID: 6})
	db.Create(&Product{Name: "Batata Frita com Cheddar e Frango", Price: 13.90, CategoryID: 6})

	return db, nil
}
