package repositories

import (
	"context"

	"photocloud/internal/domain/models"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// PhotoRepository defines the interface for photo data operations
type PhotoRepository interface {
	// Create creates a new photo record
	Create(ctx context.Context, photo *models.Photo) error

	// GetByID retrieves a photo by its ID
	GetByID(ctx context.Context, id primitive.ObjectID) (*models.Photo, error)

	// Update updates an existing photo
	Update(ctx context.Context, photo *models.Photo) error

	// Delete deletes a photo by its ID
	Delete(ctx context.Context, id primitive.ObjectID) error

	// List retrieves photos with pagination
	List(ctx context.Context, page, limit int) ([]models.Photo, error)

	// Count returns the total number of photos
	Count(ctx context.Context) (int64, error)
}
