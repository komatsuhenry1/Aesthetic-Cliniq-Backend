package dto

import "errors"

type ProfessionalRequestDTO struct {
	ClinicID  string  `json:"clinic_id" binding:"required"`
	UserID    *string `json:"user_id"`
	Name      string  `json:"name" binding:"required"`
	Specialty string  `json:"specialty" binding:"required"`
	Phone     string  `json:"phone" binding:"required"`
	Email     string  `json:"email" binding:"required"`
	Status    string  `json:"status"`
}

func (p *ProfessionalRequestDTO) Validate() error {
	if p.ClinicID == "" {
		return errors.New("clinic_id is required")
	}
	if p.Name == "" {
		return errors.New("name is required")
	}
	if p.Specialty == "" {
		return errors.New("specialty is required")
	}
	if p.Phone == "" {
		return errors.New("phone is required")
	}
	if p.Email == "" {
		return errors.New("email is required")
	}
	return nil
}
