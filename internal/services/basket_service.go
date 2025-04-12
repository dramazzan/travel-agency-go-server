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
	repository repositories.BasketRepository
}

func NewBasketService(repository repositories.BasketRepository) BasketService {
	return &basketService{repository: repository}
}

func (s *basketService) CreateBasket(userID uint) (models.Basket, error) {
	existingBasket, _ := s.repository.FindByUserID(userID)
	if existingBasket.ID != 0 {
		return models.Basket{}, errors.New("basket already exists")
	}

	basket := models.Basket{
		UserID: userID,
	}

	err := s.repository.Create(&basket)
	if err != nil {
		return models.Basket{}, err
	}

	return basket, nil
}

func (s *basketService) GetBasketByUserID(userID uint) (models.Basket, error) {
	basket, err := s.repository.FindByUserID(userID)
	if err != nil {
		return models.Basket{}, err
	}
	return basket, nil
}

func (s *basketService) AddTourToBasket(basketID, tourID uint) error {
	return s.repository.AddTour(basketID, tourID)
}

func (s *basketService) RemoveTourFromBasket(basketID, tourID uint) error {
	return s.repository.RemoveTour(basketID, tourID)
}

func (s *basketService) DeleteBasket(basketID uint) error {
	return s.repository.Delete(basketID)
}
