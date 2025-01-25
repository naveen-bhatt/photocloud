package dto

import (
	"mime/multipart"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// PhotoUploadRequest represents the request data for photo upload
type PhotoUploadRequest struct {
	Name        string                `form:"name" binding:"required"`
	Description string                `form:"description"`
	File        *multipart.FileHeader `form:"file" binding:"required"`
}

// PhotoResponse represents the response data for photo operations
type PhotoResponse struct {
	ID          primitive.ObjectID `json:"id"`
	Name        string             `json:"name"`
	Description string             `json:"description"`
	Size        int64              `json:"size"`
	ContentType string             `json:"content_type"`
	URL         string             `json:"url"`
	UploadedAt  time.Time          `json:"uploaded_at"`
	UpdatedAt   time.Time          `json:"updated_at"`
}
