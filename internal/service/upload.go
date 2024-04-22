package service

import (
	"todoApp/internal/models"
	"todoApp/internal/repository"
)

type UploadVideoService struct {
	repo repository.UploadVideo
}

func NewUploadVideoService(repo repository.UploadVideo) *UploadVideoService {
	return &UploadVideoService{repo: repo}
}

func (u *UploadVideoService) CreateVideoFile(video models.VideoFile) error {
	return u.repo.CreateVideoFile(video)
}
