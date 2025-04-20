package services

import (
	"log"
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

func NewTourService(repository repositories.TourRepository) *tourService {
	return &tourService{repository: repository}
}

func (s *tourService) GetAllTours() ([]models.Tour, error) {
	tours, err := s.repository.FindAll()
	if err != nil {
		log.Printf("Error fetching all tours: %v", err)
		return nil, err
	}
	return tours, nil
}

func (s *tourService) GetTourByID(id uint) (models.Tour, error) {
	tour, err := s.repository.FindByID(id)
	if err != nil {
		return models.Tour{}, err
	}
	return *tour, nil
}

func (s *tourService) CreateTour(tour *models.Tour) error {
	err := s.repository.Create(tour)
	if err != nil {
		log.Printf("Error creating tour: %v", err)
		return err
	}
	return nil
}

func (s *tourService) UpdateTour(tour *models.Tour) error {
	err := s.repository.Update(tour)
	if err != nil {
		log.Printf("Error updating tour (ID: %d): %v", tour.ID, err)
		return err
	}
	return nil
}

func (s *tourService) DeleteTour(id uint) error {
	err := s.repository.Delete(id)
	if err != nil {
		log.Printf("Error deleting tour (ID: %d): %v", id, err)
		return err
	}
	return nil
}
