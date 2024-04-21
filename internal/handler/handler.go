package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"todoApp/internal/service"
)

const (
	notFound = "user id not found"
)

type Handler struct {
	service *service.Service
}

func NewHandler(srv *service.Service) *Handler {
	return &Handler{
		service: srv,
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	//gin.SetMode(gin.ReleaseMode)
	router := gin.New()

	router.LoadHTMLGlob("web/public/*")
	router.Static("/web", "./web/")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}
	api := router.Group("/api")
	{
		api.POST("/upload-video", h.uploadVideo)
	}

	return router
}
