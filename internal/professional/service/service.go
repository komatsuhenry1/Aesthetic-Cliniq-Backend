package service

import (
	"clinicprobackend/internal/professional/dto"
	"clinicprobackend/internal/professional/model"
	"clinicprobackend/internal/professional/repository"
)

type ProfessionalService interface {
	CreateProfessional(requestDto *dto.ProfessionalRequestDTO) error
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

func (s *professionalService) CreateProfessional(requestDto *dto.ProfessionalRequestDTO) error {
	status := requestDto.Status
	if status == "" {
		status = "ativo"
	}

	professional := model.Professional{
		ClinicID:  requestDto.ClinicID,
		UserID:    requestDto.UserID,
		Name:      requestDto.Name,
		Specialty: requestDto.Specialty,
		Phone:     requestDto.Phone,
		Email:     requestDto.Email,
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
