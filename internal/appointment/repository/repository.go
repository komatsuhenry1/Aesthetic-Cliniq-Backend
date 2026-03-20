package repository

import (
	"clinicprobackend/internal/appointment/dto"
	"clinicprobackend/internal/appointment/model"

	"gorm.io/gorm"
)

type AppointmentRepository interface {
	CreateAppointment(appointment *model.Appointment) error
	GetAppointmentsByDate(date string) ([]dto.AppointmentResponseDTO, error)
	GetAppointmentsByDateAndProfessional(date string, professionalId string) ([]dto.AppointmentResponseDTO, error)
	GetAppointmentsWeek(startDate string, endDate string) ([]dto.AppointmentResponseDTO, error)
	GetAppointmentCountsByMonth(yearMonth string) ([]dto.AppointmentCountDTO, error)
	GetNextFiveAppointments() ([]dto.AppointmentResponseDTO, error)
	UpdateAppointment(id string, updates map[string]interface{}) (*model.Appointment, error)
	DeleteAppointment(id string) error
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

func (r *appointmentRepository) GetAppointmentsByDate(date string) ([]dto.AppointmentResponseDTO, error) {
	var appointments []dto.AppointmentResponseDTO

	err := r.db.Table("appointments").
		Select("appointments.start_time, appointments.end_time, appointments.price, appointments.payment_method, appointments.patient_name as patient_name, appointments.procedure, appointments.professional_name as professional_name, appointments.status").
		Joins("LEFT JOIN users ON users.id = appointments.professional_id").
		Where("DATE(appointments.start_time) = ?", date).
		Order("appointments.start_time ASC").
		Scan(&appointments).Error

	return appointments, err
}

func (r *appointmentRepository) GetAppointmentsWeek(startDate string, endDate string) ([]dto.AppointmentResponseDTO, error) {
	var appointments []dto.AppointmentResponseDTO

	err := r.db.Table("appointments").
		Select("appointments.start_time, appointments.end_time, appointments.price, appointments.payment_method, appointments.patient_name as patient_name, appointments.procedure, appointments.professional_name as professional_name, appointments.status").
		Joins("LEFT JOIN users ON users.id = appointments.professional_id").
		Where("appointments.start_time >= ? AND appointments.start_time <= ?", startDate, endDate).
		Order("appointments.start_time ASC").
		Scan(&appointments).Error

	return appointments, err
}

func (r *appointmentRepository) GetAppointmentsByDateAndProfessional(date string, professionalId string) ([]dto.AppointmentResponseDTO, error) {
	var appointments []dto.AppointmentResponseDTO

	err := r.db.Table("appointments").
		Select("appointments.start_time, appointments.end_time, appointments.price, appointments.payment_method, appointments.patient_name as patient_name, appointments.procedure, appointments.professional_name as professional_name, appointments.status").
		Joins("LEFT JOIN users ON users.id = appointments.professional_id").
		Where("DATE(appointments.start_time) = ? AND appointments.professional_id = ?", date, professionalId).
		Order("appointments.start_time ASC").
		Scan(&appointments).Error

	return appointments, err
}

func (r *appointmentRepository) GetAppointmentCountsByMonth(yearMonth string) ([]dto.AppointmentCountDTO, error) {
	var counts []dto.AppointmentCountDTO

	err := r.db.Table("appointments").
		Select("TO_CHAR(DATE(start_time), 'YYYY-MM-DD') as date, count(*) as count").
		Where("TO_CHAR(start_time, 'YYYY-MM') = ?", yearMonth).
		Group("DATE(start_time)").
		Order("DATE(start_time) ASC").
		Scan(&counts).Error

	return counts, err
}

func (r *appointmentRepository) GetNextFiveAppointments() ([]dto.AppointmentResponseDTO, error) {
	var appointments []dto.AppointmentResponseDTO

	err := r.db.Table("appointments").
		Select("appointments.start_time, appointments.end_time, appointments.price, appointments.payment_method, appointments.patient_name as patient_name, appointments.procedure, appointments.professional_name as professional_name, appointments.status").
		Joins("LEFT JOIN users ON users.id = appointments.professional_id").
		Where("appointments.start_time >= CURRENT_TIMESTAMP - INTERVAL '3 hours' AND DATE(appointments.start_time) = CURRENT_DATE").
		Order("appointments.start_time ASC").
		Limit(5).
		Scan(&appointments).Error

	return appointments, err
}

func (r *appointmentRepository) UpdateAppointment(id string, updates map[string]interface{}) (*model.Appointment, error) {
	var appointment model.Appointment
	if err := r.db.First(&appointment, "id = ?", id).Error; err != nil {
		return nil, err
	}

	if err := r.db.Model(&appointment).Updates(updates).Error; err != nil {
		return nil, err
	}

	return &appointment, nil
}

func (r *appointmentRepository) DeleteAppointment(id string) error {
	var appointment model.Appointment
	if err := r.db.First(&appointment, "id = ?", id).Error; err != nil {
		return err
	}
	return r.db.Delete(&appointment).Error
}
