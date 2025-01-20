package main

import (
	"bucketX/middlewares"
	"bucketX/routes"
	"bucketX/services"
	"log"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func main() {
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	if err := services.LoadMetadataMapFromFile(); err != nil {
		log.Fatalf("Error loading metadata map: %v", err)
	}

	router := gin.Default()

	router.Use(middlewares.LoggerMiddleware(logger))

	routes.RegisterRoutes(router)

	router.Run(":8080")
}
