package services

import (
	"bucketX/services/metadataObject"
	"bucketX/utils"
	"encoding/json"
	"fmt"
	"io"
	"path/filepath"

	"crypto/sha256"

	"github.com/gin-gonic/gin"
)

type File struct {
	FileKey      string
	FileMetadata metadataObject.FileMetadata
}

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

	// Node based on file key
	targetNode, err := GetNodeForKey(fileKey)
	if err != nil {
		return "", "", fmt.Errorf("failed to resolve target node: %v", err)
	}

	if !isLocalNode(targetNode) {
		resp, err := ForwardUploadRequest(c, targetNode)
		if err != nil {
			return "", "", fmt.Errorf("failed to forward upload request: %v", err)
		}
		defer resp.Body.Close()

		var result struct {
			Filename string `json:"filename"`
			FileKey  string `json:"file_key"`
		}
		if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
			return "", "", fmt.Errorf("failed to decode forwarded response: %v", err)
		}

		return result.FileKey, result.Filename, nil
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

	fileExt := filepath.Ext(filename)
	filetype := utils.FindFileType(fileExt)

	fileMetadataObject := metadataObject.FileMetadata{
		BucketId:   bucketId,
		Filename:   filename,
		FileType:   filetype,
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
	targetNode, err := GetNodeForKey(fileKey)
	if err != nil {
		return "", fmt.Errorf("failed to resolve target node: %v", err)
	}

	if !isLocalNode(targetNode) {
		remoteFilePath, err := ForwardFetchRequest(fileKey, fileQuery, targetNode)
		if err != nil {
			return "", fmt.Errorf("failed to fetch file from remote node: %v", err)
		}
		return remoteFilePath, nil
	}

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

	return filepath.Join("uploads", fileObject.BucketId, fileObject.Filename), nil
}

func ListAllFiles(c *gin.Context) ([]File, error) {
	var files []File

	metadataObject.FileMetadataMu.RLock()
	defer metadataObject.FileMetadataMu.RUnlock()

	for fileKey, fileMetadata := range metadataObject.FileMetadataMap {
		files = append(files, File{
			FileKey:      fileKey,
			FileMetadata: fileMetadata,
		})
	}

	return files, nil
}
