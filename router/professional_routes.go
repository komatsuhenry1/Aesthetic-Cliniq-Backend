package router

import (
	"clinicprobackend/internal/di"
	"clinicprobackend/internal/middleware"

	"github.com/gin-gonic/gin"
)

func SetupProfessionalRoutes(r *gin.RouterGroup, container *di.Container) {
	professional  := r.Group("/professional")
	{
		professional.GET("/", middleware.AuthUser(), container.ProfessionalHandler.GetAllProfessionals)
		professional.PATCH("/:id", middleware.AuthUser(), container.ProfessionalHandler.UpdateProfessional)
		professional.POST("/", middleware.AuthUser(), container.ProfessionalHandler.CreateProfessional)
		professional.DELETE("/:id", middleware.AuthUser(), container.ProfessionalHandler.DeleteProfessional)
		professional.GET("/names-ids", middleware.AuthUser(), container.ProfessionalHandler.GetAllNamesAndIds)
	}
}
