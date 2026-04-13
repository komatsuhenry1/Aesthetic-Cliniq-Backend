package handler

import (
	"aestheticcliniq/internal/user/service"
	"aestheticcliniq/internal/utils"
	"net/http"

	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	service service.UserService
}

func NewUserHandler(service service.UserService) *UserHandler {
	return &UserHandler{service: service}
}

// GetProfessionals godoc
// @Summary      Get all professional users
// @Description  Retrieves users with the PROFESSIONAL role
// @Tags         Users
// @Produce      json
// @Success      200  {object}  map[string]interface{}
// @Failure      500  {object}  map[string]interface{}
// @Security     BearerAuth
// @Router       /user/professionals [get]
func (h *UserHandler) GetProfessionals(c *gin.Context) {
	professionals, err := h.service.GetProfessionals()
	if err != nil {
		utils.SendErrorResponse(c, "Erro ao buscar profissionais", http.StatusInternalServerError)
		return
	}

	utils.SendSuccessResponse(c, "Profissionais encontrados com sucesso", professionals)
}

// UpdateUser godoc
// @Summary      Update a user
// @Description  Partially updates a user record
// @Tags         Users
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "User ID"
// @Param        updates body   map[string]interface{} true "Updates"
// @Success      200  {object}  map[string]interface{}
// @Failure      400  {object}  map[string]interface{}
// @Security     BearerAuth
// @Router       /user/{id} [patch]
func (h *UserHandler) UpdateUser(c *gin.Context) {
	userId := c.Param("id")

	var updates map[string]interface{}
	if err := c.ShouldBindJSON(&updates); err != nil {
		utils.SendErrorResponse(c, "JSON inválido", http.StatusBadRequest)
		return
	}

	protectedFields := map[string]bool{
		"id":         true,
		"created_at": true,
		"updated_at": true,
		"password":   true, // Extra safety for user entity updates
	}

	for key := range updates {
		if protectedFields[strings.ToLower(key)] {
			utils.SendErrorResponse(c, fmt.Sprintf("Campo(s) %s não pode ser atualizado.", key), http.StatusBadRequest)
			return
		}
	}

	user, err := h.service.UpdateUserPartial(userId, updates)
	if err != nil {
		utils.SendErrorResponse(c, err.Error(), http.StatusBadRequest)
		return
	}

	utils.SendSuccessResponse(c, "Usuário atualizado com sucesso.", user)
}

// GetAdminDashboard godoc
// @Summary      Admin dashboard summary
// @Description  Returns hardcoded summary metrics for the ADMIN role
// @Tags         Users
// @Produce      json
// @Success      200  {object}  map[string]interface{}
// @Security     BearerAuth
// @Router       /user/dashboard [get]
func (h *UserHandler) GetAdminDashboard(c *gin.Context) {
	data := map[string]interface{}{
		"total_revenue_month":      15800.00,
		"total_appointments_month": 134,
		"new_patients_month":       27,
		"active_professionals":     6,
		"pending_appointments":     8,
		"cancellation_rate":        "4.5%",
		"top_service":              "Limpeza de Pele",
		"upcoming_today": []map[string]interface{}{
			{"time": "09:00", "patient": "Ana Souza", "service": "Botox", "professional": "Dra. Carla"},
			{"time": "10:30", "patient": "João Lima", "service": "Peeling", "professional": "Dr. Marcos"},
			{"time": "14:00", "patient": "Beatriz Costa", "service": "Limpeza de Pele", "professional": "Dra. Carla"},
		},
	}

	utils.SendSuccessResponse(c, "Dashboard do administrador carregado com sucesso.", data)
}
