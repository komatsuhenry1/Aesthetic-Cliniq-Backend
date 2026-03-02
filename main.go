package main

import (
	"clinicprobackend/config"
	"clinicprobackend/router"
	"os"
)

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
