package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Photo struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name        string             `bson:"name" json:"name"`
	Description string             `bson:"description" json:"description"`
	Size        int64              `bson:"size" json:"size"`
	ContentType string             `bson:"content_type" json:"content_type"`
	S3Key       string             `bson:"s3_key" json:"s3_key"`
	UploadedAt  time.Time          `bson:"uploaded_at" json:"uploaded_at"`
	UpdatedAt   time.Time          `bson:"updated_at" json:"updated_at"`
}
