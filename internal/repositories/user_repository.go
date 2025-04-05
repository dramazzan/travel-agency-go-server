package repositories

import (
	"gorm.io/gorm"
	"server-go/internal/models"
)

// AuthRepository interface defines methods for user authentication
type AuthRepository interface {
	GetUserByEmail(email string) (*models.User, error)
	CreateUser(user *models.User) error
}

// authRepositoryImpl implements the AuthRepository interface
type authRepositoryImpl struct {
	DB *gorm.DB
}

// NewAuthRepository creates a new AuthRepository instance
func NewAuthRepository(db *gorm.DB) AuthRepository {
	return &authRepositoryImpl{DB: db}
}

// GetUserByEmail returns a user with the specified email
func (repo *authRepositoryImpl) GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	if err := repo.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// CreateUser creates a new user in the database
func (repo *authRepositoryImpl) CreateUser(user *models.User) error {
	return repo.DB.Create(user).Error
}
