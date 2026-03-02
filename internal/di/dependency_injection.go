package di

import (
	userHandler "clinicprobackend/internal/auth/handler"
	userRepository "clinicprobackend/internal/auth/repository"
	userService "clinicprobackend/internal/auth/service"
)

type Container struct {
	UserHandler *userHandler.UserHandler
}

func NewContainer() *Container {
	//db := config.GetDB()
	//supabaseClient := config.GetClient()

	userRepository := userRepository.NewUserRepository()
	userService := userService.NewUserService(userRepository)
	userHandler := userHandler.NewUserHandler(userService)

	return &Container{
		UserHandler: userHandler,
	}
}
