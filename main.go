package main

import (
	"context"
	"log"
	"os"

	"photocloud/config"
	"photocloud/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	// Initialize MongoDB connection
	mongoClient, err := config.ConnectMongoDB()
	if err != nil {
		log.Fatal("Error connecting to MongoDB:", err)
	}
	defer mongoClient.Disconnect(context.Background())

	// Initialize AWS S3 client
	s3Client, err := config.InitializeAWS()
	if err != nil {
		log.Fatal("Error initializing AWS:", err)
	}

	// Initialize Gin router
	router := gin.Default()

	// Setup routes
	routes.SetupRoutes(router, mongoClient, s3Client)

	// Start server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	if err := router.Run(":" + port); err != nil {
		log.Fatal("Error starting server:", err)
	}
}
