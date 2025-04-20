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
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) AuthRepository {
	return &authRepositoryImpl{db: db}
}

func (r *authRepositoryImpl) GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	err := r.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *authRepositoryImpl) CreateUser(user *models.User) error {
	return r.db.Create(user).Error
}

func (r *authRepositoryImpl) GetUserById(userId uint) (*models.User, error) {
	var user models.User
	err := r.db.First(&user, userId).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *authRepositoryImpl) Update(user *models.User) error {
	return r.db.Model(&models.User{}).
		Where("id = ?", user.ID).
		Updates(map[string]interface{}{
			"username": user.Username,
			"email":    user.Email,
		}).Error
}
