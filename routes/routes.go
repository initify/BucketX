package routes

import (
	"bucketX/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	api := router.Group("/api")
	{
		api.GET("/", controllers.WelcomeController)
		api.GET("/image/:bucket_id/:filename", controllers.FetchImageController)
		api.POST("/image", controllers.UploadImageController)
	}
}
