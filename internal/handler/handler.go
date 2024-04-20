package handler

import (
	"todoApp/internal/service"
	"todoApp/pkg/logger"

	"github.com/gin-gonic/gin"
)

const (
	notFound = "user id not found"
)

type Handler struct {
	log     logger.Logger
	service *service.Service
}

func NewHandler(log logger.Logger, srv *service.Service) *Handler {
	return &Handler{
		log:     log,
		service: srv,
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	//gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}

	return router
}
