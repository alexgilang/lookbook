package repositories

import (
	"fashion-backend/internal/models"

	"gorm.io/gorm"
)

type DesainRepository interface {
	GetAll() ([]models.Desain, error)
	GetByID(id uint) (*models.Desain, error)
	Create(desain *models.Desain) error
	Update(desain *models.Desain) error
	Delete(id uint) error
}

type desainRepo struct {
	db *gorm.DB
}

func NewDesainRepository(db *gorm.DB) DesainRepository {
	return &desainRepo{db: db}
}

func (r *desainRepo) GetAll() ([]models.Desain, error) {
	var desains []models.Desain
	err := r.db.Find(&desains).Error
	return desains, err
}

func (r *desainRepo) GetByID(id uint) (*models.Desain, error) {
	var desain models.Desain
	err := r.db.First(&desain, id).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return &desain, err
}

func (r *desainRepo) Create(desain *models.Desain) error {
	return r.db.Create(desain).Error
}

func (r *desainRepo) Update(desain *models.Desain) error {
	return r.db.Save(desain).Error
}

func (r *desainRepo) Delete(id uint) error {
	return r.db.Delete(&models.Desain{}, id).Error
}
