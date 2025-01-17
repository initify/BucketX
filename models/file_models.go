package models

type File struct {
	Filename string `json:"filename" bson:"filename"`
	BucketId string `json:"bucket_id" bson:"bucket_id"`
	Size     int64  `json:"size" bson:"size"`
	Hash     string `json:"hash" bson:"hash"`
}
