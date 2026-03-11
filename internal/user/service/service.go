package service

import (
	"clinicprobackend/internal/user/model"
	"clinicprobackend/internal/user/repository"
)

type UserService interface {
	GetProfessionals() ([]model.User, error)
}

type userService struct {
	userRepository repository.UserRepository
}

func NewUserService(userRepository repository.UserRepository) UserService {
	return &userService{userRepository: userRepository}
}

func (s *userService) GetProfessionals() ([]model.User, error) {
	return s.userRepository.GetProfessionals()
}
