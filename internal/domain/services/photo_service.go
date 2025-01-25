package services

import (
	"context"
	"fmt"
	"io"
	"path/filepath"
	"time"

	"photocloud/internal/domain/models"
	"photocloud/internal/domain/repositories"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type PhotoService interface {
	UploadPhoto(ctx context.Context, name, description string, content io.Reader, contentType string, size int64) (*models.Photo, error)
	GetPhoto(ctx context.Context, id primitive.ObjectID) (*models.Photo, error)
	GetPhotoContent(ctx context.Context, id primitive.ObjectID) (io.ReadCloser, error)
	DeletePhoto(ctx context.Context, id primitive.ObjectID) error
	ListPhotos(ctx context.Context, page, limit int) ([]models.Photo, error)
	GetPhotoURL(ctx context.Context, id primitive.ObjectID) (string, error)
}

type photoService struct {
	photoRepo   repositories.PhotoRepository
	storageRepo repositories.StorageRepository
}

func NewPhotoService(photoRepo repositories.PhotoRepository, storageRepo repositories.StorageRepository) PhotoService {
	return &photoService{
		photoRepo:   photoRepo,
		storageRepo: storageRepo,
	}
}

func (s *photoService) UploadPhoto(ctx context.Context, name, description string, content io.Reader, contentType string, size int64) (*models.Photo, error) {
	// Create a unique S3 key
	s3Key := fmt.Sprintf("photos/%s/%s%s",
		time.Now().Format("2006/01/02"),
		primitive.NewObjectID().Hex(),
		filepath.Ext(name))

	// Upload to S3
	if err := s.storageRepo.UploadFile(ctx, s3Key, content, contentType); err != nil {
		return nil, fmt.Errorf("failed to upload file to storage: %w", err)
	}

	// Create photo record
	photo := &models.Photo{
		Name:        name,
		Description: description,
		Size:        size,
		ContentType: contentType,
		S3Key:       s3Key,
		UploadedAt:  time.Now(),
		UpdatedAt:   time.Now(),
	}

	if err := s.photoRepo.Create(ctx, photo); err != nil {
		// Try to cleanup the uploaded file if database insert fails
		_ = s.storageRepo.DeleteFile(ctx, s3Key)
		return nil, fmt.Errorf("failed to create photo record: %w", err)
	}

	return photo, nil
}

func (s *photoService) GetPhoto(ctx context.Context, id primitive.ObjectID) (*models.Photo, error) {
	return s.photoRepo.GetByID(ctx, id)
}

func (s *photoService) GetPhotoContent(ctx context.Context, id primitive.ObjectID) (io.ReadCloser, error) {
	photo, err := s.photoRepo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if photo == nil {
		return nil, fmt.Errorf("photo not found")
	}

	return s.storageRepo.DownloadFile(ctx, photo.S3Key)
}

func (s *photoService) DeletePhoto(ctx context.Context, id primitive.ObjectID) error {
	photo, err := s.photoRepo.GetByID(ctx, id)
	if err != nil {
		return err
	}
	if photo == nil {
		return fmt.Errorf("photo not found")
	}

	// Delete from S3
	if err := s.storageRepo.DeleteFile(ctx, photo.S3Key); err != nil {
		return fmt.Errorf("failed to delete file from storage: %w", err)
	}

	// Delete from database
	if err := s.photoRepo.Delete(ctx, id); err != nil {
		return fmt.Errorf("failed to delete photo record: %w", err)
	}

	return nil
}

func (s *photoService) ListPhotos(ctx context.Context, page, limit int) ([]models.Photo, error) {
	return s.photoRepo.List(ctx, page, limit)
}

func (s *photoService) GetPhotoURL(ctx context.Context, id primitive.ObjectID) (string, error) {
	photo, err := s.photoRepo.GetByID(ctx, id)
	if err != nil {
		return "", err
	}
	if photo == nil {
		return "", fmt.Errorf("photo not found")
	}

	// Get a presigned URL that expires in 15 minutes
	return s.storageRepo.GetFileURL(ctx, photo.S3Key, 15)
}
