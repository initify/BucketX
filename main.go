package main

import (
	"bucketX/database"
	"bucketX/middlewares"
	"bucketX/routes"
	"log"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func main() {
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	if err := database.Connect_to_mongodb(); err != nil {
		log.Fatal("Could not connect to MongoDB")
	}

	router := gin.Default()

	router.Use(middlewares.LoggerMiddleware(logger))

	routes.RegisterRoutes(router)

	router.Run(":8080")
}
