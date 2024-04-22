package models

import "time"

type VideoFile struct {
	ID               int       `json:"id"`
	UserID           int       `json:"user_id"`
	S3Key            string    `json:"s3_key"`
	UploadDate       time.Time `json:"upload_date"`
	ProcessingStatus string    `json:"processing_status"`
}
