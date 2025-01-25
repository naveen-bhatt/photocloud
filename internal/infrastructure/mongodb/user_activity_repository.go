package mongodb

import (
	"context"
	"time"

	"photocloud/internal/domain/models"
	"photocloud/internal/domain/repositories"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const userActivityCollection = "user_activities"

type mongoUserActivityRepository struct {
	*BaseRepository
}

// NewUserActivityRepository creates a new MongoDB user activity repository
func NewUserActivityRepository(db *mongo.Database) repositories.UserActivityRepository {
	return &mongoUserActivityRepository{
		BaseRepository: NewBaseRepository(db, userActivityCollection),
	}
}

func (r *mongoUserActivityRepository) Create(ctx context.Context, activity *models.UserActivity) error {
	id, err := r.InsertOne(ctx, activity)
	if err != nil {
		return err
	}
	activity.ID = id
	return nil
}

func (r *mongoUserActivityRepository) GetByID(ctx context.Context, id primitive.ObjectID) (*models.UserActivity, error) {
	var activity models.UserActivity
	err := r.FindOne(ctx, bson.M{"_id": id}, &activity)
	if err == mongo.ErrNoDocuments {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &activity, nil
}

func (r *mongoUserActivityRepository) GetUserActivities(ctx context.Context, userID string, page, limit int) ([]models.UserActivity, error) {
	skip := (page - 1) * limit
	opts := options.Find().
		SetSkip(int64(skip)).
		SetLimit(int64(limit)).
		SetSort(bson.D{{Key: "timestamp", Value: -1}})

	var activities []models.UserActivity
	err := r.FindMany(ctx, bson.M{"user_id": userID}, opts, &activities)
	if err != nil {
		return nil, err
	}
	return activities, nil
}

func (r *mongoUserActivityRepository) GetPhotoActivities(ctx context.Context, photoID primitive.ObjectID, page, limit int) ([]models.UserActivity, error) {
	skip := (page - 1) * limit
	opts := options.Find().
		SetSkip(int64(skip)).
		SetLimit(int64(limit)).
		SetSort(bson.D{{Key: "timestamp", Value: -1}})

	var activities []models.UserActivity
	err := r.FindMany(ctx, bson.M{"photo_id": photoID}, opts, &activities)
	if err != nil {
		return nil, err
	}
	return activities, nil
}

func (r *mongoUserActivityRepository) GetActivitiesByTimeRange(ctx context.Context, startTime, endTime time.Time, page, limit int) ([]models.UserActivity, error) {
	skip := (page - 1) * limit
	opts := options.Find().
		SetSkip(int64(skip)).
		SetLimit(int64(limit)).
		SetSort(bson.D{{Key: "timestamp", Value: -1}})

	filter := bson.M{
		"timestamp": bson.M{
			"$gte": startTime,
			"$lte": endTime,
		},
	}

	var activities []models.UserActivity
	err := r.FindMany(ctx, filter, opts, &activities)
	if err != nil {
		return nil, err
	}
	return activities, nil
}

func (r *mongoUserActivityRepository) Count(ctx context.Context) (int64, error) {
	return r.CountDocuments(ctx, bson.M{})
}
