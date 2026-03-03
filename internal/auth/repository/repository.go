package repository

import (
	"clinicprobackend/internal/auth/model"
	"errors"

	"gorm.io/gorm"
)

type UserRepository interface {
	LoginUser(requestDto string) (string, error)
	CreateUser(user *model.User) error
	GetUserByUserNameOrEmail(name, email string) (*model.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) CreateUser(user *model.User) error {
	return r.db.Create(user).Error
}

func (r *userRepository) GetUserByUserNameOrEmail(name, email string) (*model.User, error) {
	var user model.User
	if err := r.db.Where("name = ? OR email = ?", name, email).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
	}
	return &user, nil
}

func (r *userRepository) LoginUser(requestDto string) (string, error) {
	return "token", nil
}
