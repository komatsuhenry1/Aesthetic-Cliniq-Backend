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
		appointment.GET("/day", middleware.AuthUser(), container.AppointmentHandler.GetAppointmentsByDate)
		appointment.GET("/week", middleware.AuthUser(), container.AppointmentHandler.GetAppointmentsWeek)
		appointment.GET("/next-five", middleware.AuthUser(), container.AppointmentHandler.GetNextFiveAppointments)
		appointment.PATCH("/:id", middleware.AuthUser(), container.AppointmentHandler.UpdateAppointment)
		appointment.DELETE("/:id", middleware.AuthUser(), container.AppointmentHandler.DeleteAppointment)
	}
}
