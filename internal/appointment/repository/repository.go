package repository

import (
	"clinicprobackend/internal/appointment/model"

	"gorm.io/gorm"
)

type AppointmentRepository interface {
	CreateAppointment(appointment *model.Appointment) error
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
