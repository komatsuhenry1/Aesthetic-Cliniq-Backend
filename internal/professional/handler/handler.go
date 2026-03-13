package handler

import (
	"clinicprobackend/internal/professional/dto"
	"clinicprobackend/internal/professional/service"
	"clinicprobackend/internal/utils"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type ProfessionalHandler struct {
	service service.ProfessionalService
}

func NewProfessionalHandler(s service.ProfessionalService) *ProfessionalHandler {
	return &ProfessionalHandler{service: s}
}

func (h *ProfessionalHandler) CreateProfessional(c *gin.Context) {
	var requestDto dto.ProfessionalRequestDTO

	if err := c.ShouldBindJSON(&requestDto); err != nil {
		utils.SendErrorResponse(c, err.Error(), http.StatusBadRequest)
		return
	}

	if err := requestDto.Validate(); err != nil {
		utils.SendErrorResponse(c, err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.service.CreateProfessional(&requestDto); err != nil {
		utils.SendErrorResponse(c, err.Error(), http.StatusInternalServerError)
		return
	}

	utils.SendSuccessResponse(c, "Profissional criado com sucesso.", nil)
}

func (h *ProfessionalHandler) GetAllProfessionals(c *gin.Context) {
	professionals, err := h.service.GetAllProfessionals()
	if err != nil {
		utils.SendErrorResponse(c, "Erro ao buscar profissionais", http.StatusInternalServerError)
		return
	}

	utils.SendSuccessResponse(c, "Profissionais encontrados com sucesso.", professionals)
}

func (h *ProfessionalHandler) GetProfessionalByID(c *gin.Context) {
	id := c.Param("id")

	professional, err := h.service.GetProfessionalByID(id)
	if err != nil {
		utils.SendErrorResponse(c, "Profissional não encontrado", http.StatusNotFound)
		return
	}

	utils.SendSuccessResponse(c, "Profissional encontrado com sucesso.", professional)
}

func (h *ProfessionalHandler) UpdateProfessional(c *gin.Context) {
	id := c.Param("id")

	var updates map[string]interface{}
	if err := c.ShouldBindJSON(&updates); err != nil {
		utils.SendErrorResponse(c, "JSON inválido", http.StatusBadRequest)
		return
	}

	protectedFields := map[string]bool{
		"id":         true,
		"created_at": true,
		"updated_at": true,
	}

	for key := range updates {
		if protectedFields[strings.ToLower(key)] {
			utils.SendErrorResponse(c, fmt.Sprintf("Campo(s) %s não pode ser atualizado.", key), http.StatusBadRequest)
			return
		}
	}

	professional, err := h.service.UpdateProfessionalPartial(id, updates)
	if err != nil {
		utils.SendErrorResponse(c, err.Error(), http.StatusBadRequest)
		return
	}

	utils.SendSuccessResponse(c, "Profissional atualizado com sucesso.", professional)
}

func (h *ProfessionalHandler) DeleteProfessional(c *gin.Context) {
	id := c.Param("id")

	if err := h.service.DeleteProfessional(id); err != nil {
		utils.SendErrorResponse(c, "Erro ao deletar profissional", http.StatusInternalServerError)
		return
	}

	utils.SendSuccessResponse(c, "Profissional deletado com sucesso.", nil)
}
