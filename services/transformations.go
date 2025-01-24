package services

import (
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"io"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	compression "github.com/nurlantulemisov/imagecompression"
)

func ApplyTransformations(filename string, bucketId string, fileKey string, query string) (string, error) {
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

	fileMetadataObject, isPresent := GetFileMetadata(fileKey)
	if !isPresent {
		return "", fmt.Errorf("file not found")
	}

	fileMetadataObject.TransForms = append(fileMetadataObject.TransForms, query)
	SetFileMetadata(fileKey, fileMetadataObject)

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
