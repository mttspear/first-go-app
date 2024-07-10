package main

import (
	"go-app/internal/album"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	album.RegisterRoutes(router)

	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
