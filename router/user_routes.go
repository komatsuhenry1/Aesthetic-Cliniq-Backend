package router

import (
	"clinicprobackend/internal/di"
	"clinicprobackend/internal/middleware"

	"github.com/gin-gonic/gin"
)

func SetupUserRoutes(r *gin.RouterGroup, container *di.Container) {
	user := r.Group("/user")
	{
		user.GET("/professionals", middleware.AuthUser(), container.UserHandler.GetProfessionals)
	}
}
