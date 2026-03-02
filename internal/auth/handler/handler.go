package handler

import (
	"clinicprobackend/internal/auth/service"

	"io"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	service service.UserService
}

func NewUserHandler(service service.UserService) *UserHandler {
	return &UserHandler{service: service}
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
