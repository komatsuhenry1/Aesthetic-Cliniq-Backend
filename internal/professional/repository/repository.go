package repository

import (
	"aestheticcliniq/internal/professional/dto"
	userModel "aestheticcliniq/internal/user/model"

	"gorm.io/gorm"
)

type ProfessionalRepository interface {
	CreateProfessional(user *userModel.User) error
	GetAllProfessionals() ([]userModel.User, error)
	GetProfessionalByID(id string) (*userModel.User, error)
	UpdateProfessionalPartial(id string, updates map[string]interface{}) (*userModel.User, error)
	DeleteProfessional(id string) error
	GetAllNamesAndIds() ([]dto.ProfessionalNamesAndIdsDTO, error)
}

type professionalRepository struct {
	db *gorm.DB
}

func NewProfessionalRepository(db *gorm.DB) ProfessionalRepository {
	return &professionalRepository{db: db}
}

func (r *professionalRepository) CreateProfessional(user *userModel.User) error {
	return r.db.Create(user).Error
}

func (r *professionalRepository) GetAllProfessionals() ([]userModel.User, error) {
	var professionals []userModel.User
	if err := r.db.Where("role = ?", "PROFESSIONAL").Find(&professionals).Error; err != nil {
		return nil, err
	}
	return professionals, nil
}

func (r *professionalRepository) GetProfessionalByID(id string) (*userModel.User, error) {
	var user userModel.User
	if err := r.db.Where("id = ? AND role = ?", id, "PROFESSIONAL").First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *professionalRepository) UpdateProfessionalPartial(id string, updates map[string]interface{}) (*userModel.User, error) {
	var user userModel.User
	if err := r.db.Where("id = ? AND role = ?", id, "PROFESSIONAL").First(&user).Error; err != nil {
		return nil, err
	}

	if err := r.db.Model(&user).Updates(updates).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *professionalRepository) DeleteProfessional(id string) error {
	return r.db.Where("id = ? AND role = ?", id, "PROFESSIONAL").Delete(&userModel.User{}).Error
}

func (r *professionalRepository) GetAllNamesAndIds() ([]dto.ProfessionalNamesAndIdsDTO, error) {
	var professionals []dto.ProfessionalNamesAndIdsDTO
	if err := r.db.Table("users").Select("id", "name").Where("role = ?", "PROFESSIONAL").Find(&professionals).Error; err != nil {
		return nil, err
	}

	return professionals, nil
}
