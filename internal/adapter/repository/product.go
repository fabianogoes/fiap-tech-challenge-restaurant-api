package repository

import (
	"fmt"

	"github.com/fiap/challenge-gofood/internal/adapter/repository/dbo"
	"github.com/fiap/challenge-gofood/internal/core/domain"
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

func (p *ProductRepository) CreateProduct(name string, price float64, categoryID uint) (*domain.Product, error) {
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

func (p *ProductRepository) GetProductByName(name string) (*domain.Product, error) {
	var result dbo.Product
	if err := p.db.Where("name = ?", name).First(&result).Error; err != nil {
		return nil, fmt.Errorf("error to find product with name %s - %v", name, err)
	}

	return result.ToEntity(), nil
}

func (p *ProductRepository) GetProductById(id uint) (*domain.Product, error) {
	var result dbo.Product
	if err := p.db.Model(&dbo.Product{}).Preload("Category").First(&result, id).Error; err != nil {
		return nil, fmt.Errorf("error to find product with id %d - %v", id, err)
	}

	return result.ToEntity(), nil
}

func (p *ProductRepository) GetProducts() ([]*domain.Product, error) {
	var results []*dbo.Product
	if err := p.db.Model(&dbo.Product{}).Preload("Category").Find(&results).Error; err != nil {
		return nil, err
	}

	var products []*domain.Product
	for _, result := range results {
		products = append(products, result.ToEntity())
	}

	return products, nil
}

func (p *ProductRepository) UpdateProduct(product *domain.Product) (*domain.Product, error) {
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
