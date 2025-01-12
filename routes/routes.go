package routes

import (
	"bucketX/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	api := router.Group("/api")
	{
		api.GET("/", controllers.WelcomeController)
		api.POST("/image", controllers.UploadImageController)
	}
}
