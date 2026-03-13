package repository

import (
	"clinicprobackend/internal/professional/model"

	"gorm.io/gorm"
)

type ProfessionalRepository interface {
	CreateProfessional(professional *model.Professional) error
	GetAllProfessionals() ([]model.Professional, error)
	GetProfessionalByID(id string) (*model.Professional, error)
	UpdateProfessionalPartial(id string, updates map[string]interface{}) (*model.Professional, error)
	DeleteProfessional(id string) error
}

type professionalRepository struct {
	db *gorm.DB
}

func NewProfessionalRepository(db *gorm.DB) ProfessionalRepository {
	return &professionalRepository{db: db}
}

func (r *professionalRepository) CreateProfessional(professional *model.Professional) error {
	return r.db.Create(professional).Error
}

func (r *professionalRepository) GetAllProfessionals() ([]model.Professional, error) {
	var professionals []model.Professional
	if err := r.db.Find(&professionals).Error; err != nil {
		return nil, err
	}
	return professionals, nil
}

func (r *professionalRepository) GetProfessionalByID(id string) (*model.Professional, error) {
	var professional model.Professional
	if err := r.db.First(&professional, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &professional, nil
}

func (r *professionalRepository) UpdateProfessionalPartial(id string, updates map[string]interface{}) (*model.Professional, error) {
	var professional model.Professional
	if err := r.db.First(&professional, "id = ?", id).Error; err != nil {
		return nil, err
	}

	if err := r.db.Model(&professional).Updates(updates).Error; err != nil {
		return nil, err
	}

	return &professional, nil
}

func (r *professionalRepository) DeleteProfessional(id string) error {
	return r.db.Delete(&model.Professional{}, "id = ?", id).Error
}
