package repositories

import (
	"fashion-backend/internal/models"

	"gorm.io/gorm"
)

type PromosiRepository interface {
	GetAll() ([]models.Promosi, error)
	GetByID(id uint) (*models.Promosi, error)
	Create(promosi *models.Promosi) error
	Update(promosi *models.Promosi) error
	Delete(id uint) error
}

type promosiRepo struct {
	db *gorm.DB
}

func NewPromosiRepository(db *gorm.DB) PromosiRepository {
	return &promosiRepo{db: db}
}

func (r *promosiRepo) GetAll() ([]models.Promosi, error) {
	var promosi []models.Promosi
	err := r.db.Find(&promosi).Error
	return promosi, err
}

func (r *promosiRepo) GetByID(id uint) (*models.Promosi, error) {
	var promosi models.Promosi
	err := r.db.First(&promosi, id).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return &promosi, err
}

func (r *promosiRepo) Create(promosi *models.Promosi) error {
	return r.db.Create(promosi).Error
}

func (r *promosiRepo) Update(promosi *models.Promosi) error {
	return r.db.Save(promosi).Error
}

func (r *promosiRepo) Delete(id uint) error {
	return r.db.Delete(&models.Promosi{}, id).Error
}
