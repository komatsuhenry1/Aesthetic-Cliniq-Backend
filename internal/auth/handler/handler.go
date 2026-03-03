package handler

import (
	"clinicprobackend/internal/auth/service"

	"io"

	"github.com/gin-gonic/gin"
	"clinicprobackend/internal/auth/dto"
	"clinicprobackend/internal/utils"
	"net/http"
)

type UserHandler struct {
	service service.UserService
}

func NewUserHandler(service service.UserService) *UserHandler {
	return &UserHandler{service: service}
}


func (h *UserHandler) RegisterUser(c *gin.Context) {
	var userRequestDto dto.UserRequestDTO

	if err := c.ShouldBindJSON(&userRequestDto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := userRequestDto.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.service.RegisterUser(&userRequestDto)
	if err != nil {
		utils.SendErrorResponse(c, err.Error(), http.StatusBadRequest)
		return
	}

	utils.SendSuccessResponse(c, "Usuário registrado com sucesso.", nil)
}

func (h *UserHandler) LoginUser(c *gin.Context) {
	body, err := io.ReadAll(c.Request.Body)

	token, err := h.service.LoginUser(string(body))
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"token": token,
		"user":  gin.H{"name": "user", "email": "email", "role": "role"},
	})
}
