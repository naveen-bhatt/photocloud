package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ActivityType string

const (
	ActivityTypeUpload   ActivityType = "upload"
	ActivityTypeDownload ActivityType = "download"
	ActivityTypeDelete   ActivityType = "delete"
	ActivityTypeView     ActivityType = "view"
)

type UserActivity struct {
	ID        primitive.ObjectID     `bson:"_id,omitempty" json:"id"`
	UserID    string                 `bson:"user_id" json:"user_id"`
	PhotoID   primitive.ObjectID     `bson:"photo_id" json:"photo_id"`
	Type      ActivityType           `bson:"type" json:"type"`
	Timestamp time.Time              `bson:"timestamp" json:"timestamp"`
	Metadata  map[string]interface{} `bson:"metadata,omitempty" json:"metadata,omitempty"`
}
