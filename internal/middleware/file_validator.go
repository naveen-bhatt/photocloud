package middleware

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

var (
	defaultMaxSize   int64 = 10 * 1024 * 1024 // 10MB
	allowedMimeTypes       = map[string]bool{
		"image/jpeg": true,
		"image/png":  true,
		"image/gif":  true,
		"image/webp": true,
	}
)

// FileValidator validates uploaded files
func FileValidator() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get max file size from env or use default
		maxSize := defaultMaxSize
		if sizeStr := os.Getenv("MAX_UPLOAD_SIZE"); sizeStr != "" {
			if size, err := strconv.ParseInt(sizeStr, 10, 64); err == nil {
				maxSize = size
			}
		}

		// Parse multipart form with size limit
		if err := c.Request.ParseMultipartForm(maxSize); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Failed to parse form: %v. Make sure the content-type is multipart/form-data", err)})
			c.Abort()
			return
		}

		// Get file from form
		file, fileHeader, err := c.Request.FormFile("file")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "No file uploaded. The file must be sent as a form field named 'file'",
				"required_fields": map[string]string{
					"file":        "file field containing the image (required)",
					"name":        "name of the photo (required)",
					"description": "description of the photo (optional)",
				},
			})
			c.Abort()
			return
		}
		defer file.Close()

		// Check file size
		if fileHeader.Size > maxSize {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": fmt.Sprintf("File size %d bytes exceeds maximum limit of %d bytes", fileHeader.Size, maxSize),
			})
			c.Abort()
			return
		}

		// Get content type
		contentType := fileHeader.Header.Get("Content-Type")
		if contentType == "" {
			// Try to detect content type from file extension
			ext := strings.ToLower(filepath.Ext(fileHeader.Filename))
			switch ext {
			case ".jpg", ".jpeg":
				contentType = "image/jpeg"
			case ".png":
				contentType = "image/png"
			case ".gif":
				contentType = "image/gif"
			case ".webp":
				contentType = "image/webp"
			}
		}

		// Validate content type
		if !allowedMimeTypes[contentType] {
			c.JSON(http.StatusBadRequest, gin.H{
				"error":         fmt.Sprintf("Invalid file type %s. Allowed types: JPEG, PNG, GIF, WebP", contentType),
				"allowed_types": []string{"image/jpeg", "image/png", "image/gif", "image/webp"},
			})
			c.Abort()
			return
		}

		// Check file extension
		ext := strings.ToLower(filepath.Ext(fileHeader.Filename))
		validExt := false
		switch ext {
		case ".jpg", ".jpeg", ".png", ".gif", ".webp":
			validExt = true
		}

		if !validExt {
			c.JSON(http.StatusBadRequest, gin.H{
				"error":              fmt.Sprintf("Invalid file extension %s. Allowed extensions: .jpg, .jpeg, .png, .gif, .webp", ext),
				"allowed_extensions": []string{".jpg", ".jpeg", ".png", ".gif", ".webp"},
			})
			c.Abort()
			return
		}

		// Store validated file header in context for later use
		c.Set("validatedFile", fileHeader)
		c.Next()
	}
}
