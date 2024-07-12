package main

import (
	"go-app/internal/album"

	"go-app/pkg/logging"

	"github.com/gin-gonic/gin"
)

func main() {
	logging.SetupLogging()
	log := logging.Log

	router := gin.Default()
	album.RegisterRoutes(router)

	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
