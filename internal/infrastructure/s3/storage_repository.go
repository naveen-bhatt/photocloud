package s3

import (
	"context"
	"io"
	"time"

	"photocloud/internal/domain/repositories"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type s3StorageRepository struct {
	client     *s3.Client
	bucketName string
}

// NewStorageRepository creates a new S3 storage repository
func NewStorageRepository(client *s3.Client, bucketName string) repositories.StorageRepository {
	return &s3StorageRepository{
		client:     client,
		bucketName: bucketName,
	}
}

func (r *s3StorageRepository) UploadFile(ctx context.Context, key string, content io.Reader, contentType string) error {
	_, err := r.client.PutObject(ctx, &s3.PutObjectInput{
		Bucket:      aws.String(r.bucketName),
		Key:         aws.String(key),
		Body:        content,
		ContentType: aws.String(contentType),
	})
	return err
}

func (r *s3StorageRepository) DownloadFile(ctx context.Context, key string) (io.ReadCloser, error) {
	result, err := r.client.GetObject(ctx, &s3.GetObjectInput{
		Bucket: aws.String(r.bucketName),
		Key:    aws.String(key),
	})
	if err != nil {
		return nil, err
	}
	return result.Body, nil
}

func (r *s3StorageRepository) DeleteFile(ctx context.Context, key string) error {
	_, err := r.client.DeleteObject(ctx, &s3.DeleteObjectInput{
		Bucket: aws.String(r.bucketName),
		Key:    aws.String(key),
	})
	return err
}

func (r *s3StorageRepository) GetFileURL(ctx context.Context, key string, expiryMinutes int) (string, error) {
	presignClient := s3.NewPresignClient(r.client)

	request, err := presignClient.PresignGetObject(ctx, &s3.GetObjectInput{
		Bucket: aws.String(r.bucketName),
		Key:    aws.String(key),
	}, func(opts *s3.PresignOptions) {
		opts.Expires = time.Duration(expiryMinutes) * time.Minute
	})

	if err != nil {
		return "", err
	}

	return request.URL, nil
}
