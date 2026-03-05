package di

import (
	"clinicprobackend/config"
	appointmentHandler "clinicprobackend/internal/appointment/handler"
	appointmentRepository "clinicprobackend/internal/appointment/repository"
	appointmentService "clinicprobackend/internal/appointment/service"
	userHandler "clinicprobackend/internal/auth/handler"
	userRepository "clinicprobackend/internal/auth/repository"
	userService "clinicprobackend/internal/auth/service"
)

type Container struct {
	UserHandler        *userHandler.UserHandler
	AppointmentHandler *appointmentHandler.AppointmentHandler
}

func NewContainer() *Container {
	db := config.GetDB()
	//supabaseClient := config.GetClient()

	userRepository := userRepository.NewUserRepository(db)
	userService := userService.NewUserService(userRepository)
	userHandler := userHandler.NewUserHandler(userService)

	appointmentRepo := appointmentRepository.NewAppointmentRepository(db)
	appointmentSvc := appointmentService.NewAppointmentService(appointmentRepo)
	appointmentHdl := appointmentHandler.NewAppointmentHandler(appointmentSvc)

	return &Container{
		UserHandler:        userHandler,
		AppointmentHandler: appointmentHdl,
	}
}
