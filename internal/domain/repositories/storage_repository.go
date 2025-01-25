package repositories

import (
	"context"
	"io"
)

// StorageRepository defines the interface for storage operations (S3)
type StorageRepository interface {
	// UploadFile uploads a file to storage and returns the file key
	UploadFile(ctx context.Context, key string, content io.Reader, contentType string) error

	// DownloadFile downloads a file from storage
	DownloadFile(ctx context.Context, key string) (io.ReadCloser, error)

	// DeleteFile deletes a file from storage
	DeleteFile(ctx context.Context, key string) error

	// GetFileURL gets a presigned URL for the file
	GetFileURL(ctx context.Context, key string, expiryMinutes int) (string, error)
}
