package repositories

import (
	"context"
	"time"

	"photocloud/internal/domain/models"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserActivityRepository interface {
	// Create creates a new activity record
	Create(ctx context.Context, activity *models.UserActivity) error

	// GetByID retrieves an activity by its ID
	GetByID(ctx context.Context, id primitive.ObjectID) (*models.UserActivity, error)

	// GetUserActivities retrieves activities for a specific user with pagination
	GetUserActivities(ctx context.Context, userID string, page, limit int) ([]models.UserActivity, error)

	// GetPhotoActivities retrieves activities for a specific photo with pagination
	GetPhotoActivities(ctx context.Context, photoID primitive.ObjectID, page, limit int) ([]models.UserActivity, error)

	// GetActivitiesByTimeRange retrieves activities within a time range
	GetActivitiesByTimeRange(ctx context.Context, startTime, endTime time.Time, page, limit int) ([]models.UserActivity, error)

	// Count returns the total number of activities
	Count(ctx context.Context) (int64, error)
}
