package controllers

import (
	"bucketX/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

// godoc
// @Summary Welcome to bucketX API
// @Description Welcome to bucketX API
// @Produce json
// @Success 200 {object} map[string]interface{} "{"message": "Welcome to bucketX API!"}"
// @Router / [get]
func WelcomeController(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Welcome to bucketX API!",
	})
}

// godoc
//
// @Summary Upload a file
// @Description Uploads a file to the server and returns the file's key and name.
// @Tags files
// @Accept multipart/form-data
// @Produce json
// @Param file formData file true "File to upload"
// @Success 200 {object} map[string]interface{} "{"message": "File uploaded successfully", "filename": "example.txt", "file_key": "unique_file_key"}"
// @Failure 500 {object} map[string]interface{} "{"error": "Detailed error message"}"
// @Router /api/v1/file [post]
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


// godoc
// @Summary Fetch a file
// @Description Fetch a file
// @Accept json
// @Produce json
// @Param file_key path string true "File key"
// @Param tr query string false "Transformation query"
// @Success 200 {file} file
// @Failure 404 {object} map[string]interface{} "{"error": "Detailed error message"}"
// @Router /api/v1/file/{file_key} [get]
func FetchFileController(c *gin.Context) {
	fileKey := c.Param("file_key")

	fileQuery := c.Query("tr")

	filePath, err := services.FetchFilePath(fileKey, fileQuery)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.File(filePath)
}
