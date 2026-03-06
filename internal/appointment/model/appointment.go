package model

import "time"

type Appointment struct {
	ID             string    `gorm:"primaryKey;type:uuid;default:gen_random_uuid()" json:"id"`
	PatientID      string    `gorm:"type:uuid;not null" json:"patient_id"`
	ProfessionalID string    `gorm:"type:uuid;not null" json:"professional_id"`
	PatientName    string    `gorm:"type:varchar(255);not null" json:"patient_name"`
	ProfessionalName string    `gorm:"type:varchar(255);not null" json:"professional_name"`
	Procedure      string    `gorm:"type:varchar(255);not null" json:"procedure"`
	Price          float64   `gorm:"type:decimal(10,2);not null" json:"price"`
	StartTime      time.Time `gorm:"not null" json:"start_time"`
	EndTime        time.Time `gorm:"not null" json:"end_time"`
	Notes          string    `gorm:"type:text;not null" json:"notes"`
	PaymentMethod  string    `gorm:"type:varchar(20);not null" json:"payment_method"`
	Status         string    `gorm:"type:varchar(20);not null" json:"status"` // "confirmado", "pendente", "em-andamento", "concluido"
	CreatedAt      time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt      time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}
