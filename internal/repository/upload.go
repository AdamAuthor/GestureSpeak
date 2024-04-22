package repository

import (
	"github.com/jmoiron/sqlx"
	"todoApp/internal/models"
)

type UploadVideoPostgres struct {
	db *sqlx.DB
}

func NewUploadVideoPostgres(db *sqlx.DB) *UploadVideoPostgres {
	return &UploadVideoPostgres{db: db}
}

func (r *UploadVideoPostgres) CreateVideoFile(video models.VideoFile) error {
	query := "INSERT INTO video_files (user_id, s3_key, upload_date, processing_status) VALUES ($1, $2, $3, $4) RETURNING id"
	row := r.db.QueryRow(query, video.UserID, video.S3Key, video.UploadDate, video.ProcessingStatus)
	if err := row.Scan(&video.ID); err != nil {
		return err
	}

	return nil
}
