package handlers

import (
	"fmt"
	"mime/multipart"
	"net/http"

	"photocloud/internal/domain/dto"
	"photocloud/internal/domain/services"

	"github.com/gin-gonic/gin"
)

type PhotoHandler struct {
	photoService services.PhotoService
}

func NewPhotoHandler(photoService services.PhotoService) *PhotoHandler {
	return &PhotoHandler{
		photoService: photoService,
	}
}

// UploadPhoto handles photo upload requests
func (h *PhotoHandler) UploadPhoto(c *gin.Context) {
	var req dto.PhotoUploadRequest
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Invalid request data: %v", err)})
		return
	}

	// Get validated file from context
	fileHeader, exists := c.Get("validatedFile")
	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No validated file found"})
		return
	}

	// Type assert the file header
	file, ok := fileHeader.(*multipart.FileHeader)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid file data"})
		return
	}

	// Open the file
	src, err := file.Open()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to open file: %v", err)})
		return
	}
	defer src.Close()

	// Upload the photo using the service
	photo, err := h.photoService.UploadPhoto(
		c.Request.Context(),
		req.Name,
		req.Description,
		src,
		file.Header.Get("Content-Type"),
		file.Size,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to upload photo: %v", err)})
		return
	}

	// Get the photo URL
	url, err := h.photoService.GetPhotoURL(c.Request.Context(), photo.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to generate photo URL: %v", err)})
		return
	}

	// Create response
	response := dto.PhotoResponse{
		ID:          photo.ID,
		Name:        photo.Name,
		Description: photo.Description,
		Size:        photo.Size,
		ContentType: photo.ContentType,
		URL:         url,
		UploadedAt:  photo.UploadedAt,
		UpdatedAt:   photo.UpdatedAt,
	}

	c.JSON(http.StatusCreated, response)
}
