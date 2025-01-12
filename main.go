package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Welcome to bucketX API!",
		})
	})

	router.POST("/image", func(c *gin.Context) {
		file, err := c.FormFile("file")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "File upload failed. Please ensure a file is provided.",
			})
			return
		}

		filename := file.Filename
		if err := c.SaveUploadedFile(file, filename); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Failed to save the uploaded file.",
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message":  "File uploaded successfully",
			"filename": filename,
		})
	})

	router.Run(":8080")
}
