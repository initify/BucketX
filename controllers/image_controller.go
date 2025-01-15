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
	filename, err := services.SaveUploadedFile(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":  "File uploaded successfully",
		"filename": filename,
	})
}

func FetchFileController(c *gin.Context) {
	filename := c.Param("filename")
	bucketId := c.Param("bucket_id")

	filepath, err := services.FetchFilePath(filename, bucketId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.File(filepath)
}
