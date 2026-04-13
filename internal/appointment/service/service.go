package service

import (
	"aestheticcliniq/internal/appointment/dto"
	"aestheticcliniq/internal/appointment/model"
	"aestheticcliniq/internal/appointment/repository"
	"aestheticcliniq/internal/utils"
	"fmt"
)

type AppointmentService interface {
	CreateAppointment(requestDto *dto.AppointmentRequestDTO, userId string, userRole string) error
	GetAppointmentsByDate(date string) ([]dto.AppointmentResponseDTO, error)
	GetAppointmentsByDateAndProfessional(date string, professionalId string) ([]dto.AppointmentResponseDTO, error)
	GetAppointmentsWeek(startDate string, endDate string) ([]dto.AppointmentResponseDTO, error)
	GetAppointmentCountsByMonth(yearMonth string) ([]dto.AppointmentCountDTO, error)
	GetNextFiveAppointments() ([]dto.AppointmentResponseDTO, error)
	UpdateAppointment(id string, updates map[string]interface{}) (*model.Appointment, error)
	DeleteAppointment(id string) error
}

type appointmentService struct {
	appointmentRepository repository.AppointmentRepository
}

func NewAppointmentService(repo repository.AppointmentRepository) AppointmentService {
	return &appointmentService{appointmentRepository: repo}
}

func (s *appointmentService) CreateAppointment(requestDto *dto.AppointmentRequestDTO, userId string, userRole string) error {

	if userRole == "USER"{
		fmt.Println("caiu no if")
		
		requestDto.PatientID = userId
	}


	fmt.Println("===========================: ")
	fmt.Println("userRole: ", userRole)
	fmt.Println("patientID: ", userId)
	fmt.Println("===========================: ")

	formatedNote := utils.CapitalizeFirstWord(requestDto.Notes)

	appointment := model.Appointment{
		PatientID:        requestDto.PatientID,
		ProfessionalID:   requestDto.ProfessionalID,
		PatientName:      requestDto.PatientName,
		ProfessionalName: requestDto.ProfessionalName,
		Procedure:        requestDto.Procedure,
		Price:            requestDto.Price,
		StartTime:        requestDto.StartTime,
		EndTime:          requestDto.EndTime,
		Notes:            formatedNote,
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

func (s *appointmentService) GetAppointmentsByDateAndProfessional(date string, professionalId string) ([]dto.AppointmentResponseDTO, error) {
	return s.appointmentRepository.GetAppointmentsByDateAndProfessional(date, professionalId)
}

func (s *appointmentService) GetAppointmentCountsByMonth(yearMonth string) ([]dto.AppointmentCountDTO, error) {
	return s.appointmentRepository.GetAppointmentCountsByMonth(yearMonth)
}

func (s *appointmentService) GetNextFiveAppointments() ([]dto.AppointmentResponseDTO, error) {
	return s.appointmentRepository.GetNextFiveAppointments()
}

func (s *appointmentService) UpdateAppointment(id string, updates map[string]interface{}) (*model.Appointment, error) {
	return s.appointmentRepository.UpdateAppointment(id, updates)
}

func (s *appointmentService) DeleteAppointment(id string) error {
	return s.appointmentRepository.DeleteAppointment(id)
}
