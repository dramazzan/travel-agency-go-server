package repositories

import (
	"gorm.io/gorm"
	"server-go/internal/models"
)

type TourRepository interface {
	FindAll() ([]models.Tour, error)
	FindByID(id uint) (models.Tour, error)
	Create(tour *models.Tour) error
	Update(tour *models.Tour) error
	Delete(id uint) error
}

type tourRepository struct {
	db *gorm.DB
}

func NewTourRepository(db *gorm.DB) TourRepository {
	return &tourRepository{db: db}
}

func (r *tourRepository) FindAll() ([]models.Tour, error) {
	var tours []models.Tour
	result := r.db.Find(&tours)
	return tours, result.Error
}

func (r *tourRepository) FindByID(id uint) (models.Tour, error) {
	var tour models.Tour
	result := r.db.First(&tour, id)
	return tour, result.Error
}

func (r *tourRepository) Create(tour *models.Tour) error {
	return r.db.Create(tour).Error
}

func (r *tourRepository) Update(tour *models.Tour) error {
	return r.db.Save(tour).Error
}

func (r *tourRepository) Delete(id uint) error {
	return r.db.Delete(&models.Tour{}, id).Error
}
