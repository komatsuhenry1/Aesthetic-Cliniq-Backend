package router

import (
	"aestheticcliniq/internal/di"
	"aestheticcliniq/internal/middleware"

	"github.com/gin-gonic/gin"
)

func SetupProfessionalRoutes(r *gin.RouterGroup, container *di.Container) {
	professional  := r.Group("/professional")
	{
		professional.GET("/", middleware.AuthRoles("ADMIN", "PROFESSIONAL", "USER"), container.ProfessionalHandler.GetAllProfessionals)
		professional.PATCH("/:id", middleware.AuthRoles("ADMIN", "PROFESSIONAL"), container.ProfessionalHandler.UpdateProfessional)
		professional.POST("/", middleware.AuthRoles("ADMIN", "PROFESSIONAL"), container.ProfessionalHandler.CreateProfessional)
		professional.DELETE("/:id", middleware.AuthRoles("ADMIN", "PROFESSIONAL"), container.ProfessionalHandler.DeleteProfessional)
		professional.GET("/names-ids", middleware.AuthRoles("ADMIN", "PROFESSIONAL"), container.ProfessionalHandler.GetAllNamesAndIds)
		professional.GET("/dashboard", middleware.AuthRoles("PROFESSIONAL"), container.ProfessionalHandler.GetProfessionalDashboard)
	}
}
