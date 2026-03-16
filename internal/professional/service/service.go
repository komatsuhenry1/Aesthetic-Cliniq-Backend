package service

import (
	"clinicprobackend/internal/professional/dto"
	"clinicprobackend/internal/professional/model"
	"clinicprobackend/internal/professional/repository"
	"clinicprobackend/internal/utils"
)

type ProfessionalService interface {
	CreateProfessional(requestDto *dto.ProfessionalRequestDTO, clinicId string) error
	GetAllProfessionals() ([]model.Professional, error)
	GetProfessionalByID(id string) (*model.Professional, error)
	UpdateProfessionalPartial(id string, updates map[string]interface{}) (*model.Professional, error)
	DeleteProfessional(id string) error
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

	status := requestDto.Status
	if status == "" {
		status = "ativo"
	}

	professional := model.Professional{
		ClinicID:  clinicId,
		UserID:    requestDto.UserID,
		Name:      utils.CapitalizeWords(requestDto.Name),
		Specialty: utils.CapitalizeWords(requestDto.Specialty),
		Phone:     normalizedPhone,
		Email:     normalizedEmail,
		Status:    status,
	}
	return s.professionalRepository.CreateProfessional(&professional)
}

func (s *professionalService) GetAllProfessionals() ([]model.Professional, error) {
	return s.professionalRepository.GetAllProfessionals()
}

func (s *professionalService) GetProfessionalByID(id string) (*model.Professional, error) {
	return s.professionalRepository.GetProfessionalByID(id)
}

func (s *professionalService) UpdateProfessionalPartial(id string, updates map[string]interface{}) (*model.Professional, error) {
	return s.professionalRepository.UpdateProfessionalPartial(id, updates)
}

func (s *professionalService) DeleteProfessional(id string) error {
	return s.professionalRepository.DeleteProfessional(id)
}
