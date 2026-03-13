package model

import "time"

type Professional struct {
	ID        string    `gorm:"primaryKey;type:uuid;default:gen_random_uuid()" json:"id"`
	ClinicID  string    `json:"clinic_id"`
	UserID    *string   `json:"user_id"`
	Name      string    `json:"name"`
	Specialty string    `json:"specialty"`
	Phone     string    `json:"phone"`
	Email     string    `json:"email"`
	Status    string    `json:"status"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}
