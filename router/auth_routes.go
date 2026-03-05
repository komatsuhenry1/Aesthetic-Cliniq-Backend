// router/auth_routes.go
package router

import (
	"clinicprobackend/internal/di"

	"github.com/gin-gonic/gin"
)

func SetupAuthRoutes(r *gin.RouterGroup, container *di.Container) {
	auth := r.Group("/auth")
	{
		auth.POST("/register", container.UserHandler.RegisterUser)
		auth.POST("/login", container.UserHandler.LoginUser)
	}
}
