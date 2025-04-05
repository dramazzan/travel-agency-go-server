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

// NewAuthService returns an implementation of the AuthService interface
func NewAuthService(repository repositories.AuthRepository) AuthService {
	return &authService{repository: repository}
}

// Register creates a new user account
func (s *authService) Register(username, email, password string) error {
	// Check if user with this email already exists
	existingUser, err := s.repository.GetUserByEmail(email)
	if err != nil {
		// Only return the error if it's not a "record not found" error
		if err.Error() != "record not found" {
			return fmt.Errorf("error checking existing user: %w", err)
		}
		// If it's "record not found", that's what we want - continue with registration
	} else if existingUser != nil {
		return fmt.Errorf("user with email %s already exists", email)
	}

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("error hashing password: %w", err)
	}

	// Create new user
	user := &models.User{
		Username: username,
		Email:    email,
		Password: string(hashedPassword),
	}

	// Save user to repository
	return s.repository.CreateUser(user)
}

// Login authenticates a user and returns a JWT token
func (s *authService) Login(email, password string) (string, error) {
	user, err := s.repository.GetUserByEmail(email)
	if err != nil {
		return "", fmt.Errorf("invalid email or password")
	}

	// Verify password
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", fmt.Errorf("invalid email or password")
	}

	// Generate token
	token, err := auth.GenerateToken(user.Username)
	if err != nil {
		return "", fmt.Errorf("failed to generate token: %w", err)
	}

	return token, nil
}
