package services

import (
	"encoding/json"
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"io"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"crypto/sha256"

	"github.com/gin-gonic/gin"
	compression "github.com/nurlantulemisov/imagecompression"
)

type FileDataObject struct {
	BucketId   string
	FileKey    string
	Filename   string
	Hash       string
	TransForms []string
}

var fileMetadataMap = make(map[string]FileDataObject)

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

	if isDuplicate, checkErr := checkDuplicateHash(bucketId, hashHex); checkErr != nil {
		return "", "", fmt.Errorf("failed to check for duplicate hash: %v", checkErr)
	} else if isDuplicate {
		return "", "", fmt.Errorf("file with the same content already exists")
	}

	filename := filepath.Base(file.Filename)

	if uploadErr := c.SaveUploadedFile(file, filepath.Join("uploads", bucketId, filename)); uploadErr != nil {
		return "", "", fmt.Errorf("failed to save file: %v", uploadErr)
	}

	fileObject := FileDataObject{
		BucketId:   bucketId,
		FileKey:    fileKey,
		Filename:   filename,
		Hash:       hashHex,
		TransForms: make([]string, 0),
	}

	fileMetadataMap[fileKey] = fileObject

	if err := saveMetadataMapToFile(); err != nil {
		return "", "", fmt.Errorf("failed to save metadata map: %v", err)
	}

	return fileKey, filename, nil
}

func checkDuplicateHash(bucketId string, hash string) (bool, error) {
	for _, fileObject := range fileMetadataMap {
		if fileObject.BucketId == bucketId && fileObject.Hash == hash {
			return true, nil
		}
	}
	return false, nil
}

func saveMetadataMapToFile() error {
	file, err := os.Create("file_metadata.json")
	if err != nil {
		return fmt.Errorf("failed to create metadata file: %v", err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	if err := encoder.Encode(fileMetadataMap); err != nil {
		return fmt.Errorf("failed to encode metadata map: %v", err)
	}

	return nil
}

func FetchFilePath(fileKey string, fileQuery string) (string, error) {
	fileObject, exists := fileMetadataMap[fileKey]
	if !exists {
		return "", fmt.Errorf("file with key %s does not exist", fileKey)
	}

	if fileQuery != "" {
		for _, value := range fileObject.TransForms {
			if value == fileQuery {
				log.Println("File already transformed")
				transformedFileExt := filepath.Ext(fileObject.Filename)
				transformedFilename := fileObject.Filename[:len(fileObject.Filename)-len(transformedFileExt)]
				return filepath.Join("transformed-uploads", fileObject.BucketId, transformedFilename+"_"+fileQuery+transformedFileExt), nil
			}
		}

		return applyTransformations(fileObject.Filename, fileObject.BucketId, fileKey, fileQuery)
	}

	filePath := filepath.Join("uploads", fileObject.BucketId, fileObject.Filename)

	return filePath, nil
}

func applyTransformations(filename string, bucketId string, fileKey string, query string) (string, error) {
	file, err := os.Open(filepath.Join("uploads", bucketId, filename))
	if err != nil {
		return "", fmt.Errorf("failed to open file: %v", err)
	}
	defer file.Close()

	transformedDir := filepath.Join("transformed-uploads", bucketId)
	transformedFileExt := filepath.Ext(filename)
	transformedFilename := filename[:len(filename)-len(transformedFileExt)]
	transformedFilePath := filepath.Join(transformedDir, transformedFilename+"_"+query+transformedFileExt)

	if _, err := os.Stat(transformedDir); os.IsNotExist(err) {
		err = os.MkdirAll(transformedDir, os.ModePerm)
		if err != nil {
			return "", fmt.Errorf("failed to create directory: %v", err)
		}
	}

	transformedFile, er := os.Create(transformedFilePath)
	if er != nil {
		return "", fmt.Errorf("failed to create transformed file: %v", er)
	}
	defer transformedFile.Close()

	_, copyErr := io.Copy(transformedFile, file)
	if copyErr != nil {
		return "", fmt.Errorf("failed to copy file: %v", copyErr)
	}

	fileObject := fileMetadataMap[fileKey]
	fileObject.TransForms = append(fileObject.TransForms, query)
	fileMetadataMap[fileKey] = fileObject

	if err := saveMetadataMapToFile(); err != nil {
		return "", fmt.Errorf("failed to save metadata map: %v", err)
	}

	pairs := strings.Split(query, ",")
	for _, pair := range pairs {
		kv := strings.Split(pair, "-")
		if len(kv) == 2 {
			applyErr := applyTransformation(transformedFile, kv[0], kv[1])
			if applyErr != nil {
				return "", fmt.Errorf("failed to apply transformation: %v", applyErr)
			}
		}
	}

	return transformedFilePath, nil
}

func applyTransformation(file *os.File, key string, value string) error {
	switch key {
	case "quality":
		{
			quality, err := strconv.Atoi(value)
			if err != nil {
				return fmt.Errorf("failed to parse quality value: %v", err)
			}
			return qualityImageTransformation(file, quality)
		}
	}

	return fmt.Errorf("invalid transformation key")
}

func qualityImageTransformation(file *os.File, quality int) error {
	file.Seek(0, io.SeekStart)
	img, format, err := image.Decode(file)
	if err != nil {
		return fmt.Errorf("failed to decode image: %v", err)
	}

	compressing, er := compression.New(quality)
	if er != nil {
		return fmt.Errorf("failed to compress image: %v", er)
	}

	compressedImage := compressing.Compress(img)

	if _, err := file.Seek(0, 0); err != nil {
		return fmt.Errorf("failed to reset file pointer: %v", err)
	}

	if err := file.Truncate(0); err != nil {
		return fmt.Errorf("failed to truncate file: %v", err)
	}

	switch format {
	case "jpeg", "jpg":
		if err := jpeg.Encode(file, compressedImage, &jpeg.Options{Quality: quality}); err != nil {
			return fmt.Errorf("failed to write compressed image as JPEG: %v", err)
		}
	case "png":
		if err := png.Encode(file, compressedImage); err != nil {
			return fmt.Errorf("failed to write compressed image as PNG: %v", err)
		}
	default:
		return fmt.Errorf("unsupported image format: %s", format)
	}

	return nil
}

func LoadMetadataMapFromFile() error {
	file, err := os.Open("file_metadata.json")
	if err != nil {
		if os.IsNotExist(err) {
			fileMetadataMap = make(map[string]FileDataObject)
			return nil
		}
		return fmt.Errorf("failed to open metadata file: %v", err)
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&fileMetadataMap); err != nil {
		return fmt.Errorf("failed to decode metadata map: %v", err)
	}

	return nil
}
