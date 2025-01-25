package routes

import (
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func SetupRoutes(router *gin.Engine, mongoClient *mongo.Client, s3Client *s3.Client) {
	// Health check route
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "ok",
		})
	})

	// API v1 group
	v1 := router.Group("/api/v1")
	{
		// Photo routes will be added here in phase 2
		photos := v1.Group("/photos")
		{
			// Routes will be implemented in phase 2
		}

		// Other routes will be added here in phase 2
	}
}
