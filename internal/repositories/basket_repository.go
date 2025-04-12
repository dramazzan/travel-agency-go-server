package repositories

import (
	"gorm.io/gorm"
	"server-go/internal/models"
)

type BasketRepository interface {
	Create(basket *models.Basket) error
	FindByUserID(userID uint) (models.Basket, error)
	AddTour(basketID, tourID uint) error
	RemoveTour(basketID, tourID uint) error
	Delete(basketID uint) error
}

type basketRepository struct {
	db *gorm.DB
}

func NewBasketRepository(db *gorm.DB) BasketRepository {
	return &basketRepository{db: db}
}

func (r *basketRepository) Create(basket *models.Basket) error {
	return r.db.Create(basket).Error
}

func (r *basketRepository) FindByUserID(userID uint) (models.Basket, error) {
	var basket models.Basket
	err := r.db.Preload("Tours").Where("user_id = ?", userID).First(&basket).Error
	return basket, err
}

func (r *basketRepository) AddTour(basketID, tourID uint) error {
	return r.db.Exec("INSERT INTO basket_tours (basket_id, tour_id) VALUES (?, ?)", basketID, tourID).Error
}

func (r *basketRepository) RemoveTour(basketID, tourID uint) error {
	return r.db.Exec("DELETE FROM basket_tours WHERE basket_id = ? AND tour_id = ?", basketID, tourID).Error
}

func (r *basketRepository) Delete(basketID uint) error {
	return r.db.Delete(&models.Basket{}, basketID).Error
}
