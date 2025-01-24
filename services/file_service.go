package services

import (
	metadataObject "bucketX/services/file_metadataObject"
	transformations "bucketX/services/file_transformations"
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

	if isDuplicate, checkErr := checkDuplicateHash(hashHex); checkErr != nil {
		return "", "", fmt.Errorf("failed to check for duplicate hash: %v", checkErr)
	} else if isDuplicate {
		return "", "", fmt.Errorf("file with the same content already exists")
	}

	filename := filepath.Base(file.Filename)

	if uploadErr := c.SaveUploadedFile(file, filepath.Join("uploads", bucketId, filename)); uploadErr != nil {
		return "", "", fmt.Errorf("failed to save file: %v", uploadErr)
	}

	fileObject := metadataObject.FileDataObject{
		BucketId:   bucketId,
		FileKey:    fileKey,
		Filename:   filename,
		Hash:       hashHex,
		TransForms: make([]string, 0),
	}

	metadataObject.FileMetadataMap[fileKey] = fileObject
	metadataObject.FileHashes[hashHex] = fileKey

	if err := metadataObject.SaveMetadataMapToFile(); err != nil {
		return "", "", fmt.Errorf("failed to save metadata map: %v", err)
	}
	if err := metadataObject.SaveFileHashesToFile(); err != nil {
		return "", "", fmt.Errorf("failed to save file hashes: %v", err)
	}

	return fileKey, filename, nil
}

func checkDuplicateHash(hash string) (bool, error) {
	_, exists := metadataObject.FileHashes[hash]
	return exists, nil
}

func FetchFilePath(fileKey string, fileQuery string) (string, error) {
	fileObject, exists := metadataObject.FileMetadataMap[fileKey]
	if !exists {
		return "", fmt.Errorf("file with key %s does not exist", fileKey)
	}

	if fileQuery != "" {
		for _, value := range fileObject.TransForms {
			if value == fileQuery {
				transformedFileExt := filepath.Ext(fileObject.Filename)
				transformedFilename := fileObject.Filename[:len(fileObject.Filename)-len(transformedFileExt)]
				return filepath.Join("transformed-uploads", fileObject.BucketId, transformedFilename+"_"+fileQuery+transformedFileExt), nil
			}
		}

		return transformations.ApplyTransformations(fileObject.Filename, fileObject.BucketId, fileKey, fileQuery)
	}

	filePath := filepath.Join("uploads", fileObject.BucketId, fileObject.Filename)

	return filePath, nil
}
