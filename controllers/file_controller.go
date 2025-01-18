package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"bucketX/services"
)

func WelcomeController(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Welcome to bucketX API!",
	})
}

func UploadFileController(c *gin.Context) {
	fileKey, filename, err := services.SaveUploadedFile(c)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":  "File uploaded successfully",
		"filename": filename,
		"file_key": fileKey,
	})
}

func FetchFileController(c *gin.Context) {
	fileKey := c.Param("file_key")

	filepath, err := services.FetchFilePath(fileKey)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.File(filepath)
}
