package repository

import (
	"fmt"
	"github.com/fabianogoes/fiap-challenge/domain/entities"
	"github.com/fabianogoes/fiap-challenge/frameworks/repository/dbo"
	"log"

	"gorm.io/gorm"
)

type ProductRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{
		db,
	}
}

func (p *ProductRepository) CreateProduct(name string, price float64, categoryID uint) (*entities.Product, error) {
	var err error
	product := &dbo.Product{
		Name:       name,
		Price:      price,
		CategoryID: categoryID,
	}

	if err = p.db.Create(product).Error; err != nil {
		return nil, err
	}

	return p.GetProductByName(name)
}

func (p *ProductRepository) GetProductByName(name string) (*entities.Product, error) {
	var result dbo.Product
	if err := p.db.Where("name = ?", name).First(&result).Error; err != nil {
		return nil, fmt.Errorf("error to find product with name %s - %v", name, err)
	}

	return result.ToEntity(), nil
}

func (p *ProductRepository) GetProductById(id uint) (*entities.Product, error) {
	var result dbo.Product
	if err := p.db.Model(&dbo.Product{}).Preload("Category").First(&result, id).Error; err != nil {
		return nil, fmt.Errorf("error to find product with id %d - %v", id, err)
	}

	return result.ToEntity(), nil
}

func (p *ProductRepository) GetProducts() ([]*entities.Product, error) {
	var results []*dbo.Product
	if err := p.db.Model(&dbo.Product{}).Preload("Category").Find(&results).Error; err != nil {
		return nil, err
	}

	var products []*entities.Product
	for _, result := range results {
		products = append(products, result.ToEntity())
	}

	return products, nil
}

func (p *ProductRepository) UpdateProduct(product *entities.Product) (*entities.Product, error) {
	var result dbo.Product
	if err := p.db.First(&result, product.ID).Error; err != nil {
		return nil, err
	}

	result.Name = product.Name
	result.Price = product.Price
	result.Category = dbo.Category{
		Name: product.Category.Name,
	}

	if err := p.db.Save(&result).Error; err != nil {
		return nil, err
	}

	return p.GetProductById(product.ID)
}

func (p *ProductRepository) DeleteProduct(id uint) error {
	if err := p.db.Delete(&dbo.Product{}, id).Error; err != nil {
		return err
	}

	return nil
}

func InitialDataProducts(db *gorm.DB) {
	if count := db.Find(&[]*dbo.Product{}).RowsAffected; count == 0 {
		log.Print("Inserting Products...")
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
	}
}
