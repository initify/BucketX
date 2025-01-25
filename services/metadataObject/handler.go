package metadataObject

import (
	"encoding/json"
	"log"
	"sync"
)

var Handlers = map[string]func(FileMapType){
	"METADATA": setMetadata,
	"HASH":     setHash,
}

type FileMetadata struct {
	BucketId   string
	Filename   string
	Hash       string
	TransForms []string
}

var FileMetadataMap = make(map[string]FileMetadata)
var FileMetadataMu = sync.RWMutex{}

var FileHashesMap = make(map[string]string)
var FileHashesMu = sync.RWMutex{}

func SetFileMetadata(fileKey string, fileMetadata FileMetadata) {
	FileMetadataMu.Lock()
	FileMetadataMap[fileKey] = fileMetadata
	FileMetadataMu.Unlock()
}

func GetFileMetadata(fileKey string) (FileMetadata, bool) {
	FileMetadataMu.RLock()
	fileMetadata, exists := FileMetadataMap[fileKey]
	FileMetadataMu.RUnlock()

	if exists {
		return fileMetadata, true
	}

	return FileMetadata{}, false
}

func SetFileHash(hash string, fileKey string) {
	FileHashesMu.Lock()
	FileHashesMap[hash] = fileKey
	FileHashesMu.Unlock()
}

func GetFileKey(hash string) (string, bool) {
	FileHashesMu.RLock()
	fileKey, exists := FileHashesMap[hash]
	FileHashesMu.RUnlock()

	if exists {
		return fileKey, true
	}

	return "", false
}

func setMetadata(value FileMapType) {
	obj, ok := StringToMetadataObj(value.FileMetadata)
	if !ok {
		return
	}
	SetFileMetadata(value.Filekey, obj)
}

func setHash(value FileMapType) {
	SetFileHash(value.FileHash, value.Filekey)
}

func StringToMetadataObj(s string) (FileMetadata, bool) {
	var obj FileMetadata
	err := json.Unmarshal([]byte(s), &obj)
	if err != nil {
		log.Fatal(err)
		return obj, false
	}
	return obj, true
}
