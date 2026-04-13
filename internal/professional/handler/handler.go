package handler

import (
	"aestheticcliniq/internal/professional/dto"
	"aestheticcliniq/internal/professional/service"
	"aestheticcliniq/internal/utils"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type ProfessionalHandler struct {
	service service.ProfessionalService
}

func NewProfessionalHandler(s service.ProfessionalService) *ProfessionalHandler {
	return &ProfessionalHandler{service: s}
}

func GetClinicId(c *gin.Context) string {
	raw, exists := c.Get("claims")
	if !exists {
		return ""
	}
	claims, ok := raw.(jwt.MapClaims)
	if !ok {
		return ""
	}
	clinicId, ok := claims["clinic_id"].(string)
	if !ok {
		return ""
	}
	return clinicId
}

// CreateProfessional godoc
// @Summary      Create a new professional
// @Description  Creates a user with role PROFESSIONAL
// @Tags         Professionals
// @Accept       json
// @Produce      json
// @Param        professional body dto.ProfessionalRequestDTO true "Professional Data"
// @Success      200  {object}  map[string]interface{}
// @Failure      400  {object}  map[string]interface{}
// @Failure      401  {object}  map[string]interface{}
// @Failure      500  {object}  map[string]interface{}
// @Security     BearerAuth
// @Router       /professionals [post]
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

	clinicId := GetClinicId(c)
	if clinicId == "" {
		utils.SendErrorResponse(c, "Usuário não autenticado", http.StatusUnauthorized)
		return
	}

	if err := h.service.CreateProfessional(&requestDto, clinicId); err != nil {
		utils.SendErrorResponse(c, err.Error(), http.StatusInternalServerError)
		return
	}

	utils.SendSuccessResponse(c, "Profissional criado com sucesso.", nil)
}

// GetAllProfessionals godoc
// @Summary      Get all professionals
// @Description  Retrieves users with role PROFESSIONAL
// @Tags         Professionals
// @Produce      json
// @Success      200  {object}  map[string]interface{}
// @Failure      500  {object}  map[string]interface{}
// @Security     BearerAuth
// @Router       /professionals [get]
func (h *ProfessionalHandler) GetAllProfessionals(c *gin.Context) {
	professionals, err := h.service.GetAllProfessionals()
	if err != nil {
		utils.SendErrorResponse(c, "Erro ao buscar profissionais", http.StatusInternalServerError)
		return
	}

	utils.SendSuccessResponse(c, "Profissionais encontrados com sucesso.", professionals)
}

// GetProfessionalByID godoc
// @Summary      Get a professional by ID
// @Description  Retrieves a user with role PROFESSIONAL by ID
// @Tags         Professionals
// @Produce      json
// @Param        id   path      string  true  "Professional ID"
// @Success      200  {object}  map[string]interface{}
// @Failure      404  {object}  map[string]interface{}
// @Security     BearerAuth
// @Router       /professionals/{id} [get]
func (h *ProfessionalHandler) GetProfessionalByID(c *gin.Context) {
	id := c.Param("id")

	professional, err := h.service.GetProfessionalByID(id)
	if err != nil {
		utils.SendErrorResponse(c, "Profissional não encontrado", http.StatusNotFound)
		return
	}

	utils.SendSuccessResponse(c, "Profissional encontrado com sucesso.", professional)
}

// UpdateProfessional godoc
// @Summary      Update a professional
// @Description  Partially updates a user with role PROFESSIONAL
// @Tags         Professionals
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "Professional ID"
// @Param        updates body   map[string]interface{} true "Updates"
// @Success      200  {object}  map[string]interface{}
// @Failure      400  {object}  map[string]interface{}
// @Security     BearerAuth
// @Router       /professionals/{id} [patch]
func (h *ProfessionalHandler) UpdateProfessional(c *gin.Context) {
	id := c.Param("id")

	var updates map[string]interface{}
	if err := c.ShouldBindJSON(&updates); err != nil {
		utils.SendErrorResponse(c, "JSON inválido", http.StatusBadRequest)
		return
	}

	protectedFields := map[string]bool{
		"id":         true,
		"role":       true,
		"password":   true,
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

// DeleteProfessional godoc
// @Summary      Delete a professional
// @Description  Deletes a user with role PROFESSIONAL
// @Tags         Professionals
// @Produce      json
// @Param        id   path      string  true  "Professional ID"
// @Success      200  {object}  map[string]interface{}
// @Failure      500  {object}  map[string]interface{}
// @Security     BearerAuth
// @Router       /professionals/{id} [delete]
func (h *ProfessionalHandler) DeleteProfessional(c *gin.Context) {
	id := c.Param("id")

	if err := h.service.DeleteProfessional(id); err != nil {
		utils.SendErrorResponse(c, "Erro ao deletar profissional", http.StatusInternalServerError)
		return
	}

	utils.SendSuccessResponse(c, "Profissional deletado com sucesso.", nil)
}

// GetAllNamesAndIds godoc
// @Summary      Get all professional names and IDs
// @Description  Retrieves a lightweight list of professional names and IDs
// @Tags         Professionals
// @Produce      json
// @Success      200  {object}  map[string]interface{}
// @Failure      500  {object}  map[string]interface{}
// @Security     BearerAuth
// @Router       /professionals/names [get]
func (h *ProfessionalHandler) GetAllNamesAndIds(c *gin.Context) {
	professionals, err := h.service.GetAllNamesAndIds()
	if err != nil {
		utils.SendErrorResponse(c, "Erro ao buscar profissionais", http.StatusInternalServerError)
		return
	}

	utils.SendSuccessResponse(c, "Profissionais encontrados com sucesso.", professionals)
}
