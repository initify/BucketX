package routes

import (
	"bucketX/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	api := router.Group("/api/v1")
	{
		api.GET("/", controllers.WelcomeController)
		api.GET("/file/:file_key", controllers.FetchFileController)
		api.POST("/file", controllers.UploadFileController)
		// godoc
		// @Summary Health check
		// @Description Health check
		// @Produce json
		// @Success 200 {object} gin.H
		// @Router /health [get]
		api.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status":   "ok",
			"response": "pong",
		})
	})
	}
}
