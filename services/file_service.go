package services

import (
	"fmt"
	"io"
	"path/filepath"

	"crypto/sha256"

	"github.com/gin-gonic/gin"
)

func SaveUploadedFile(c *gin.Context) (string, string, error) {
	file, err := c.FormFile("file")
	bucketId := c.PostForm("bucket_id")
	fileKey := c.PostForm("file_key")
	if err != nil {
		return "", "", fmt.Errorf("failed to get uploaded file: %v", err)
	}

	if bucketId == "" || fileKey == "" {
		missingFields := []string{}
		if bucketId == "" {
			missingFields = append(missingFields, "bucket_id")
		}
		if fileKey == "" {
			missingFields = append(missingFields, "file_key")
		}
		return "", "", fmt.Errorf("missing required fields: %v", missingFields)
	}

	fileContent, err := file.Open()
	if err != nil {
		return "", "", fmt.Errorf("failed to open file: %v", err)
	}
	defer fileContent.Close()

	h := sha256.New()
	if _, hashErr := io.Copy(h, fileContent); hashErr != nil {
		return "", "", fmt.Errorf("failed to compute hash: %v", hashErr)
	}

	hashHex := fmt.Sprintf("%x", h.Sum(nil))

	if _, isPresent := GetFileKey(hashHex); isPresent {
		return "", "", fmt.Errorf("file with the same content already exists")
	}

	filename := filepath.Base(file.Filename)

	if uploadErr := c.SaveUploadedFile(file, filepath.Join("uploads", bucketId, filename)); uploadErr != nil {
		return "", "", fmt.Errorf("failed to save file: %v", uploadErr)
	}

	fileMetadataObject := FileMetadata{
		BucketId:   bucketId,
		FileKey:    fileKey,
		Filename:   filename,
		Hash:       hashHex,
		TransForms: make([]string, 0),
	}

	SetFileMetadata(fileKey, fileMetadataObject)
	SetFileHash(hashHex, fileKey)

	return fileKey, filename, nil
}

func FetchFilePath(fileKey string, fileQuery string) (string, error) {
	fileObject, isPresent := GetFileMetadata(fileKey)
	if !isPresent {
		return "", fmt.Errorf("file not found")
	}

	if fileQuery != "" {
		for _, value := range fileObject.TransForms {
			if value == fileQuery {
				transformedFileExt := filepath.Ext(fileObject.Filename)
				transformedFilename := fileObject.Filename[:len(fileObject.Filename)-len(transformedFileExt)]
				return filepath.Join("transformed-uploads", fileObject.BucketId, transformedFilename+"_"+fileQuery+transformedFileExt), nil
			}
		}

		return ApplyTransformations(fileObject.Filename, fileObject.BucketId, fileKey, fileQuery)
	}

	filePath := filepath.Join("uploads", fileObject.BucketId, fileObject.Filename)

	return filePath, nil
}
