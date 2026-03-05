// router/auth_routes.go
package router

import (
	"clinicprobackend/internal/di"

	"github.com/gin-gonic/gin"
)

func SetupAppointmentRoutes(r *gin.RouterGroup, container *di.Container) {
	appointment := r.Group("/appointment")
	{
		appointment.POST("/", container.AppointmentHandler.CreateAppointment)
	}
}
