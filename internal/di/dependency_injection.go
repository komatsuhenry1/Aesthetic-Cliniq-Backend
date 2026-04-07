package di

import (
	"aestheticcliniq/config"
	appointmentHandler "aestheticcliniq/internal/appointment/handler"
	appointmentRepository "aestheticcliniq/internal/appointment/repository"
	appointmentService "aestheticcliniq/internal/appointment/service"
	authHandler "aestheticcliniq/internal/auth/handler"
	authService "aestheticcliniq/internal/auth/service"
	userHandler "aestheticcliniq/internal/user/handler"
	userRepository "aestheticcliniq/internal/user/repository"
	userService "aestheticcliniq/internal/user/service"
	professionalHandler "aestheticcliniq/internal/professional/handler"
	professionalRepository "aestheticcliniq/internal/professional/repository"
	professionalService "aestheticcliniq/internal/professional/service"
)

type Container struct {
	AuthHandler        *authHandler.UserHandler
	UserHandler        *userHandler.UserHandler
	AppointmentHandler *appointmentHandler.AppointmentHandler
	ProfessionalHandler *professionalHandler.ProfessionalHandler		
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

	professionalRepo := professionalRepository.NewProfessionalRepository(db)
	professionalSvc := professionalService.NewProfessionalService(professionalRepo)
	professionalHdl := professionalHandler.NewProfessionalHandler(professionalSvc)

	return &Container{
		AuthHandler:        authHdl,
		UserHandler:        userHdl,
		AppointmentHandler: appointmentHdl,
		ProfessionalHandler: professionalHdl,
	}
}
