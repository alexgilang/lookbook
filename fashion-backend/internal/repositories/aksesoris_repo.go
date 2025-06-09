package repositories

import (
	"fashion-backend/internal/models"

	"gorm.io/gorm"
)

type AksesorisRepository interface {
	GetAll() ([]models.Aksesoris, error)
	GetByID(id uint) (*models.Aksesoris, error)
	Create(aksesoris *models.Aksesoris) error
	Update(aksesoris *models.Aksesoris) error
	Delete(id uint) error
}

type aksesorisRepo struct {
	db *gorm.DB
}

func NewAksesorisRepository(db *gorm.DB) AksesorisRepository {
	return &aksesorisRepo{db: db}
}

func (r *aksesorisRepo) GetAll() ([]models.Aksesoris, error) {
	var aksesoris []models.Aksesoris
	err := r.db.Find(&aksesoris).Error
	return aksesoris, err
}

func (r *aksesorisRepo) GetByID(id uint) (*models.Aksesoris, error) {
	var aksesoris models.Aksesoris
	err := r.db.First(&aksesoris, id).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return &aksesoris, err
}

func (r *aksesorisRepo) Create(aksesoris *models.Aksesoris) error {
	return r.db.Create(aksesoris).Error
}

func (r *aksesorisRepo) Update(aksesoris *models.Aksesoris) error {
	return r.db.Save(aksesoris).Error
}

func (r *aksesorisRepo) Delete(id uint) error {
	return r.db.Delete(&models.Aksesoris{}, id).Error
}
