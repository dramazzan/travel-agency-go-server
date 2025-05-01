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
	GetAllUsers() ([]*models.User, error)
	BlockUserByID(id uint) error
	UnblockUserByID(id uint) error
	DeleteUserByID(id uint) error
}

type authRepositoryImpl struct {
	DB *gorm.DB
}

func NewAuthRepository(db *gorm.DB) *authRepositoryImpl {
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
	return repo.DB.Model(&models.User{}).Where("id = ?", user.ID).Updates(user).Error
}

func (repo *authRepositoryImpl) GetAllUsers() ([]*models.User, error) {
	var users []*models.User
	if err := repo.DB.Where("role = ?", "user").Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (repo *authRepositoryImpl) BlockUserByID(id uint) error {
	return repo.DB.Model(&models.User{}).
		Where("id = ?", id).
		Update("is_blocked", true).Error
}

func (repo *authRepositoryImpl) UnblockUserByID(id uint) error {
	return repo.DB.Model(&models.User{}).
		Where("id = ?", id).
		Update("is_blocked", false).Error
}

func (repo *authRepositoryImpl) DeleteUserByID(id uint) error {
	return repo.DB.Delete(&models.User{}, id).Error
}
