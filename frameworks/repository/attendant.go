package repository

import (
	"fmt"
	"github.com/fabianogoes/fiap-challenge/domain/entities"
	"github.com/fabianogoes/fiap-challenge/frameworks/repository/dbo"
	"github.com/fabianogoes/fiap-challenge/shared"
	"gorm.io/gorm"
	"log"
)

type AttendantRepository struct {
	db     *gorm.DB
	crypto *shared.Crypto
}

func NewAttendantRepository(db *gorm.DB, crypto *shared.Crypto) *AttendantRepository {
	return &AttendantRepository{
		db,
		crypto,
	}
}

func (c *AttendantRepository) CreateAttendant(name string) (*entities.Attendant, error) {
	attendant := &dbo.Attendant{
		Name: c.crypto.EncryptAES(name),
	}

	var err error
	if err = c.db.Create(attendant).Error; err != nil {
		return nil, err
	}

	return c.GetAttendantByName(name)
}

func (c *AttendantRepository) GetAttendantByName(name string) (*entities.Attendant, error) {
	var result dbo.Attendant
	if err := c.db.Where("name = ?", c.crypto.EncryptAES(name)).First(&result).Error; err != nil {
		return nil, fmt.Errorf("error to find attendant with name %s - %v", name, err)
	}

	return result.ToEntity(c.crypto), nil
}

func (c *AttendantRepository) GetAttendantById(id uint) (*entities.Attendant, error) {
	var result dbo.Attendant
	if err := c.db.First(&result, id).Error; err != nil {
		return nil, fmt.Errorf("error to find attendant with id %d - %v", id, err)
	}

	return result.ToEntity(c.crypto), nil
}

func (c *AttendantRepository) GetAttendants() ([]*entities.Attendant, error) {
	var results []*dbo.Attendant
	if err := c.db.Find(&results).Error; err != nil {
		return nil, err
	}

	var attendants []*entities.Attendant
	for _, result := range results {
		attendants = append(attendants, result.ToEntity(c.crypto))
	}

	return attendants, nil
}

func (c *AttendantRepository) UpdateAttendant(attendant *entities.Attendant) (*entities.Attendant, error) {
	var result dbo.Attendant
	if err := c.db.First(&result, attendant.ID).Error; err != nil {
		return nil, err
	}

	result.Name = c.crypto.EncryptAES(attendant.Name)

	if err := c.db.Save(&result).Error; err != nil {
		return nil, err
	}

	return c.GetAttendantById(attendant.ID)
}

func (c *AttendantRepository) DeleteAttendant(id uint) error {
	if err := c.db.Delete(&dbo.Attendant{}, id).Error; err != nil {
		return err
	}

	return nil
}

func InitialDataAttendants(db *gorm.DB, crypto *shared.Crypto) {
	if count := db.Find(&[]*dbo.Attendant{}).RowsAffected; count == 0 {
		log.Print("Inserting Attendants...")
		db.Create(&dbo.Attendant{Name: crypto.EncryptAES("Miguel")})
		db.Create(&dbo.Attendant{Name: crypto.EncryptAES("Sophia")})
		db.Create(&dbo.Attendant{Name: crypto.EncryptAES("Alice")})
		db.Create(&dbo.Attendant{Name: crypto.EncryptAES("Pedro")})
		db.Create(&dbo.Attendant{Name: crypto.EncryptAES("Manuela")})
	}
}
