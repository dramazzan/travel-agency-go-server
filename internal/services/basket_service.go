package services

import (
	"errors"
	"server-go/internal/models"
	"server-go/internal/repositories"
)

type BasketService interface {
	CreateBasket(userID uint) (models.Basket, error)
	GetBasketByUserID(userID uint) (models.Basket, error)
	AddTourToBasket(basketID, tourID uint) error
	RemoveTourFromBasket(basketID, tourID uint) error
	DeleteBasket(basketID uint) error
}

type basketService struct {
	repo repositories.BasketRepository
}

func NewBasketService(repo repositories.BasketRepository) BasketService {
	return &basketService{repo: repo}
}

func (s *basketService) CreateBasket(userID uint) (models.Basket, error) {
	if basket, err := s.repo.FindByUserID(userID); err == nil && basket.ID != 0 {
		return models.Basket{}, errors.New("basket already exists")
	}

	basket := models.Basket{UserID: userID}
	if err := s.repo.Create(&basket); err != nil {
		return models.Basket{}, err
	}
	return basket, nil
}

func (s *basketService) GetBasketByUserID(userID uint) (models.Basket, error) {
	return s.repo.FindByUserID(userID)
}

func (s *basketService) AddTourToBasket(basketID, tourID uint) error {
	return s.repo.AddTour(basketID, tourID)
}

func (s *basketService) RemoveTourFromBasket(basketID, tourID uint) error {
	return s.repo.RemoveTour(basketID, tourID)
}

func (s *basketService) DeleteBasket(basketID uint) error {
	return s.repo.Delete(basketID)
}
