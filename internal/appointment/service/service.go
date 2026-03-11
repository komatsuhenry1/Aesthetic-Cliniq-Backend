package service

import (
	"clinicprobackend/internal/appointment/dto"
	"clinicprobackend/internal/appointment/model"
	"clinicprobackend/internal/appointment/repository"
)

type AppointmentService interface {
	CreateAppointment(requestDto *dto.AppointmentRequestDTO) error
	GetAppointmentsByDate(date string) ([]dto.AppointmentResponseDTO, error)
	GetAppointmentsWeek(startDate string, endDate string) ([]dto.AppointmentResponseDTO, error)
	GetNextFiveAppointments() ([]dto.AppointmentResponseDTO, error)
	UpdateAppointment(id string, updates map[string]interface{}) (*model.Appointment, error)
}

type appointmentService struct {
	appointmentRepository repository.AppointmentRepository
}

func NewAppointmentService(repo repository.AppointmentRepository) AppointmentService {
	return &appointmentService{appointmentRepository: repo}
}

func (s *appointmentService) CreateAppointment(requestDto *dto.AppointmentRequestDTO) error {
	appointment := model.Appointment{
		PatientID:        requestDto.PatientID,
		ProfessionalID:   requestDto.ProfessionalID,
		PatientName:      requestDto.PatientName,
		ProfessionalName: requestDto.ProfessionalName,
		Procedure:        requestDto.Procedure,
		Price:            requestDto.Price,
		StartTime:        requestDto.StartTime,
		EndTime:          requestDto.EndTime,
		Notes:            requestDto.Notes,
		PaymentMethod:    requestDto.PaymentMethod,
		Status:           "pendente",
	}

	return s.appointmentRepository.CreateAppointment(&appointment)
}

func (s *appointmentService) GetAppointmentsByDate(date string) ([]dto.AppointmentResponseDTO, error) {
	return s.appointmentRepository.GetAppointmentsByDate(date)
}

func (s *appointmentService) GetAppointmentsWeek(startDate string, endDate string) ([]dto.AppointmentResponseDTO, error) {
	return s.appointmentRepository.GetAppointmentsWeek(startDate, endDate)
}

func (s *appointmentService) GetNextFiveAppointments() ([]dto.AppointmentResponseDTO, error) {
	return s.appointmentRepository.GetNextFiveAppointments()
}

func (s *appointmentService) UpdateAppointment(id string, updates map[string]interface{}) (*model.Appointment, error) {
	return s.appointmentRepository.UpdateAppointment(id, updates)
}

