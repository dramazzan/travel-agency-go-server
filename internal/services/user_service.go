package services

import (
	"fmt"
	"log"

	"golang.org/x/crypto/bcrypt"
	"server-go/internal/auth"
	"server-go/internal/models"
	"server-go/internal/repositories"
)

type AuthService interface {
	Register(username, email, password string) error
	Login(email, password string) (string, error)
	GetUserDataById(id uint) (*models.User, error)
	UpdateUser(user *models.User) error
}

type authService struct {
	repo          repositories.AuthRepository
	basketService BasketService
}

func NewAuthService(repo repositories.AuthRepository, basketService BasketService) AuthService {
	return &authService{repo: repo, basketService: basketService}
}

func (s *authService) Register(username, email, password string) error {
	existingUser, err := s.repo.GetUserByEmail(email)
	if err == nil && existingUser != nil {
		return fmt.Errorf("user with email %s already exists", email)
	}
	if err != nil && err.Error() != "record not found" {
		return fmt.Errorf("error checking existing user: %w", err)
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("error hashing password: %w", err)
	}

	user := &models.User{
		Username: username,
		Email:    email,
		Password: string(hashedPassword),
	}

	if err := s.repo.CreateUser(user); err != nil {
		return fmt.Errorf("error creating user: %w", err)
	}

	if _, err := s.basketService.CreateBasket(user.ID); err != nil {
		log.Printf("error creating basket for user %d: %v", user.ID, err)
		return fmt.Errorf("error creating basket: %w", err)
	}

	return nil
}

func (s *authService) Login(email, password string) (string, error) {
	user, err := s.repo.GetUserByEmail(email)
	if err != nil {
		return "", fmt.Errorf("invalid email or password")
	}

	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)) != nil {
		return "", fmt.Errorf("invalid email or password")
	}

	token, err := auth.GenerateToken(user.ID, user.Username, user.Role, 24)
	if err != nil {
		return "", fmt.Errorf("failed to generate token: %w", err)
	}

	return token, nil
}

func (s *authService) GetUserDataById(id uint) (*models.User, error) {
	user, err := s.repo.GetUserById(id)
	if err != nil {
		return nil, fmt.Errorf("failed to get user by id %d: %w", id, err)
	}
	return user, nil
}

func (s *authService) UpdateUser(user *models.User) error {
	return s.repo.Update(user)
}
