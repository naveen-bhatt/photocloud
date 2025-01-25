# PhotoCloud

PhotoCloud is a photo storage and management platform built with Go, using the Gin framework, MongoDB for data storage, and AWS S3 for photo storage.

## Prerequisites

- Go 1.21 or later
- MongoDB
- AWS Account with S3 access
- Docker (optional)

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

Update the following variables in `.env`:

- `MONGODB_URI`: Your MongoDB connection string
- `MONGODB_DATABASE`: Database name
- `AWS_REGION`: Your AWS region
- `AWS_ACCESS_KEY_ID`: Your AWS access key
- `AWS_SECRET_ACCESS_KEY`: Your AWS secret key
- `AWS_S3_BUCKET`: Your S3 bucket name
- `PORT`: Server port (default: 8080)

4. Run the application:

```bash
go run main.go
```

The server will start on `http://localhost:8080`

## Project Structure

```
photocloud/
├── config/
│   ├── aws.go         # AWS S3 configuration
│   └── mongodb.go     # MongoDB configuration
├── routes/
│   └── routes.go      # API routes
├── .env               # Environment variables
├── go.mod            # Go modules file
├── main.go           # Application entry point
└── README.md         # This file
```

## API Endpoints

### Current Endpoints

- `GET /health` - Health check endpoint

More endpoints will be added in Phase 2 of the project.

## Development Phases

### Phase 1 (Current)

- Basic project setup
- MongoDB connection
- AWS S3 integration
- Project structure and configuration

### Phase 2 (Upcoming)

- Photo upload/download APIs
- User management
- Additional features (TBD)
