package handler

import (
	"errors"
	"net/http"
	"strings"
	"todoApp/internal/models"

	"github.com/gin-gonic/gin"
)

const (
	authorizationHeader = "Authorization"
	emptyHeader         = "empty auth header"
	invalidHeader       = "invalid auth header"
	userCtx             = "userID"
)

func (h *Handler) userIdentify(c *gin.Context) {
	header := c.GetHeader(authorizationHeader)
	if header == "" {
		h.log.Error(emptyHeader)
		c.AbortWithStatusJSON(http.StatusUnauthorized, models.ErrorResponse{Message: emptyHeader})
		return
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 {
		h.log.Error(invalidHeader)
		c.AbortWithStatusJSON(http.StatusUnauthorized, models.ErrorResponse{Message: invalidHeader})
		return
	}

	userID, err := h.service.Authorization.ParseToken(headerParts[1])
	if err != nil {
		h.log.Error(err)
		c.AbortWithStatusJSON(http.StatusUnauthorized, models.ErrorResponse{Message: err.Error()})
		return
	}
	c.Set(userCtx, userID)
}

func (h *Handler) getUserID(c *gin.Context) (int, error) {
	id, ok := c.Get(userCtx)
	if !ok {
		h.log.Error(emptyHeader)
		c.AbortWithStatusJSON(http.StatusInternalServerError, models.ErrorResponse{Message: notFound})
		return 0, errors.New(notFound)
	}
	idInt, ok := id.(int)
	if !ok {
		h.log.Error(emptyHeader)
		c.AbortWithStatusJSON(http.StatusInternalServerError, models.ErrorResponse{Message: notFound})
		return 0, errors.New(notFound)
	}

	return idInt, nil
}
