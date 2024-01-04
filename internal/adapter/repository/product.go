package repository

import (
	"github.com/fiap/challenge-gofood/internal/core/domain"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name       string `gorm:"unique"`
	Price      float64
	Quantity   int
	CategoryID int
	Category   Category `gorm:"ForeignKey:CategoryID"`
}

func (p *Product) ToModel() *domain.Product {
	return &domain.Product{
		ID:       p.ID,
		Name:     p.Name,
		Price:    p.Price,
		Quantity: p.Quantity,
		Category: &domain.Category{
			ID:   p.Category.ID,
			Name: p.Category.Name,
		},
		CreatedAt: p.CreatedAt,
		UpdatedAt: p.UpdatedAt,
	}
}

type Category struct {
	gorm.Model
	Name string
}

func (c *Category) ToModel() *domain.Category {
	return &domain.Category{
		ID:   c.ID,
		Name: c.Name,
	}
}

type ProductRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{
		db,
	}
}

func (p *ProductRepository) CreateProduct(name string, price float64, quantity int, categoryID int) (*domain.Product, error) {
	var err error
	product := &Product{
		Name:       name,
		Price:      price,
		Quantity:   quantity,
		CategoryID: categoryID,
	}

	if err = p.db.Create(product).Error; err != nil {
		return nil, err
	}

	var result Product
	p.db.Where("name = ?", product.Name).First(&result)

	return result.ToModel(), nil
}

func (p *ProductRepository) GetProductById(id uint) (*domain.Product, error) {
	var result Product
	if err := p.db.Model(&Product{}).Preload("Category").First(&result, id).Error; err != nil {
		return nil, err
	}

	return result.ToModel(), nil
}

func (p *ProductRepository) GetProducts() ([]*domain.Product, error) {
	var results []*Product
	if err := p.db.Model(&Product{}).Preload("Category").Find(&results).Error; err != nil {
		return nil, err
	}

	var products []*domain.Product
	for _, result := range results {
		products = append(products, result.ToModel())
	}

	return products, nil
}

func (p *ProductRepository) UpdateProduct(product *domain.Product) (*domain.Product, error) {
	var result Product
	if err := p.db.First(&result, product.ID).Error; err != nil {
		return nil, err
	}

	result.Name = product.Name
	result.Price = product.Price
	result.Quantity = product.Quantity
	result.Category = Category{
		Name: product.Category.Name,
	}

	if err := p.db.Save(&result).Error; err != nil {
		return nil, err
	}

	return p.GetProductById(product.ID)
}

func (p *ProductRepository) DeleteProduct(id uint) error {
	if err := p.db.Delete(&Product{}, id).Error; err != nil {
		return err
	}

	return nil
}
