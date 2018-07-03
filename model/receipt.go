package model

type Receipt struct {
	ID   string `json:"ID" bson:"_id"`
	ScanURL string `json:"scanUrl" bson:"scanUrl"`
	UploadDate int64  `json:"uploadDate" bson:"uploadDate"`
}