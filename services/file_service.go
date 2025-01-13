package services

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func SaveUploadedFile(c *gin.Context) (string, error) {
	file, err := c.FormFile("file")
	bucketId := c.PostForm("bucket_id")
	if err != nil {
		return "", fmt.Errorf("failed to retrieve file: %v", err)
	}

	filename := filepath.Base(file.Filename)
	if err := c.SaveUploadedFile(file, filepath.Join("uploads", bucketId, filename));

	err != nil {
		return "", fmt.Errorf("failed to save file: %v", err)
	}

	return filename, nil
}

func FetchImagePath(filename string, bucketId string) (string, error) {
	filePath := filepath.Join("uploads", bucketId, filename)

	if _, err := os.Stat(filePath); os.IsNotExist(err) {
			return "", fmt.Errorf("file does not exist: %v", err)
	} else if err != nil {
			return "", fmt.Errorf("failed to retrieve file info: %v", err)
	}

	return filePath, nil
}
