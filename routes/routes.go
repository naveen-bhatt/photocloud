package routes

import (
	"os"

	"photocloud/internal/domain/services"
	"photocloud/internal/handlers"
	"photocloud/internal/infrastructure/mongodb"
	s3repo "photocloud/internal/infrastructure/s3"
	"photocloud/internal/middleware"

	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func SetupRoutes(router *gin.Engine, mongoClient *mongo.Client, s3Client *s3.Client) {
	// Initialize repositories
	photoRepo := mongodb.NewPhotoRepository(mongoClient.Database(os.Getenv("MONGODB_DATABASE")))
	storageRepo := s3repo.NewStorageRepository(s3Client, os.Getenv("AWS_S3_BUCKET"))

	// Initialize services
	photoService := services.NewPhotoService(photoRepo, storageRepo)

	// Initialize handlers
	photoHandler := handlers.NewPhotoHandler(photoService)

	// Health check route
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "ok",
		})
	})

	// API v1 group
	v1 := router.Group("/api/v1")
	{
		// Photo routes
		photos := v1.Group("/photos")
		{
			// Upload photo endpoint with file validation middleware
			photos.POST("/upload", middleware.FileValidator(), photoHandler.UploadPhoto)
		}
	}
}
