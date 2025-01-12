package services

import (
	"fmt"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func SaveUploadedFile(c *gin.Context) (string, error) {
	file, err := c.FormFile("file")
	if err != nil {
		return "", fmt.Errorf("failed to retrieve file: %v", err)
	}

	filename := filepath.Base(file.Filename)
	if err := c.SaveUploadedFile(file, filename); err != nil {
		return "", fmt.Errorf("failed to save file: %v", err)
	}

	return filename, nil
}
