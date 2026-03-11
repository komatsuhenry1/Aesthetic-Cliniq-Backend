package handler

import (
	"clinicprobackend/internal/user/service"
	"clinicprobackend/internal/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	service service.UserService
}

func NewUserHandler(service service.UserService) *UserHandler {
	return &UserHandler{service: service}
}

func (h *UserHandler) GetProfessionals(c *gin.Context) {
	professionals, err := h.service.GetProfessionals()
	if err != nil {
		utils.SendErrorResponse(c, "Erro ao buscar profissionais", http.StatusInternalServerError)
		return
	}

	utils.SendSuccessResponse(c, "Profissionais encontrados com sucesso", professionals)
}
