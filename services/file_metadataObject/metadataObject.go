package metadataObject

import (
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

func LoadMetadataMapFromFile() error {
	file, err := os.Open("file_metadata.json")
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
	file, err := os.Create("file_metadata.json")
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
