package main

import (
	"aestheticcliniq/config"
	"aestheticcliniq/router"
	"os"
)

// @title           AestheticCliniq API
// @version         1.0
// @description     API backend in Go for AestheticCliniq.
// @host            localhost:8080
// @BasePath        /api/v1
//
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {
	// go middleware.CleanupClients()
	if err := config.ConnectDB(); err != nil {
    	panic(err) // ou log.Fatal(err)
  	}
	// utils.CreateKeys()
	r := router.InitRouter()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	r.Run("0.0.0.0:" + port)
}
