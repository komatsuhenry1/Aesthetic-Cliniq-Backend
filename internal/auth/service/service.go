package service

import "clinicprobackend/internal/auth/repository"

type UserService interface {
	LoginUser(requestDto string) (string, error)
}

type userService struct {
	userRepository repository.UserRepository
}

func NewUserService(userRepository repository.UserRepository) UserService {
	return &userService{userRepository: userRepository}
}

func (s *userService) LoginUser(requestDto string) (string, error) {
	return "token", nil
}
