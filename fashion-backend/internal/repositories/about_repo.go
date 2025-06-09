package repositories

import (
	"fashion-backend/internal/models"

	"gorm.io/gorm"
)

type AboutRepository interface {
	GetAbout() (*models.About, error)
	CreateAbout(about *models.About) error
	UpdateAbout(about *models.About) error
	DeleteAbout(id uint) error
}

type aboutRepo struct {
	db *gorm.DB
}

func NewAboutRepository(db *gorm.DB) AboutRepository {
	return &aboutRepo{db: db}
}

func (r *aboutRepo) GetAbout() (*models.About, error) {
	var about models.About
	err := r.db.First(&about).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return &about, err
}

func (r *aboutRepo) CreateAbout(about *models.About) error {
	return r.db.Create(about).Error
}

func (r *aboutRepo) UpdateAbout(about *models.About) error {
	return r.db.Save(about).Error
}

func (r *aboutRepo) DeleteAbout(id uint) error {
	return r.db.Delete(&models.About{}, id).Error
}
