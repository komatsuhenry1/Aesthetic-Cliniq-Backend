package handler

import (
	"clinicprobackend/internal/appointment/dto"
	"clinicprobackend/internal/appointment/service"
	"clinicprobackend/internal/utils"
	"fmt"
	"net/http"
	"strings"

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
	dateStr := c.Query("date")

	if dateStr == "" {
		utils.SendErrorResponse(c, "Data não informada. Ex: /appointment?date=2026-03-05", http.StatusBadRequest)
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
	startDate := c.Query("start_date")
	endDate := c.Query("end_date")

	if startDate == "" || endDate == "" {
		utils.SendErrorResponse(c, "Data inicial e final não informadas. Ex: /appointment/week?start_date=2026-03-01&end_date=2026-03-07", http.StatusBadRequest)
		return
	}

	appointments, err := h.service.GetAppointmentsWeek(startDate, endDate)
	if err != nil {
		utils.SendErrorResponse(c, "Erro ao buscar agendamentos do período", http.StatusInternalServerError)
		return
	}

	utils.SendSuccessResponse(c, "Agendamentos do período", appointments)
}

func (h *AppointmentHandler) GetNextFiveAppointments(c *gin.Context) {
	appointments, err := h.service.GetNextFiveAppointments()
	fmt.Println(appointments)
	if err != nil {
		utils.SendErrorResponse(c, "Erro ao buscar próximos eventos", http.StatusInternalServerError)
		return
	}

	utils.SendSuccessResponse(c, "Próximos 5 agendamentos.", appointments)
}

func (h *AppointmentHandler) UpdateAppointment(c *gin.Context) {
	appointmentId := c.Param("id")

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
	
	appointment, err := h.service.UpdateAppointment(appointmentId, updates)
	if err != nil{
		utils.SendErrorResponse(c, err.Error(), http.StatusBadRequest)
		return
	}

	utils.SendSuccessResponse(c, "Agendamento atualizado com sucesso.", appointment)
}

func (h *AppointmentHandler) DeleteAppointment(c *gin.Context) {
	id := c.Param("id")

	if err := h.service.DeleteAppointment(id); err != nil {
		utils.SendErrorResponse(c, "Erro ao deletar agendamento", http.StatusInternalServerError)
		return
	}

	utils.SendSuccessResponse(c, "Agendamento deletado com sucesso.", nil)
}
