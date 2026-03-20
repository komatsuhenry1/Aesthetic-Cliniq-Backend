package handler

import (
	"clinicprobackend/internal/user/service"
	"clinicprobackend/internal/utils"
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
