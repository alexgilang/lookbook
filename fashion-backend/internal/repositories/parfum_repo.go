package repositories

import (
	"fashion-backend/internal/models"

	"gorm.io/gorm"
)

type ParfumRepository interface {
	GetAll() ([]models.Parfum, error)
	GetByID(id uint) (*models.Parfum, error)
	Create(parfum *models.Parfum) error
	Update(parfum *models.Parfum) error
	Delete(id uint) error
}

type parfumRepo struct {
	db *gorm.DB
}

func NewParfumRepository(db *gorm.DB) ParfumRepository {
	return &parfumRepo{db: db}
}

func (r *parfumRepo) GetAll() ([]models.Parfum, error) {
	var parfums []models.Parfum
	err := r.db.Find(&parfums).Error
	return parfums, err
}

func (r *parfumRepo) GetByID(id uint) (*models.Parfum, error) {
	var parfum models.Parfum
	err := r.db.First(&parfum, id).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return &parfum, err
}

func (r *parfumRepo) Create(parfum *models.Parfum) error {
	return r.db.Create(parfum).Error
}

func (r *parfumRepo) Update(parfum *models.Parfum) error {
	return r.db.Save(parfum).Error
}

func (r *parfumRepo) Delete(id uint) error {
	return r.db.Delete(&models.Parfum{}, id).Error
}
