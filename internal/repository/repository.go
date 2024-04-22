package repository

import (
	"todoApp/internal/models"

	"github.com/jmoiron/sqlx"
)

const (
	usersTable = "users"
)

type Authorization interface {
	CreateUser(user models.User) (int, error)
	GetUser(username string) (models.User, error)
}

type UploadVideo interface {
	CreateVideoFile(video models.VideoFile) error
}

type Repository struct {
	Authorization
	UploadVideo
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		UploadVideo:   NewUploadVideoPostgres(db),
	}
}
