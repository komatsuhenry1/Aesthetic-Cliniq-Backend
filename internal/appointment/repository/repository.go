package repository

import (
	"clinicprobackend/internal/appointment/dto"
	"clinicprobackend/internal/appointment/model"

	"gorm.io/gorm"
)

type AppointmentRepository interface {
	CreateAppointment(appointment *model.Appointment) error
	GetAppointmentsToday() ([]dto.AppointmentResponseDTO, error)
	GetAppointmentsWeek() ([]dto.AppointmentResponseDTO, error)
}

type appointmentRepository struct {
	db *gorm.DB
}

func NewAppointmentRepository(db *gorm.DB) AppointmentRepository {
	return &appointmentRepository{db: db}
}

func (r *appointmentRepository) CreateAppointment(appointment *model.Appointment) error {
	return r.db.Create(appointment).Error
}

func (r *appointmentRepository) GetAppointmentsToday() ([]dto.AppointmentResponseDTO, error) {
	var appointments []dto.AppointmentResponseDTO

	err := r.db.Table("appointments").
		Select("appointments.start_time, appointments.end_time, appointments.patient_id as patient, appointments.procedure, users.name as professional, appointments.status").
		Joins("LEFT JOIN users ON users.id = appointments.professional_id").
		Where("appointments.start_time >= CURRENT_DATE AND appointments.start_time < CURRENT_DATE + INTERVAL '1 day'").
		Order("appointments.start_time ASC").
		Scan(&appointments).Error

	return appointments, err
}

func (r *appointmentRepository) GetAppointmentsWeek() ([]dto.AppointmentResponseDTO, error) {
	var appointments []dto.AppointmentResponseDTO

	err := r.db.Table("appointments").
		Select("appointments.start_time, appointments.end_time, appointments.patient_id as patient, appointments.procedure, users.name as professional, appointments.status").
		Joins("LEFT JOIN users ON users.id = appointments.professional_id").
		Where("appointments.start_time >= date_trunc('week', CURRENT_DATE) AND appointments.start_time < date_trunc('week', CURRENT_DATE) + INTERVAL '1 week'").
		Order("appointments.start_time ASC").
		Scan(&appointments).Error

	return appointments, err
}
