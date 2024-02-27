package repository

import (
	"fmt"
	"github.com/fabianogoes/fiap-challenge/entities"

	"github.com/fabianogoes/fiap-challenge/adapters/repository/dbo"
	"gorm.io/gorm"
)

type AttendantRepository struct {
	db *gorm.DB
}

func NewAttendantRepository(db *gorm.DB) *AttendantRepository {
	return &AttendantRepository{
		db,
	}
}

func (c *AttendantRepository) CreateAttendant(name string) (*entities.Attendant, error) {
	attendant := &dbo.Attendant{
		Name: name,
	}

	var err error
	if err = c.db.Create(attendant).Error; err != nil {
		return nil, err
	}

	return c.GetAttendantByName(name)
}

func (c *AttendantRepository) GetAttendantByName(name string) (*entities.Attendant, error) {
	var result dbo.Attendant
	if err := c.db.Where("name = ?", name).First(&result).Error; err != nil {
		return nil, fmt.Errorf("error to find attendant with name %s - %v", name, err)
	}

	return result.ToEntity(), nil
}

func (c *AttendantRepository) GetAttendantById(id uint) (*entities.Attendant, error) {
	var result dbo.Attendant
	if err := c.db.First(&result, id).Error; err != nil {
		return nil, fmt.Errorf("error to find attendant with id %d - %v", id, err)
	}

	return result.ToEntity(), nil
}

func (c *AttendantRepository) GetAttendants() ([]*entities.Attendant, error) {
	var results []*dbo.Attendant
	if err := c.db.Find(&results).Error; err != nil {
		return nil, err
	}

	var attendants []*entities.Attendant
	for _, result := range results {
		attendants = append(attendants, result.ToEntity())
	}

	return attendants, nil
}

func (c *AttendantRepository) UpdateAttendant(attendant *entities.Attendant) (*entities.Attendant, error) {
	var result dbo.Attendant
	if err := c.db.First(&result, attendant.ID).Error; err != nil {
		return nil, err
	}

	result.Name = attendant.Name

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
