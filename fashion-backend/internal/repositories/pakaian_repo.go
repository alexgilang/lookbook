package repositories

import (
	"fashion-backend/internal/models"

	"gorm.io/gorm"
)

type PakaianRepository interface {
	GetAll() ([]models.Pakaian, error)
	GetByID(id uint) (*models.Pakaian, error)
	Create(pakaian *models.Pakaian) error
	Update(pakaian *models.Pakaian) error
	Delete(id uint) error
}

type pakaianRepo struct {
	db *gorm.DB
}

func NewPakaianRepository(db *gorm.DB) PakaianRepository {
	return &pakaianRepo{db: db}
}

func (r *pakaianRepo) GetAll() ([]models.Pakaian, error) {
	var pakaian []models.Pakaian
	err := r.db.Find(&pakaian).Error
	return pakaian, err
}

func (r *pakaianRepo) GetByID(id uint) (*models.Pakaian, error) {
	var pakaian models.Pakaian
	err := r.db.First(&pakaian, id).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return &pakaian, err
}

func (r *pakaianRepo) Create(pakaian *models.Pakaian) error {
	return r.db.Create(pakaian).Error
}

func (r *pakaianRepo) Update(pakaian *models.Pakaian) error {
	return r.db.Save(pakaian).Error
}

func (r *pakaianRepo) Delete(id uint) error {
	return r.db.Delete(&models.Pakaian{}, id).Error
}
