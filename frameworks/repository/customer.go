package repository

import (
	"fmt"
	"github.com/fabianogoes/fiap-challenge/domain/entities"
	"github.com/fabianogoes/fiap-challenge/frameworks/repository/dbo"
	"github.com/fabianogoes/fiap-challenge/shared"
	"log"

	"gorm.io/gorm"
)

type CustomerRepository struct {
	db     *gorm.DB
	crypto *shared.Crypto
}

func NewCustomerRepository(db *gorm.DB, crypto *shared.Crypto) *CustomerRepository {
	return &CustomerRepository{
		db,
		crypto,
	}
}

func (c *CustomerRepository) CreateCustomer(customer *entities.Customer) (*entities.Customer, error) {
	var err error

	customer.Name = c.crypto.EncryptAES(customer.Name)
	customer.Email = c.crypto.EncryptAES(customer.Email)
	if err = c.db.Create(customer).Error; err != nil {
		return nil, err
	}

	return c.GetCustomerByCPF(customer.CPF)
}

func (c *CustomerRepository) GetCustomerByCPF(cpf string) (*entities.Customer, error) {
	var result dbo.Customer

	if err := c.db.Where("cpf = ?", cpf).First(&result).Error; err != nil {
		return nil, fmt.Errorf("error to find customer with cpf %s - %v", cpf, err)
	}

	return result.ToEntity(c.crypto), nil
}

func (c *CustomerRepository) GetCustomerById(id uint) (*entities.Customer, error) {
	var result dbo.Customer
	if err := c.db.First(&result, id).Error; err != nil {
		return nil, fmt.Errorf("error to find customer with id %d - %v", id, err)
	}

	return result.ToEntity(c.crypto), nil
}

func (c *CustomerRepository) GetCustomers() ([]*entities.Customer, error) {
	var results []*dbo.Customer
	if err := c.db.Find(&results).Error; err != nil {
		return nil, err
	}

	var customers []*entities.Customer
	for _, result := range results {
		customers = append(customers, result.ToEntity(c.crypto))
	}

	return customers, nil
}

func (c *CustomerRepository) UpdateCustomer(customer *entities.Customer) (*entities.Customer, error) {
	var result dbo.Customer
	if err := c.db.First(&result, customer.ID).Error; err != nil {
		return nil, err
	}

	result.Name = c.crypto.EncryptAES(customer.Name)
	result.Email = c.crypto.EncryptAES(customer.Email)
	result.CPF = c.crypto.EncryptAES(customer.CPF)

	if err := c.db.Save(&result).Error; err != nil {
		return nil, err
	}

	return c.GetCustomerById(customer.ID)
}

func (c *CustomerRepository) DeleteCustomer(id uint) error {
	if err := c.db.Delete(&dbo.Customer{}, id).Error; err != nil {
		return err
	}

	return nil
}

func InitialDataCustomers(db *gorm.DB, crypto *shared.Crypto) {
	if count := db.Find(&[]*dbo.Customer{}).RowsAffected; count == 0 {
		log.Print("Inserting Customers...")
		db.Create(&dbo.Customer{
			Name:  crypto.EncryptAES("Bernardo"),
			Email: crypto.EncryptAES("bernardo@gmail.com"),
			CPF:   "29381510040",
		})
		db.Create(&dbo.Customer{
			Name:  crypto.EncryptAES("Laura"),
			Email: crypto.EncryptAES("laura@hotmail.com"),
			CPF:   "15204180001",
		})
		db.Create(&dbo.Customer{
			Name:  crypto.EncryptAES("Lucas"),
			Email: crypto.EncryptAES("lucas@gmail.com"),
			CPF:   "43300921074",
		})
		db.Create(&dbo.Customer{
			Name:  crypto.EncryptAES("Maria Eduarda"),
			Email: crypto.EncryptAES("meduarda@uol.com.br"),
			CPF:   "85752055016",
		})
		db.Create(&dbo.Customer{
			Name:  crypto.EncryptAES("Guilherme"),
			Email: crypto.EncryptAES("guilherme@microsoft.com"),
			CPF:   "17148604001",
		})
	}
}
