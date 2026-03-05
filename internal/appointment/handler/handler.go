package handler

import (
	"clinicprobackend/internal/appointment/dto"
	"clinicprobackend/internal/appointment/service"
	"clinicprobackend/internal/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AppointmentHandler struct {
	service service.AppointmentService
}

func NewAppointmentHandler(s service.AppointmentService) *AppointmentHandler {
	return &AppointmentHandler{service: s}
}

func (h *AppointmentHandler) CreateAppointment(c *gin.Context) {
	var requestDto dto.AppointmentRequestDTO

	if err := c.ShouldBindJSON(&requestDto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := requestDto.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.service.CreateAppointment(&requestDto)
	if err != nil {
		utils.SendErrorResponse(c, err.Error(), http.StatusBadRequest)
		return
	}

	utils.SendSuccessResponse(c, "Agendamento criado com sucesso.", nil)
}
