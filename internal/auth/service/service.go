package service

import (
	"clinicprobackend/internal/auth/repository"
	"clinicprobackend/internal/auth/dto"
	"clinicprobackend/internal/utils"
	"clinicprobackend/internal/auth/model"
	"errors"
	"fmt"
	"strings"
)

type UserService interface {
	RegisterUser(userRequestDTO *dto.UserRequestDTO) error
	LoginUser(requestDto string) (string, error)
}

type userService struct {
	userRepository repository.UserRepository
}

func NewUserService(userRepository repository.UserRepository) UserService {
	return &userService{userRepository: userRepository}
}

func (s *userService) RegisterUser(userRequestDTO *dto.UserRequestDTO) error {
	userRequestDTO.Name = strings.ToLower(userRequestDTO.Name)
	existingUser, err := s.userRepository.GetUserByUserNameOrEmail(userRequestDTO.Name, userRequestDTO.Email)
	if err != nil {
		return err
	}
	if existingUser != nil {
		return errors.New("usuário já existe")
	}

	fmt.Println(userRequestDTO.Password)
	err = utils.HashPassword(&userRequestDTO.Password)
	if err != nil {
		return err
	}

	user := model.User{
		Name:         utils.CapitalizeWords(userRequestDTO.Name),
		Email:        userRequestDTO.Email,
		Password:     userRequestDTO.Password,
		Clinic:       userRequestDTO.Clinic,
		Phone:        userRequestDTO.Phone,
		Role:         "USER",
	}

	err = s.userRepository.CreateUser(&user)
	if err != nil {
		return err
	}

	return nil
}

func (s *userService) LoginUser(requestDto string) (string, error) {
	return "token", nil
}
