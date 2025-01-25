# PhotoCloud

PhotoCloud is a photo storage and management platform built with Go, using the Gin framework, MongoDB for data storage, and AWS S3 for photo storage.

## Features

- Photo upload with validation
- Secure file storage in AWS S3
- Metadata storage in MongoDB
- File type and size validation
- Presigned URLs for photo access

## Prerequisites

- Go 1.21 or later
- MongoDB
- AWS Account with S3 access
- Docker (optional)

## Project Structure

```
photocloud/
├── internal/
│   ├── domain/
│   │   ├── dto/           # Data Transfer Objects
│   │   ├── models/        # Domain models
│   │   ├── repositories/  # Repository interfaces
│   │   └── services/      # Business logic services
│   ├── handlers/          # HTTP request handlers
│   ├── middleware/        # HTTP middleware
│   └── infrastructure/    # External services implementation
│       ├── mongodb/       # MongoDB repositories
│       └── s3/           # AWS S3 storage
├── routes/               # API routes
├── config/              # Application configuration
├── .env                # Environment variables (not in git)
├── .env.example        # Environment variables template
└── main.go            # Application entry point
```

## Setup

1. Clone the repository:

```bash
git clone <repository-url>
cd photocloud
```

2. Install dependencies:

```bash
go mod tidy
```

3. Configure environment variables:
   Copy the `.env.example` file to `.env` and update the values:

```bash
cp .env.example .env
```

Required environment variables:

```env
# MongoDB Configuration
MONGODB_URI=mongodb://localhost:27017
MONGODB_DATABASE=photocloud

# AWS Configuration
AWS_REGION=ap-south-1
AWS_ACCESS_KEY_ID=your_access_key_here
AWS_SECRET_ACCESS_KEY=your_secret_key_here
AWS_S3_BUCKET=your_bucket_name

# Server Configuration
PORT=8080

# Optional Configurations
LOG_LEVEL=debug
ENVIRONMENT=development
MAX_UPLOAD_SIZE=10485760  # 10MB in bytes
ALLOWED_ORIGINS=http://localhost:3000,http://localhost:8080
```

4. Run the application:

```bash
go run main.go
```

The server will start on `http://localhost:8080`

## API Documentation

### Current Endpoints

#### Health Check

- `GET /health`
  - Response: `{"status": "ok"}`

#### Upload Photo

- `POST /api/v1/photos/upload`
  - Content-Type: `multipart/form-data`
  - Request Body:
    ```
    name: string (required)
    description: string (optional)
    file: file (required, image file)
    ```
  - Response:
    ```json
    {
      "id": "photo_id",
      "name": "photo_name",
      "description": "photo_description",
      "size": 1234567,
      "content_type": "image/jpeg",
      "url": "presigned_s3_url",
      "uploaded_at": "2024-01-25T12:00:00Z",
      "updated_at": "2024-01-25T12:00:00Z"
    }
    ```

### File Upload Restrictions

- Supported file types: JPEG, PNG, GIF, WebP
- Maximum file size: 10MB (configurable via MAX_UPLOAD_SIZE)
- File validation includes:
  - File size check
  - MIME type validation
  - File extension validation

## Development Phases

### Phase 1 (Completed)

- Basic project setup
- MongoDB connection
- AWS S3 integration
- Project structure and configuration

### Phase 2 (Current)

- Photo upload API with validation
- File storage in S3
- Metadata storage in MongoDB
- Presigned URL generation

### Future Phases

- User management
- Photo galleries
- Sharing capabilities
- Advanced photo management features

## Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

This project is licensed under the MIT License - see the LICENSE file for details.
