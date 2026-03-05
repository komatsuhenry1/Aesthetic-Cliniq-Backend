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

func (h *AppointmentHandler) GetAppointmentsByDate(c *gin.Context) {
	dateStr := c.Param("date")

	if dateStr == "" {
		utils.SendErrorResponse(c, "Data não informada. Ex: /appointment/2026-03-05", http.StatusBadRequest)
		return
	}

	appointments, err := h.service.GetAppointmentsByDate(dateStr)
	if err != nil {
		utils.SendErrorResponse(c, "Erro ao buscar agendamentos do dia", http.StatusInternalServerError)
		return
	}

	utils.SendSuccessResponse(c, "Agendamentos do dia "+dateStr, appointments)
}

func (h *AppointmentHandler) GetAppointmentsWeek(c *gin.Context) {
	appointments, err := h.service.GetAppointmentsWeek()
	if err != nil {
		utils.SendErrorResponse(c, "Erro ao buscar agendamentos da semana", http.StatusInternalServerError)
		return
	}

	utils.SendSuccessResponse(c, "Agendamentos da semana.", appointments)
}
