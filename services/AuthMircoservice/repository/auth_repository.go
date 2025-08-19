package repository

import (
	"auth-micro/models"

	"gorm.io/gorm"
)

type UserRepository interface {
	Authenticate(email string) (*models.User, error)
	Create(user *models.User) error
}

type UserRepoImpl struct {
	DB *gorm.DB
}

func (r *UserRepoImpl) Authenticate(email string) (*models.User, error) {
	var user models.User
	err := r.DB.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepoImpl) Create(user *models.User) error {
	return r.DB.Create(user).Error
}
