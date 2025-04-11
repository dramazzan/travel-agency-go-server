package services

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"server-go/internal/auth"
	"server-go/internal/models"
	"server-go/internal/repositories"
)

type AuthService interface {
	Register(username, email, password string) error
	Login(email, password string) (string, error)
}

type authService struct {
	repository repositories.AuthRepository
}

func NewAuthService(repository repositories.AuthRepository) AuthService {
	return &authService{repository: repository}
}

func (s *authService) Register(username, email, password string) error {
	existingUser, err := s.repository.GetUserByEmail(email)
	if err != nil {
		if err.Error() != "record not found" {
			return fmt.Errorf("error checking existing user: %w", err)
		}
	} else if existingUser != nil {
		return fmt.Errorf("user with email %s already exists", email)
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

	return s.repository.CreateUser(user)
}

func (s *authService) Login(email, password string) (string, error) {
	user, err := s.repository.GetUserByEmail(email)
	if err != nil {
		return "", fmt.Errorf("invalid email or password")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", fmt.Errorf("invalid email or password")
	}

	token, err := auth.GenerateToken(user.Username, user.Role)
	if err != nil {
		return "", fmt.Errorf("failed to generate token: %w", err)
	}

	return token, nil
}
