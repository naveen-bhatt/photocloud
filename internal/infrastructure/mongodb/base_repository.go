package mongodb

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// BaseRepository provides generic CRUD operations for MongoDB
type BaseRepository struct {
	db         *mongo.Database
	collection string
}

// NewBaseRepository creates a new base repository
func NewBaseRepository(db *mongo.Database, collection string) *BaseRepository {
	return &BaseRepository{
		db:         db,
		collection: collection,
	}
}

// Collection returns the MongoDB collection
func (r *BaseRepository) Collection() *mongo.Collection {
	return r.db.Collection(r.collection)
}

// InsertOne inserts a single document
func (r *BaseRepository) InsertOne(ctx context.Context, document interface{}) (primitive.ObjectID, error) {
	result, err := r.Collection().InsertOne(ctx, document)
	if err != nil {
		return primitive.NilObjectID, err
	}
	return result.InsertedID.(primitive.ObjectID), nil
}

// InsertMany inserts multiple documents
func (r *BaseRepository) InsertMany(ctx context.Context, documents []interface{}) ([]primitive.ObjectID, error) {
	result, err := r.Collection().InsertMany(ctx, documents)
	if err != nil {
		return nil, err
	}

	// Convert inserted IDs to ObjectIDs
	objectIDs := make([]primitive.ObjectID, len(result.InsertedIDs))
	for i, id := range result.InsertedIDs {
		objectIDs[i] = id.(primitive.ObjectID)
	}
	return objectIDs, nil
}

// FindOne finds a single document
func (r *BaseRepository) FindOne(ctx context.Context, filter interface{}, result interface{}) error {
	return r.Collection().FindOne(ctx, filter).Decode(result)
}

// FindOneWithOptions finds a single document with options
func (r *BaseRepository) FindOneWithOptions(ctx context.Context, filter interface{}, opts *options.FindOneOptions, result interface{}) error {
	return r.Collection().FindOne(ctx, filter, opts).Decode(result)
}

// FindMany finds multiple documents with pagination
func (r *BaseRepository) FindMany(ctx context.Context, filter interface{}, opts *options.FindOptions, results interface{}) error {
	cursor, err := r.Collection().Find(ctx, filter, opts)
	if err != nil {
		return err
	}
	defer cursor.Close(ctx)
	return cursor.All(ctx, results)
}

// UpdateOne updates a single document
func (r *BaseRepository) UpdateOne(ctx context.Context, filter interface{}, update interface{}) (*mongo.UpdateResult, error) {
	return r.Collection().UpdateOne(ctx, filter, update)
}

// UpdateOneWithOptions updates a single document with options
func (r *BaseRepository) UpdateOneWithOptions(ctx context.Context, filter interface{}, update interface{}, opts *options.UpdateOptions) (*mongo.UpdateResult, error) {
	return r.Collection().UpdateOne(ctx, filter, update, opts)
}

// UpdateMany updates multiple documents
func (r *BaseRepository) UpdateMany(ctx context.Context, filter interface{}, update interface{}) (*mongo.UpdateResult, error) {
	return r.Collection().UpdateMany(ctx, filter, update)
}

// UpdateManyWithOptions updates multiple documents with options
func (r *BaseRepository) UpdateManyWithOptions(ctx context.Context, filter interface{}, update interface{}, opts *options.UpdateOptions) (*mongo.UpdateResult, error) {
	return r.Collection().UpdateMany(ctx, filter, update, opts)
}

// DeleteOne deletes a single document
func (r *BaseRepository) DeleteOne(ctx context.Context, filter interface{}) (*mongo.DeleteResult, error) {
	return r.Collection().DeleteOne(ctx, filter)
}

// DeleteMany deletes multiple documents
func (r *BaseRepository) DeleteMany(ctx context.Context, filter interface{}) (*mongo.DeleteResult, error) {
	return r.Collection().DeleteMany(ctx, filter)
}

// CountDocuments counts documents based on filter
func (r *BaseRepository) CountDocuments(ctx context.Context, filter interface{}) (int64, error) {
	return r.Collection().CountDocuments(ctx, filter)
}

// CountDocumentsWithOptions counts documents with options
func (r *BaseRepository) CountDocumentsWithOptions(ctx context.Context, filter interface{}, opts *options.CountOptions) (int64, error) {
	return r.Collection().CountDocuments(ctx, filter, opts)
}

// Distinct finds distinct values for a field
func (r *BaseRepository) Distinct(ctx context.Context, fieldName string, filter interface{}) ([]interface{}, error) {
	return r.Collection().Distinct(ctx, fieldName, filter)
}

// Aggregate performs an aggregation pipeline
func (r *BaseRepository) Aggregate(ctx context.Context, pipeline interface{}, results interface{}) error {
	cursor, err := r.Collection().Aggregate(ctx, pipeline)
	if err != nil {
		return err
	}
	defer cursor.Close(ctx)
	return cursor.All(ctx, results)
}

// AggregateWithOptions performs an aggregation pipeline with options
func (r *BaseRepository) AggregateWithOptions(ctx context.Context, pipeline interface{}, opts *options.AggregateOptions, results interface{}) error {
	cursor, err := r.Collection().Aggregate(ctx, pipeline, opts)
	if err != nil {
		return err
	}
	defer cursor.Close(ctx)
	return cursor.All(ctx, results)
}

// FindOneAndUpdate finds a document and updates it
func (r *BaseRepository) FindOneAndUpdate(ctx context.Context, filter interface{}, update interface{}, opts *options.FindOneAndUpdateOptions, result interface{}) error {
	return r.Collection().FindOneAndUpdate(ctx, filter, update, opts).Decode(result)
}

// FindOneAndDelete finds a document and deletes it
func (r *BaseRepository) FindOneAndDelete(ctx context.Context, filter interface{}, opts *options.FindOneAndDeleteOptions, result interface{}) error {
	return r.Collection().FindOneAndDelete(ctx, filter, opts).Decode(result)
}

// FindOneAndReplace finds a document and replaces it
func (r *BaseRepository) FindOneAndReplace(ctx context.Context, filter interface{}, replacement interface{}, opts *options.FindOneAndReplaceOptions, result interface{}) error {
	return r.Collection().FindOneAndReplace(ctx, filter, replacement, opts).Decode(result)
}

// BulkWrite performs multiple write operations in bulk
func (r *BaseRepository) BulkWrite(ctx context.Context, models []mongo.WriteModel) (*mongo.BulkWriteResult, error) {
	return r.Collection().BulkWrite(ctx, models)
}

// BulkWriteWithOptions performs multiple write operations in bulk with options
func (r *BaseRepository) BulkWriteWithOptions(ctx context.Context, models []mongo.WriteModel, opts *options.BulkWriteOptions) (*mongo.BulkWriteResult, error) {
	return r.Collection().BulkWrite(ctx, models, opts)
}

// CreateIndex creates an index for the collection
func (r *BaseRepository) CreateIndex(ctx context.Context, keys interface{}, opts *options.IndexOptions) (string, error) {
	indexModel := mongo.IndexModel{
		Keys:    keys,
		Options: opts,
	}
	return r.Collection().Indexes().CreateOne(ctx, indexModel)
}

// CreateIndexes creates multiple indexes for the collection
func (r *BaseRepository) CreateIndexes(ctx context.Context, models []mongo.IndexModel) ([]string, error) {
	return r.Collection().Indexes().CreateMany(ctx, models)
}

// DropIndex drops an index from the collection
func (r *BaseRepository) DropIndex(ctx context.Context, name string) error {
	_, err := r.Collection().Indexes().DropOne(ctx, name)
	return err
}

// DropAllIndexes drops all indexes from the collection
func (r *BaseRepository) DropAllIndexes(ctx context.Context) error {
	_, err := r.Collection().Indexes().DropAll(ctx)
	return err
}

// Watch creates a change stream for the collection
func (r *BaseRepository) Watch(ctx context.Context, pipeline interface{}, opts *options.ChangeStreamOptions) (*mongo.ChangeStream, error) {
	return r.Collection().Watch(ctx, pipeline, opts)
}
