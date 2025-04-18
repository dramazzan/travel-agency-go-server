package repositories

import (
	"gorm.io/gorm"
	"server-go/internal/models"
)

type AuthRepository interface {
	GetUserByEmail(email string) (*models.User, error)
	CreateUser(user *models.User) error
	GetUserById(userId uint) (*models.User, error)
	Update(user *models.User) error
}

type authRepositoryImpl struct {
	DB *gorm.DB
}

func NewAuthRepository(db *gorm.DB) AuthRepository {
	return &authRepositoryImpl{DB: db}
}

func (repo *authRepositoryImpl) GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	if err := repo.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (repo *authRepositoryImpl) CreateUser(user *models.User) error {
	return repo.DB.Create(user).Error
}

func (repo *authRepositoryImpl) GetUserById(userId uint) (*models.User, error) {
	var user models.User
	if err := repo.DB.First(&user, userId).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (repo *authRepositoryImpl) Update(user *models.User) error {
	updateData := map[string]interface{}{
		"username": user.Username,
		"email":    user.Email,
	}

	return repo.DB.Model(&models.User{}).Where("id = ?", user.ID).Updates(updateData).Error
}
