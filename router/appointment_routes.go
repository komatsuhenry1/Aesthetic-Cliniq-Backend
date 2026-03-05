// router/auth_routes.go
package router

import (
	"clinicprobackend/internal/di"
	"clinicprobackend/internal/middleware"

	"github.com/gin-gonic/gin"
)

func SetupAppointmentRoutes(r *gin.RouterGroup, container *di.Container) {
	appointment := r.Group("/appointment")
	{
		appointment.POST("/", middleware.AuthUser(), container.AppointmentHandler.CreateAppointment)
		appointment.GET("/today", middleware.AuthUser(), container.AppointmentHandler.GetAppointmentsToday)
		appointment.GET("/week", middleware.AuthUser(), container.AppointmentHandler.GetAppointmentsWeek)
	}
}
