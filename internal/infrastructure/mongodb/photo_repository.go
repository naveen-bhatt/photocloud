package mongodb

import (
	"context"

	"photocloud/internal/domain/models"
	"photocloud/internal/domain/repositories"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const photoCollection = "photos"

type mongoPhotoRepository struct {
	*BaseRepository
}

// NewPhotoRepository creates a new MongoDB photo repository
func NewPhotoRepository(db *mongo.Database) repositories.PhotoRepository {
	return &mongoPhotoRepository{
		BaseRepository: NewBaseRepository(db, photoCollection),
	}
}

func (r *mongoPhotoRepository) Create(ctx context.Context, photo *models.Photo) error {
	id, err := r.InsertOne(ctx, photo)
	if err != nil {
		return err
	}
	photo.ID = id
	return nil
}

func (r *mongoPhotoRepository) GetByID(ctx context.Context, id primitive.ObjectID) (*models.Photo, error) {
	var photo models.Photo
	err := r.FindOne(ctx, bson.M{"_id": id}, &photo)
	if err == mongo.ErrNoDocuments {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &photo, nil
}

func (r *mongoPhotoRepository) Update(ctx context.Context, photo *models.Photo) error {
	result, err := r.UpdateOne(ctx, bson.M{"_id": photo.ID}, bson.M{"$set": photo})
	if err != nil {
		return err
	}
	if result.MatchedCount == 0 {
		return mongo.ErrNoDocuments
	}
	return nil
}

func (r *mongoPhotoRepository) Delete(ctx context.Context, id primitive.ObjectID) error {
	result, err := r.DeleteOne(ctx, bson.M{"_id": id})
	if err != nil {
		return err
	}
	if result.DeletedCount == 0 {
		return mongo.ErrNoDocuments
	}
	return nil
}

func (r *mongoPhotoRepository) List(ctx context.Context, page, limit int) ([]models.Photo, error) {
	skip := (page - 1) * limit
	opts := options.Find().
		SetSkip(int64(skip)).
		SetLimit(int64(limit)).
		SetSort(bson.D{{Key: "uploaded_at", Value: -1}})

	var photos []models.Photo
	err := r.FindMany(ctx, bson.M{}, opts, &photos)
	if err != nil {
		return nil, err
	}
	return photos, nil
}

func (r *mongoPhotoRepository) Count(ctx context.Context) (int64, error) {
	return r.CountDocuments(ctx, bson.M{})
}
