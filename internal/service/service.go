package service

import (
	"todoApp/internal/models"
	"todoApp/internal/repository"
)

type Authorization interface {
	CreateUser(user models.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}

type UploadVideo interface {
	CreateVideoFile(video models.VideoFile) error
}

type Service struct {
	Authorization
	UploadVideo
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repo.Authorization),
		UploadVideo:   NewUploadVideoService(repo.UploadVideo),
	}
}
