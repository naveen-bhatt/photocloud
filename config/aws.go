package config

import (
	"context"
	"os"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func InitializeAWS() (*s3.Client, error) {
	// Load AWS configuration
	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion(os.Getenv("AWS_REGION")),
	)
	if err != nil {
		return nil, err
	}

	// Create S3 client
	s3Client := s3.NewFromConfig(cfg)
	return s3Client, nil
}

// GetBucketName returns the configured S3 bucket name
func GetBucketName() string {
	return os.Getenv("AWS_S3_BUCKET")
}
