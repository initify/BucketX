package services

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func SaveUploadedFile(c *gin.Context) (string, error) {
	file, err := c.FormFile("file")
	bucketId := c.PostForm("bucket_id")
	if err != nil {
		return "", fmt.Errorf("failed to retrieve file: %v", err)
	}

	if bucketId == "" {
		return "", fmt.Errorf("bucket_id is required")
	}

	filename := filepath.Base(file.Filename)
	uuidFilename := uuid.New().String() + "-" + filename
	
	if err := c.SaveUploadedFile(file, filepath.Join("uploads", bucketId, uuidFilename));

	err != nil {
		return "", fmt.Errorf("failed to save file: %v", err)
	}

	return uuidFilename, nil
}

func FetchFilePath(filename string, bucketId string) (string, error) {
	filePath := filepath.Join("uploads", bucketId, filename)
	fmt.Println(filePath)
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
			return "", fmt.Errorf("file does not exist: %v", err)
	} else if err != nil {
			return "", fmt.Errorf("failed to retrieve file info: %v", err)
	}

	return filePath, nil
}
