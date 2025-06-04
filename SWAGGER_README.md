# Swagger API Documentation

This document explains how to access and use the Swagger API documentation for the Construction System API.

## Overview

The Construction System API now includes comprehensive Swagger/OpenAPI documentation that provides:

- Interactive API documentation
- Request/response schemas
- Authentication requirements
- Example requests and responses
- Try-it-out functionality

## Accessing Swagger UI

Once the application is running, you can access the Swagger UI at:

```
http://localhost:8080/swagger/index.html
```

## API Endpoints Documented

The following endpoints are fully documented:

### Authentication
- `POST /api/login` - User authentication

### Health Check
- `GET /api/ping` - Simple connectivity test
- `GET /api/health` - Service health status

### Users
- `POST /api/users` - Create a new user
- `GET /api/users` - Get all users
- `GET /api/users/{id}` - Get user by ID
- `PUT /api/users/{id}` - Update user
- `DELETE /api/users/{id}` - Delete user

### Projects
- `POST /api/project` - Create a new project
- `GET /api/project` - Get all projects
- `GET /api/project/{id}` - Get project by ID
- `PUT /api/project/{id}` - Update project
- `DELETE /api/project/{id}` - Delete project
- `POST /api/project/{id}/users` - Add user to project

### Project Types
- `POST /api/project-type` - Create a new project type
- `GET /api/project-type` - Get all project types
- `GET /api/project-type/{id}` - Get project type by ID
- `PUT /api/project-type/{id}` - Update project type
- `DELETE /api/project-type/{id}` - Delete project type

### Schedules
- `GET /api/schedule/ByProject/` - Get schedules by project ID
- `POST /api/schedule` - Create a new schedule

### Shapes
- `POST /api/shapes` - Create a new shape
- `GET /api/shapes` - Get all shapes
- `GET /api/shapes/{id}` - Get shape by ID
- `DELETE /api/shapes/{id}` - Delete shape

## Authentication

Most endpoints require Bearer token authentication. To use protected endpoints:

1. First, authenticate using the `/api/login` endpoint
2. Copy the JWT token from the response
3. In Swagger UI, click the "Authorize" button
4. Enter `Bearer <your-token>` in the Authorization field

## Data Models

All request and response data models are fully documented with:
- Field types and constraints
- Required vs optional fields
- Validation rules
- Example values

## Running the Application

To start the application with Swagger documentation:

```bash
# Build the application
go build -o main .

# Run the application
./main
```

Or using Docker:

```bash
# Build and run with Docker Compose
docker-compose up --build
```

The application will start on port 8080 by default. You can then access:
- API endpoints: `http://localhost:8080/api/*`
- Swagger UI: `http://localhost:8080/swagger/index.html`

## Regenerating Documentation

If you make changes to the API endpoints or models, regenerate the Swagger docs:

```bash
# Install swag CLI (if not already installed)
go install github.com/swaggo/swag/cmd/swag@latest

# Generate documentation
swag init
```

## Files Generated

The Swagger generation creates the following files:
- `docs/docs.go` - Go documentation package
- `docs/swagger.json` - OpenAPI JSON specification
- `docs/swagger.yaml` - OpenAPI YAML specification

## Notes

- The Swagger UI provides a complete interactive interface for testing the API
- All validation rules from the Go structs are reflected in the documentation
- Security requirements are clearly marked for protected endpoints
- Response schemas include all possible HTTP status codes
