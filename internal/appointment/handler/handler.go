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

// CreateAppointment godoc
// @Summary      Create a new appointment
// @Description  Creates an appointment record
// @Tags         Appointments
// @Accept       json
// @Produce      json
// @Param        appointment body dto.AppointmentRequestDTO true "Appointment Data"
// @Success      200  {object}  map[string]interface{}
// @Failure      400  {object}  map[string]interface{}
// @Failure      401  {object}  map[string]interface{}
// @Failure      500  {object}  map[string]interface{}
// @Security     BearerAuth
// @Router       /appointment/ [post]
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

// GetAppointmentsByDate godoc
// @Summary      Get appointments by date
// @Description  Retrieves appointments for a specific day
// @Tags         Appointments
// @Produce      json
// @Param        date query     string  true  "Date (YYYY-MM-DD)"
// @Success      200  {object}  map[string]interface{}
// @Failure      400  {object}  map[string]interface{}
// @Failure      500  {object}  map[string]interface{}
// @Security     BearerAuth
// @Router       /appointment/day [get]
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

// GetAppointmentsWeek godoc
// @Summary      Get appointments for a week
// @Description  Retrieves appointments within a date range
// @Tags         Appointments
// @Produce      json
// @Param        start_date query     string  true  "Start Date (YYYY-MM-DD)"
// @Param        end_date   query     string  true  "End Date (YYYY-MM-DD)"
// @Success      200  {object}  map[string]interface{}
// @Failure      400  {object}  map[string]interface{}
// @Failure      500  {object}  map[string]interface{}
// @Security     BearerAuth
// @Router       /appointment/week [get]
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

// GetAppointmentsByDateAndProfessional godoc
// @Summary      Get appointments by date and professional
// @Description  Retrieves appointments for a specific day and professional
// @Tags         Appointments
// @Produce      json
// @Param        date query     string  true  "Date (YYYY-MM-DD)"
// @Param        professional_id query string true "Professional ID"
// @Success      200  {object}  map[string]interface{}
// @Failure      400  {object}  map[string]interface{}
// @Failure      500  {object}  map[string]interface{}
// @Security     BearerAuth
// @Router       /appointment/professional-day [get]
func (h *AppointmentHandler) GetAppointmentsByDateAndProfessional(c *gin.Context) {
	dateStr := c.Query("date")
	professionalId := c.Query("professional_id")

	if dateStr == "" || professionalId == "" {
		utils.SendErrorResponse(c, "Data ou Profissional não informados. Ex: /appointment/professional-day?date=2026-03-05&professional_id=uuid", http.StatusBadRequest)
		return
	}

	appointments, err := h.service.GetAppointmentsByDateAndProfessional(dateStr, professionalId)
	if err != nil {
		utils.SendErrorResponse(c, "Erro ao buscar agendamentos do profissional no dia", http.StatusInternalServerError)
		return
	}

	utils.SendSuccessResponse(c, "Agendamentos do profissional no dia "+dateStr, appointments)
}

// GetAppointmentCountsByMonth godoc
// @Summary      Get appointment counts by month
// @Description  Retrieves the number of appointments per day for a specific month
// @Tags         Appointments
// @Produce      json
// @Param        month query     string  true  "Month (YYYY-MM)"
// @Success      200  {object}  map[string]interface{}
// @Failure      400  {object}  map[string]interface{}
// @Failure      500  {object}  map[string]interface{}
// @Security     BearerAuth
// @Router       /appointment/month-count [get]
func (h *AppointmentHandler) GetAppointmentCountsByMonth(c *gin.Context) {
	monthStr := c.Query("month")

	if monthStr == "" {
		utils.SendErrorResponse(c, "Mês não informado. Ex: /appointment/month-count?month=2026-03", http.StatusBadRequest)
		return
	}

	counts, err := h.service.GetAppointmentCountsByMonth(monthStr)
	if err != nil {
		utils.SendErrorResponse(c, "Erro ao buscar contagens do mês", http.StatusInternalServerError)
		return
	}

	utils.SendSuccessResponse(c, "Contagens do mês "+monthStr, counts)
}

// GetNextFiveAppointments godoc
// @Summary      Get next five appointments
// @Description  Retrieves the next 5 upcoming appointments
// @Tags         Appointments
// @Produce      json
// @Success      200  {object}  map[string]interface{}
// @Failure      500  {object}  map[string]interface{}
// @Security     BearerAuth
// @Router       /appointment/next-five [get]
func (h *AppointmentHandler) GetNextFiveAppointments(c *gin.Context) {
	appointments, err := h.service.GetNextFiveAppointments()
	fmt.Println(appointments)
	if err != nil {
		utils.SendErrorResponse(c, "Erro ao buscar próximos eventos", http.StatusInternalServerError)
		return
	}

	utils.SendSuccessResponse(c, "Próximos 5 agendamentos.", appointments)
}

// UpdateAppointment godoc
// @Summary      Update an appointment
// @Description  Partially updates an appointment record
// @Tags         Appointments
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "Appointment ID"
// @Param        updates body   map[string]interface{} true "Updates"
// @Success      200  {object}  map[string]interface{}
// @Failure      400  {object}  map[string]interface{}
// @Security     BearerAuth
// @Router       /appointment/{id} [patch]
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

// DeleteAppointment godoc
// @Summary      Delete an appointment
// @Description  Deletes an appointment record
// @Tags         Appointments
// @Produce      json
// @Param        id   path      string  true  "Appointment ID"
// @Success      200  {object}  map[string]interface{}
// @Failure      500  {object}  map[string]interface{}
// @Security     BearerAuth
// @Router       /appointment/{id} [delete]
func (h *AppointmentHandler) DeleteAppointment(c *gin.Context) {
	id := c.Param("id")

	if err := h.service.DeleteAppointment(id); err != nil {
		utils.SendErrorResponse(c, "Erro ao deletar agendamento", http.StatusInternalServerError)
		return
	}

	utils.SendSuccessResponse(c, "Agendamento deletado com sucesso.", nil)
}
