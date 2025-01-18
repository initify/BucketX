package services

import (
	"bucketX/database"
	"bucketX/models"
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"

	"crypto/sha256"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func SaveUploadedFile(c *gin.Context) (string, string, error) {
	file, err := c.FormFile("file")
	bucketId := c.PostForm("bucket_id")
	fileKey := c.PostForm("file_key")
	if err != nil {
		return "", "", fmt.Errorf("failed to get uploaded file: %v", err)
	}
	if bucketId == "" {
		return "", "", fmt.Errorf("bucket_id is required")
	}
	if fileKey == "" {
		return "", "", fmt.Errorf("file_key is required")
	}

	fileContent, er := file.Open()
	if er != nil {
		return "", "", fmt.Errorf("failed to open uploaded file: %v", er)
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

	databaseUploadErr := uploadFileToDatabase(file, fileKey, bucketId, hashHex)
	if databaseUploadErr != nil {
		return "", "", fmt.Errorf("failed to upload file to database: %v", databaseUploadErr)
	}

	return fileKey, filename, nil
}

func uploadFileToDatabase(file *multipart.FileHeader, fileKey string, bucketId string, hash string) error {
	filename := file.Filename
	fileSize := file.Size

	fileDoc := models.File{
		Filename: filename,
		FileKey:  fileKey,
		BucketId: bucketId,
		Size:     fileSize,
		Hash:     hash,
	}

	if _, err := database.MongoClient.Database("bucketX").Collection("files").InsertOne(database.Ctx, fileDoc); err != nil {
		return fmt.Errorf("failed to insert file to database: %v", err)
	}

	return nil
}

func checkDuplicateHash(bucketId string, hash string) (bool, error) {
	var file models.File

	if err := database.MongoClient.Database("bucketX").Collection("files").FindOne(database.Ctx, bson.M{"bucket_id": bucketId, "hash": hash}).Decode(&file); err != nil {
		if err == mongo.ErrNoDocuments {
			return false, nil
		}
		return false, fmt.Errorf("failed to query database: %v", err)
	}

	return true, nil
}

func FetchFilePath(fileKey string) (string, error) {
	var file models.File
	if err := database.MongoClient.Database("bucketX").Collection("files").FindOne(database.Ctx, bson.M{"file_key": fileKey}).Decode(&file); err != nil {
		if err == mongo.ErrNoDocuments {
			return "", fmt.Errorf("file not found")
		}
		return "", fmt.Errorf("failed to query database: %v", err)
	}

	filePath := filepath.Join("uploads", file.BucketId, file.Filename)

	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return "", fmt.Errorf("file does not exist: %v", err)
	} else if err != nil {
		return "", fmt.Errorf("failed to retrieve file info: %v", err)
	}

	return filePath, nil
}
