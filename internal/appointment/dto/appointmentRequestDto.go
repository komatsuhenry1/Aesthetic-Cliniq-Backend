package dto

import (
	"errors"
	"time"
)

type AppointmentRequestDTO struct {
	PatientID      string    `json:"client_id" binding:"required"`
	ProfessionalID string    `json:"professional_id" binding:"required"`
	Procedure      string    `json:"procedure" binding:"required"`
	Price          float64   `json:"price" binding:"required"`
	StartTime      time.Time `json:"start_time" binding:"required"`
	EndTime        time.Time `json:"end_time" binding:"required"`
	PaymentMethod  string    `json:"payment_method" binding:"required"`
	Notes          string    `json:"notes" binding:"required"`
}

func (a *AppointmentRequestDTO) Validate() error {
	if a.PatientID == "" {
		return errors.New("client_id is required")
	}
	if a.ProfessionalID == "" {
		return errors.New("professional_id is required")
	}
	if a.Procedure == "" {
		return errors.New("procedure is required")
	}
	if a.Price < 0 {
		return errors.New("price cannot be negative")
	}
	if a.StartTime.IsZero() {
		return errors.New("start_time is required")
	}
	if a.EndTime.IsZero() {
		return errors.New("end_time is required")
	}
	if a.StartTime.After(a.EndTime) {
		return errors.New("start_time cannot be after end_time")
	}
	return nil
}
