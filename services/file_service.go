package services

import (
	"bucketX/services/metadataObject"
	"encoding/json"
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

	presentKey, isPresent := metadataObject.GetFileKey(hashHex)

	filename := filepath.Base(file.Filename)

	if !isPresent {
		uploadErr := c.SaveUploadedFile(file, filepath.Join("uploads", bucketId, filename))
		if uploadErr != nil {
			return "", "", fmt.Errorf("failed to save file: %v", uploadErr)
		}
	} else {
		presentObject, _ := metadataObject.GetFileMetadata(presentKey)
		if presentObject.BucketId == bucketId {
			filename = presentObject.Filename
		}
	}

	fileMetadataObject := metadataObject.FileMetadata{
		BucketId:   bucketId,
		Filename:   filename,
		Hash:       hashHex,
		TransForms: make([]string, 0),
	}

	fileMetadataBytes, _ := json.Marshal(fileMetadataObject)

	metadataMapObject := metadataObject.FileMapType{
		Type:         "METADATA",
		Filekey:      fileKey,
		FileMetadata: string(fileMetadataBytes),
		FileHash:     "",
	}

	hashMapObject := metadataObject.FileMapType{
		Type:         "HASH",
		Filekey:      fileKey,
		FileMetadata: "",
		FileHash:     hashHex,
	}

	metadataObject.AOF.Write(metadataMapObject)
	metadataObject.AOF.Write(hashMapObject)

	metadataObject.SetFileMetadata(fileKey, fileMetadataObject)
	metadataObject.SetFileHash(hashHex, fileKey)

	return fileKey, filename, nil
}

func FetchFilePath(fileKey string, fileQuery string) (string, error) {
	fileObject, isPresent := metadataObject.GetFileMetadata(fileKey)
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
