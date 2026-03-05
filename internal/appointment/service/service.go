package service

import (
	"clinicprobackend/internal/appointment/dto"
	"clinicprobackend/internal/appointment/model"
	"clinicprobackend/internal/appointment/repository"
)

type AppointmentService interface {
	CreateAppointment(requestDto *dto.AppointmentRequestDTO) error
}

type appointmentService struct {
	appointmentRepository repository.AppointmentRepository
}

func NewAppointmentService(repo repository.AppointmentRepository) AppointmentService {
	return &appointmentService{appointmentRepository: repo}
}

func (s *appointmentService) CreateAppointment(requestDto *dto.AppointmentRequestDTO) error {
	appointment := model.Appointment{
		PatientID:      requestDto.PatientID,
		ProfessionalID: requestDto.ProfessionalID,
		Procedure:      requestDto.Procedure,
		Price:          requestDto.Price,
		StartTime:      requestDto.StartTime,
		EndTime:        requestDto.EndTime,
		Notes:          requestDto.Notes,
		PaymentMethod:  requestDto.PaymentMethod,
		Status:         "pendente",
	}

	return s.appointmentRepository.CreateAppointment(&appointment)
}
