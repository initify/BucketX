package main

import (
	"bucketX/middlewares"
	"bucketX/routes"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func main() {
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	router := gin.Default()

	router.Use(middlewares.LoggerMiddleware(logger))

	routes.RegisterRoutes(router)

	router.Run(":8080")
}
