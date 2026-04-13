package dto

import "errors"

type ProfessionalRequestDTO struct {
	Name      string `json:"name" binding:"required"`
	Specialty string `json:"specialty" binding:"required"`
	Phone     string `json:"phone" binding:"required"`
	Email     string `json:"email" binding:"required"`
	Password  string `json:"password"`
	Status    string `json:"status"`
}

type ProfessionalNamesAndIdsDTO struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func (p *ProfessionalRequestDTO) Validate() error {
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
