package di

import (
	"clinicprobackend/config"
	appointmentHandler "clinicprobackend/internal/appointment/handler"
	appointmentRepository "clinicprobackend/internal/appointment/repository"
	appointmentService "clinicprobackend/internal/appointment/service"
	authHandler "clinicprobackend/internal/auth/handler"
	authService "clinicprobackend/internal/auth/service"
	userHandler "clinicprobackend/internal/user/handler"
	userRepository "clinicprobackend/internal/user/repository"
	userService "clinicprobackend/internal/user/service"
)

type Container struct {
	AuthHandler        *authHandler.UserHandler
	UserHandler        *userHandler.UserHandler
	AppointmentHandler *appointmentHandler.AppointmentHandler
}

func NewContainer() *Container {
	db := config.GetDB()
	//supabaseClient := config.GetClient()

	userRepo := userRepository.NewUserRepository(db)

	authSvc := authService.NewUserService(userRepo)
	authHdl := authHandler.NewUserHandler(authSvc)

	userSvc := userService.NewUserService(userRepo)
	userHdl := userHandler.NewUserHandler(userSvc)

	appointmentRepo := appointmentRepository.NewAppointmentRepository(db)
	appointmentSvc := appointmentService.NewAppointmentService(appointmentRepo)
	appointmentHdl := appointmentHandler.NewAppointmentHandler(appointmentSvc)

	return &Container{
		AuthHandler:        authHdl,
		UserHandler:        userHdl,
		AppointmentHandler: appointmentHdl,
	}
}
