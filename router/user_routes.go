package router

import (
	"aestheticcliniq/internal/di"
	"aestheticcliniq/internal/middleware"

	"github.com/gin-gonic/gin"
)

func SetupUserRoutes(r *gin.RouterGroup, container *di.Container) {
	user := r.Group("/user")
	{
		user.GET("/professionals", middleware.AuthRoles("ADMIN", "PROFESSIONAL"), container.UserHandler.GetProfessionals)
		user.PATCH("/:id", middleware.AuthRoles("ADMIN", "PROFESSIONAL"), container.UserHandler.UpdateUser)
	}
}
