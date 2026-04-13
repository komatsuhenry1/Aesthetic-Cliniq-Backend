package service

import (
	"aestheticcliniq/internal/professional/dto"
	"aestheticcliniq/internal/professional/repository"
	userModel "aestheticcliniq/internal/user/model"
	"aestheticcliniq/internal/utils"
)

type ProfessionalService interface {
	CreateProfessional(requestDto *dto.ProfessionalRequestDTO, clinicId string) error
	GetAllProfessionals() ([]userModel.User, error)
	GetProfessionalByID(id string) (*userModel.User, error)
	UpdateProfessionalPartial(id string, updates map[string]interface{}) (*userModel.User, error)
	DeleteProfessional(id string) error
	GetAllNamesAndIds() ([]dto.ProfessionalNamesAndIdsDTO, error)
}

type professionalService struct {
	professionalRepository repository.ProfessionalRepository
}

func NewProfessionalService(repo repository.ProfessionalRepository) ProfessionalService {
	return &professionalService{professionalRepository: repo}
}

func (s *professionalService) CreateProfessional(requestDto *dto.ProfessionalRequestDTO, clinicId string) error {

	normalizedEmail, err := utils.EmailRegex(requestDto.Email)
	if err != nil {
		return err
	}

	normalizedPhone, err := utils.ValidatePhone(requestDto.Phone)
	if err != nil {
		return err
	}

	if requestDto.Password != "" {
		if err := utils.ValidatePasswordRegex(requestDto.Password); err != nil {
			return err
		}
	}

	hashedPassword := requestDto.Password
	if hashedPassword == "" {
		hashedPassword = "Password1@"
	}
	
	if err := utils.HashPassword(&hashedPassword); err != nil {
		return err
	}
	status := requestDto.Status
	if status == "" {
		status = "ativo"
	}

	user := userModel.User{
		ClinicID:  clinicId,
		Name:      utils.CapitalizeWords(requestDto.Name),
		Specialty: utils.CapitalizeWords(requestDto.Specialty),
		Phone:     normalizedPhone,
		Email:     normalizedEmail,
		Password:  hashedPassword,
		Role:      "PROFESSIONAL",
		Status:    status,
	}
	return s.professionalRepository.CreateProfessional(&user)
}

func (s *professionalService) GetAllProfessionals() ([]userModel.User, error) {
	return s.professionalRepository.GetAllProfessionals()
}

func (s *professionalService) GetProfessionalByID(id string) (*userModel.User, error) {
	return s.professionalRepository.GetProfessionalByID(id)
}

func (s *professionalService) UpdateProfessionalPartial(id string, updates map[string]interface{}) (*userModel.User, error) {
	return s.professionalRepository.UpdateProfessionalPartial(id, updates)
}

func (s *professionalService) DeleteProfessional(id string) error {
	return s.professionalRepository.DeleteProfessional(id)
}

func (s *professionalService) GetAllNamesAndIds() ([]dto.ProfessionalNamesAndIdsDTO, error) {
	return s.professionalRepository.GetAllNamesAndIds()
}
