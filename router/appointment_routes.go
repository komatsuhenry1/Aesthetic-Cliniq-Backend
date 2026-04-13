// router/auth_routes.go
package router

import (
	"aestheticcliniq/internal/di"
	"aestheticcliniq/internal/middleware"

	"github.com/gin-gonic/gin"
)

func SetupAppointmentRoutes(r *gin.RouterGroup, container *di.Container) {
	appointment := r.Group("/appointment")
	{
		appointment.POST("/", middleware.AuthRoles("ADMIN", "PROFESSIONAL", "USER"), container.AppointmentHandler.CreateAppointment)
		appointment.GET("/day", middleware.AuthRoles("ADMIN", "PROFESSIONAL"), container.AppointmentHandler.GetAppointmentsByDate)
		appointment.GET("/professional-day", middleware.AuthRoles("ADMIN", "PROFESSIONAL"), container.AppointmentHandler.GetAppointmentsByDateAndProfessional)
		appointment.GET("/week", middleware.AuthRoles("ADMIN", "PROFESSIONAL"), container.AppointmentHandler.GetAppointmentsWeek)
		appointment.GET("/month-count", middleware.AuthRoles("ADMIN", "PROFESSIONAL"), container.AppointmentHandler.GetAppointmentCountsByMonth)
		appointment.GET("/next-five", middleware.AuthRoles("ADMIN", "PROFESSIONAL"), container.AppointmentHandler.GetNextFiveAppointments)
		appointment.PATCH("/:id", middleware.AuthRoles("ADMIN", "PROFESSIONAL"), container.AppointmentHandler.UpdateAppointment)
		appointment.DELETE("/:id", middleware.AuthRoles("ADMIN", "PROFESSIONAL"), container.AppointmentHandler.DeleteAppointment)
	}
}
