package services

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

type BucketData struct {
	BucketName string `json:"bucket_name"`
	Size       int64  `json:"size"`
	FileCount  int    `json:"file_count"`
}

func ListAllBuckets(c *gin.Context) ([]BucketData, error) {
	var buckets []BucketData

	entries, err := os.ReadDir("uploads")
	if err != nil {
		fmt.Errorf("Error reading uploads directory: %v", err)
		return nil, err
	}

	for _, entry := range entries {
		if entry.IsDir() {
			bucketPath := filepath.Join("uploads", entry.Name())
			size, fileCount := getBucketInfo(bucketPath)

			buckets = append(buckets, BucketData{
				BucketName: entry.Name(),
				Size:       size,
				FileCount:  fileCount,
			})
		}
	}

	return buckets, nil
}

func getBucketInfo(bucketPath string) (int64, int) {
	var totalSize int64
	var fileCount int

	err := filepath.Walk(bucketPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			totalSize += info.Size()
			fileCount++
		}
		return nil
	})

	if err != nil {
		fmt.Println("Error reading bucket:", err)
	}

	return totalSize, fileCount
}

func CreateBucket(bucketName string) error {
	bucketPath := filepath.Join("uploads", bucketName)
	err := os.MkdirAll(bucketPath, os.ModePerm)
	if err!= nil {
		return fmt.Errorf("Error creating bucket: %v", err)
	}
	return nil
}
