package controllers

import (
	"net/http"
	"os"

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

	fileContent, filename, err := services.FetchFileContent(fileKey)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	tempFile, err := os.CreateTemp("", filename)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer os.Remove(tempFile.Name())

	if _, err := tempFile.Write(fileContent); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.File(tempFile.Name())
}
