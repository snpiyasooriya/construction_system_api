# Schedule Items CRUD Implementation

This document describes the complete implementation of CRUD operations for Schedule Items in the construction system API.

## Overview

The implementation follows the Clean Architecture pattern with clear separation of concerns:
- **DTOs**: Data Transfer Objects for API communication
- **Use Cases**: Business logic implementation
- **Interfaces**: Contracts for repositories and use cases
- **Repository**: Data access layer using GORM
- **Controllers**: HTTP handlers for API endpoints

## Components Implemented

### 1. DTOs (Data Transfer Objects)
**File**: `dto/schedule_item.go`

- `ScheduleItemCreateInputDTO` - Input for creating schedule items
- `ScheduleItemCreateOutputDTO` - Output after creating schedule items
- `ScheduleItemGetByIDOutputDTO` - Output for getting schedule item by ID
- `ScheduleItemGetByScheduleOutputDTO` - Output for getting schedule items by schedule ID
- `ScheduleItemUpdateInputDTO` - Input for updating schedule items
- `ScheduleItemUpdateOutputDTO` - Output after updating schedule items
- `ScheduleItemDeleteOutputDTO` - Output after deleting schedule items

### 2. Use Case Interfaces
**Files**: 
- `usecase/schedule_item_create_usecase.go`
- `usecase/schedule_item_get_by_id_usecase.go`
- `usecase/schedule_item_get_by_schedule_usecase.go`
- `usecase/schedule_item_update_usecase.go`
- `usecase/schedule_item_delete_usecase.go`

Each interface defines a single `Execute` method following the Single Responsibility Principle.

### 3. Repository Interface
**File**: `interfaces/repository/schedule_item_repository.go`

Defines the contract for data access operations:
- `Create(scheduleItem dto.ScheduleItemCreateInputDTO) (*dto.ScheduleItemCreateOutputDTO, error)`
- `UpdateByID(scheduleItem dto.ScheduleItemUpdateInputDTO) (*dto.ScheduleItemUpdateOutputDTO, error)`
- `GetByID(id uint) (*models.ScheduleItem, error)`
- `GetByScheduleID(scheduleID uint) ([]models.ScheduleItem, error)`
- `DeleteByID(id uint) error`

### 4. Use Case Implementations
**Files**:
- `usecase/schedule_item_create_usecase_impl.go`
- `usecase/schedule_item_get_by_id_usecase_impl.go`
- `usecase/schedule_item_get_by_schedule_usecase_impl.go`
- `usecase/schedule_item_update_usecase_impl.go`
- `usecase/schedule_item_delete_usecase_impl.go`

Each implementation includes:
- Dependency injection of required repositories
- Business logic validation
- Error handling
- Data transformation between DTOs and models

### 5. Repository Implementation
**File**: `repository/schedule_item_repository.go`

GORM-based implementation with:
- Database operations using GORM ORM
- Proper error handling
- Data mapping between DTOs and models
- Preloading of related entities (Shape, Schedule)

### 6. HTTP Controller
**File**: `http/controllers/schedule_item_controller.go`

RESTful API endpoints:
- `POST /schedule-items` - Create schedule item
- `GET /schedule-items/{id}` - Get schedule item by ID
- `GET /schedules/{scheduleId}/items` - Get schedule items by schedule ID
- `PUT /schedule-items/{id}` - Update schedule item
- `DELETE /schedule-items/{id}` - Delete schedule item

### 7. Tests
**Files**:
- `usecase/schedule_item_create_usecase_test.go` - Unit tests for create use case
- `repository/schedule_item_repository_test.go` - Repository integration tests
- `tests/integration/schedule_item_integration_test.go` - Full integration tests

## API Endpoints

### Create Schedule Item
```http
POST /schedule-items
Content-Type: application/json

{
  "name": "Steel Beam",
  "shape_id": 1,
  "schedule_id": 1,
  "shape_dimensions": {"length": 10, "width": 5}
}
```

### Get Schedule Item by ID
```http
GET /schedule-items/1
```

### Get Schedule Items by Schedule ID
```http
GET /schedules/1/items
```

### Update Schedule Item
```http
PUT /schedule-items/1
Content-Type: application/json

{
  "name": "Updated Steel Beam",
  "shape_id": 2,
  "shape_dimensions": {"length": 12, "width": 6}
}
```

### Delete Schedule Item
```http
DELETE /schedule-items/1
```

## Business Logic

### Create Operation
1. Validates that the referenced schedule exists
2. Validates that the referenced shape exists
3. Creates the schedule item with provided dimensions
4. Returns the created item with timestamps

### Read Operations
- Get by ID: Retrieves single item with shape and schedule details
- Get by Schedule ID: Retrieves all items for a specific schedule

### Update Operation
1. Validates that the schedule item exists
2. Validates that the new shape exists (if changed)
3. Updates the item with new data
4. Returns the updated item

### Delete Operation
1. Validates that the schedule item exists
2. Performs soft delete using GORM
3. Returns confirmation with deletion timestamp

## Error Handling

The implementation includes comprehensive error handling:
- Validation errors for missing required fields
- Not found errors for non-existent resources
- Database errors with appropriate HTTP status codes
- Business logic errors with descriptive messages

## Testing

### Unit Tests
- Mock-based testing for use cases
- Isolated testing of business logic
- Comprehensive test coverage for success and error scenarios

### Integration Tests
- In-memory SQLite database for testing
- Full HTTP request/response testing
- End-to-end validation of the complete flow

### Repository Tests
- Database integration testing
- CRUD operation validation
- Relationship testing with preloaded entities

## Dependencies

The schedule items implementation depends on:
- **Schedule Repository**: For validating schedule existence
- **Shape Repository**: For validating shape existence
- **GORM**: For database operations
- **Gin**: For HTTP routing and handling
- **Testify**: For testing framework

## Usage Example

To integrate the schedule items functionality into the main application:

```go
// In main.go or dependency injection setup
scheduleItemRepo := repository.NewGORMScheduleItemRepository(db)
scheduleRepo := repository.NewGORMScheduleRepository(db)
shapeRepo := repository.NewGORMShapeRepository(db)

// Create use cases
scheduleItemCreateUseCase := usecase.NewScheduleItemCreateUseCaseImpl(
    scheduleItemRepo, scheduleRepo, shapeRepo)
scheduleItemGetByIDUseCase := usecase.NewScheduleItemGetByIDUseCaseImpl(scheduleItemRepo)
// ... other use cases

// Create controller
scheduleItemController := controllers.NewScheduleItemController(
    scheduleItemCreateUseCase,
    scheduleItemGetByIDUseCase,
    // ... other use cases
)

// Register routes
router.POST("/schedule-items", scheduleItemController.CreateScheduleItem)
router.GET("/schedule-items/:id", scheduleItemController.GetScheduleItemByID)
// ... other routes
```

## Future Enhancements

Potential improvements for the schedule items implementation:
1. Pagination for list operations
2. Filtering and sorting capabilities
3. Bulk operations (create/update/delete multiple items)
4. Audit logging for changes
5. Caching for frequently accessed items
6. Advanced validation rules for shape dimensions
7. File upload support for shape diagrams
8. Integration with external CAD systems
