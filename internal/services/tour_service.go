package services

import (
	"server-go/internal/models"
	"server-go/internal/repositories"
)

type TourService interface {
	GetAllTours() ([]models.Tour, error)
	GetTourByID(id uint) (models.Tour, error)
	CreateTour(tour *models.Tour) error
	UpdateTour(tour *models.Tour) error
	DeleteTour(id uint) error
}

type tourService struct {
	repository repositories.TourRepository
}

func NewTourService(repository repositories.TourRepository) TourService {
	return &tourService{repository: repository}
}

func (s *tourService) GetAllTours() ([]models.Tour, error) {
	return s.repository.FindAll()
}

func (s *tourService) GetTourByID(id uint) (models.Tour, error) {
	return s.repository.FindByID(id)
}

func (s *tourService) CreateTour(tour *models.Tour) error {
	return s.repository.Create(tour)
}

func (s *tourService) UpdateTour(tour *models.Tour) error {
	return s.repository.Update(tour)
}

func (s *tourService) DeleteTour(id uint) error {
	return s.repository.Delete(id)
}
