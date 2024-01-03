package repository

import (
	"github.com/fiap/challenge-gofood/internal/core/domain"
	"gorm.io/gorm"
)

type Attendant struct {
	gorm.Model
	Name string `gorm:"unique"`
}

type AttendantRepository struct {
	db *gorm.DB
}

func NewAttendantRepository(db *gorm.DB) *AttendantRepository {
	return &AttendantRepository{
		db,
	}
}

func (c *AttendantRepository) CreateAttendant(name string) (*domain.Attendant, error) {
	attendant := &Attendant{
		Name: name,
	}

	var err error
	if err = c.db.Create(attendant).Error; err != nil {
		return nil, err
	}

	var result Attendant
	c.db.Where("name = ?", name).First(&result)

	return &domain.Attendant{
		ID:        result.ID,
		Name:      result.Name,
		CreatedAt: result.CreatedAt,
		UpdatedAt: result.UpdatedAt,
	}, nil
}

func (c *AttendantRepository) GetAttendantById(id uint) (*domain.Attendant, error) {
	var result Attendant
	if err := c.db.First(&result, id).Error; err != nil {
		return nil, err
	}

	return &domain.Attendant{
		ID:        result.ID,
		Name:      result.Name,
		CreatedAt: result.CreatedAt,
		UpdatedAt: result.UpdatedAt,
	}, nil
}

func (c *AttendantRepository) GetAttendants() ([]*domain.Attendant, error) {
	var results []*Attendant
	if err := c.db.Find(&results).Error; err != nil {
		return nil, err
	}

	var attendants []*domain.Attendant
	for _, result := range results {
		attendants = append(attendants, &domain.Attendant{
			ID:        result.ID,
			Name:      result.Name,
			CreatedAt: result.CreatedAt,
			UpdatedAt: result.UpdatedAt,
		})
	}

	return attendants, nil
}

func (c *AttendantRepository) UpdateAttendant(attendant *domain.Attendant) (*domain.Attendant, error) {
	var result Attendant
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
	if err := c.db.Delete(&Attendant{}, id).Error; err != nil {
		return err
	}

	return nil
}
