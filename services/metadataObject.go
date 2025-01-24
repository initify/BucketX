package services

import (
	"bucketX/config"
	"encoding/json"
	"fmt"
	"os"
)

type FileDataObject struct {
	BucketId   string
	FileKey    string
	Filename   string
	Hash       string
	TransForms []string
}

var FileMetadataMap = make(map[string]FileDataObject)
var FileHashes = make(map[string]string)

var (
	metadataFile string
	hashesFile   string
)

func Initialize(cfg config.Metadata) error {
	metadataFile = cfg.FilePath + "metadata.json"
	hashesFile = cfg.FilePath + "hashes.json"

	if err := LoadMetadataMapFromFile(); err != nil {
		return err
	}
	return LoadFileHashesFromFile()
}

func LoadMetadataMapFromFile() error {
	file, err := os.Open(metadataFile)
	if err != nil {
		if os.IsNotExist(err) {
			FileMetadataMap = make(map[string]FileDataObject)
			return nil
		}
		return fmt.Errorf("failed to open metadata file: %v", err)
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&FileMetadataMap); err != nil {
		return fmt.Errorf("failed to decode metadata map: %v", err)
	}

	return nil
}

func SaveMetadataMapToFile() error {
	file, err := os.Create(metadataFile)
	if err != nil {
		return fmt.Errorf("failed to create metadata file: %v", err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	if err := encoder.Encode(FileMetadataMap); err != nil {
		return fmt.Errorf("failed to encode metadata map: %v", err)
	}

	return nil
}

func LoadFileHashesFromFile() error {
	file, err := os.Open(hashesFile)
	if err != nil {
		if os.IsNotExist(err) {
			FileHashes = make(map[string]string)
			return nil
		}
		return fmt.Errorf("failed to open file hashes file: %v", err)
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&FileHashes); err != nil {
		return fmt.Errorf("failed to decode file hashes: %v", err)
	}

	return nil
}

func SaveFileHashesToFile() error {
	file, err := os.Create(hashesFile)
	if err != nil {
		return fmt.Errorf("failed to create file hashes file: %v", err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	if err := encoder.Encode(FileHashes); err != nil {
		return fmt.Errorf("failed to encode file hashes: %v", err)
	}

	return nil
}
