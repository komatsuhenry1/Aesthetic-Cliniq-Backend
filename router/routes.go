package router

import (
	"clinicprobackend/internal/di"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	container := di.NewContainer()
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{
			"*",
			"http://localhost:3000",      // Para acesso local via localhost
			"http://192.168.18.153:3000", // Para acesso via IP na rede local
		},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// r.GET("/health", container.UserHandler.HealthCheck)

	// r.POST("/call", middleware.AuthUser(), container.CallHandler.CallRegister)
	// r.GET("/calls", middleware.AuthUser(), container.CallHandler.FindCallsByStatus)
	// r.PATCH("/solve/:id", middleware.AuthUser(), container.CallHandler.SolveCall)

	r.POST("/register", container.UserHandler.RegisterUser)
	r.POST("/login", container.UserHandler.LoginUser)

	r.POST("/appointment", container.AppointmentHandler.CreateAppointment)

	// r.POST("/refresh-token", container.UserHandler.RefreshToken)
	// // r.GET("/health", container.UserHandler.CheckMyIp)

	// r.POST("/encode", container.LogHandler.EncryptLogByText)
	// r.POST("/decode", container.LogHandler.DecryptLogByText)

	// r.GET("/logs", container.LogHandler.GetLogByCode)
	// r.POST("/upload-file", container.LogHandler.EncryptLogByFile)
	// r.POST("/upload", container.LogHandler.DecryptLogByFile)

	// r.GET("/ips", container.IpHandler.GetAllIps)
	// r.GET("/:ip", container.IpHandler.DecryptIp)

	return r
}
