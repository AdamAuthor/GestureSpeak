package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"os"
	"time"
	"todoApp/internal/models"
)

func (h *Handler) uploadVideo(c *gin.Context) {
	//userId, err := h.getUserID(c)
	//if err != nil {
	//	newErrorResponse(c, http.StatusInternalServerError, err.Error())
	//	return
	//}

	userId := 1

	file, err := c.FormFile("video")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Video file is required"})
		return
	}

	dst := fmt.Sprintf("uploads/%s", file.Filename)
	if err := c.SaveUploadedFile(file, dst); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save video file"})
		return
	}

	videoContent, err := file.Open()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to open video file"})
		return
	}
	defer videoContent.Close()

	videoData, err := io.ReadAll(videoContent)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read video file"})
		return
	}

	err = h.s3Service.UploadVideoFileToR2(c, dst, videoData)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save video file"})
		return
	}

	video := models.VideoFile{
		UserID:           userId,
		S3Key:            h.s3Service.GetFileURL(dst),
		UploadDate:       time.Now(),
		ProcessingStatus: "pending",
	}

	err = h.service.UploadVideo.CreateVideoFile(video)

	defer func() {
		if err := deleteLocalFile(dst); err != nil {
			fmt.Println("Failed to delete local file:", err)
		}
	}()

	c.JSON(http.StatusOK, gin.H{"message": "Video uploaded successfully"})
}

func deleteLocalFile(filePath string) error {
	if err := os.Remove(filePath); err != nil {
		return err
	}
	return nil
}
