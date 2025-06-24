# Schedule Get By Project ID - Issue Resolution

## Problem Identified

The "get schedules by project ID" endpoint had several issues that prevented it from working as intended:

### 1. **Route vs Swagger Documentation Mismatch**
- **Route Definition**: `/schedule/project/:project_id` (path parameter)
- **Swagger Annotation**: Documented as query parameter and wrong route path
- **Controller Implementation**: Correctly using `c.Param("project_id")`

### 2. **Swagger Documentation Issues**
- Incorrect route path in `@Router` annotation
- Parameter type mismatch (query vs path parameter)

## Issues Fixed

### ✅ **1. Corrected Swagger Annotations**

**Before:**
```go
// @Param project_id query int true "Project ID"
// @Router /api/schedule/ByProject/ [get]
```

**After:**
```go
// @Param project_id path int true "Project ID"
// @Router /api/schedule/project/{project_id} [get]
```

### ✅ **2. Route Implementation Verification**
- Confirmed route is correctly defined: `/schedule/project/:project_id`
- Controller correctly extracts parameter: `c.Param("project_id")`
- Repository and use case implementations are working correctly

### ✅ **3. Integration Testing**
- Created comprehensive integration tests
- Verified functionality with real database operations
- Tested edge cases (empty results, non-existent projects)

## Current Working Implementation

### **Route Definition**
```go
// In http/routes/routes.go
scheduleRoutes.GET("/project/:project_id", scheduleController.GetSchedulesByProjectID)
```

### **Controller Method**
```go
// In http/controllers/schedule_controller.go
func (sc *ScheduleController) GetSchedulesByProjectID(c *gin.Context) {
    projectID, err := strconv.ParseUint(c.Param("project_id"), 10, 64)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid project_id"})
        return
    }
    schedules, err := sc.scheduleGetByProjectUseCase.Execute(uint(projectID))
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get schedules"})
        return
    }
    c.JSON(http.StatusOK, schedules)
}
```

### **Use Case Implementation**
```go
// In usecase/schedule_get_by_project_usecase_impl.go
func (s *ScheduleGetByProjectUseCaseImpl) Execute(projectID uint) ([]dto.ScheduleGetByProjectOutputDTO, error) {
    schedules, err := s.scheduleRepository.GetByProjectID(projectID)
    if err != nil {
        return nil, err
    }
    // Maps models to DTOs with proper user name formatting
    // Handles nullable ReviewerID properly
    return scheduleGetByProjectOutputDTOs, nil
}
```

### **Repository Implementation**
```go
// In repository/schedule_repository.go
func (g *GORMScheduleRepository) GetByProjectID(projectID uint) ([]models.Schedule, error) {
    var schedules []models.Schedule
    if err := g.db.Where("project_id = ?", projectID).
        Preload("Schedular").
        Preload("Reviewer").
        Preload("Project").
        Find(&schedules).Error; err != nil {
        return nil, err
    }
    return schedules, nil
}
```

## How to Use the Endpoint

### **API Endpoint**
```
GET /api/schedule/project/{project_id}
```

### **Example Request**
```bash
curl -X GET "http://localhost:8080/api/schedule/project/1" \
     -H "Authorization: Bearer <your-jwt-token>" \
     -H "Content-Type: application/json"
```

### **Example Response**
```json
[
  {
    "id": 1,
    "schedule_id": "SCH001",
    "project_id": 1,
    "description": "Foundation Schedule",
    "required_date": "2025-07-25T01:53:22.915434019+05:30",
    "schedular_id": 1,
    "schedular": "John Doe",
    "reviewer_id": 0,
    "reviewer": "",
    "status": "PENDING",
    "note": "Foundation work schedule",
    "created_at": "2025-06-25T01:53:22.91545868+05:30",
    "updated_at": "2025-06-25T01:53:22.91545868+05:30"
  },
  {
    "id": 2,
    "schedule_id": "SCH002",
    "project_id": 1,
    "description": "Framing Schedule",
    "required_date": "2025-08-25T01:53:22.91551423+05:30",
    "schedular_id": 1,
    "schedular": "John Doe",
    "reviewer_id": 0,
    "reviewer": "",
    "status": "PENDING",
    "note": "Framing work schedule",
    "created_at": "2025-06-25T01:53:22.915526699+05:30",
    "updated_at": "2025-06-25T01:53:22.915526699+05:30"
  }
]
```

## Response Details

### **Success Response (200 OK)**
- Returns array of `ScheduleGetByProjectOutputDTO` objects
- Empty array `[]` if no schedules found for the project
- Includes full schedule details with user names

### **Error Responses**
- **400 Bad Request**: Invalid project_id format
- **500 Internal Server Error**: Database or server error

### **DTO Structure**
```go
type ScheduleGetByProjectOutputDTO struct {
    ID           uint                     `json:"id"`
    ScheduleID   string                   `json:"schedule_id"`
    ProjectID    uint                     `json:"project_id"`
    Description  string                   `json:"description"`
    RequiredDate time.Time                `json:"required_date"`
    SchedularID  uint                     `json:"schedular_id"`
    Schedular    string                   `json:"schedular"`        // Full name
    ReviewerID   uint                     `json:"reviewer_id"`      // 0 if no reviewer
    Reviewer     string                   `json:"reviewer"`         // Empty if no reviewer
    Status       constants.ScheduleStatus `json:"status"`
    Note         string                   `json:"note"`
    CreatedAt    time.Time                `json:"created_at"`
    UpdatedAt    time.Time                `json:"updated_at"`
}
```

## Testing

### **Integration Tests**
- ✅ Success case with multiple schedules
- ✅ Empty result case (project with no schedules)
- ✅ Non-existent project case
- ✅ Proper DTO mapping verification
- ✅ User name formatting verification

### **Test File Location**
```
tests/integration/schedule_get_by_project_test.go
```

### **Run Tests**
```bash
go test -run TestScheduleGetByProjectIntegrationTestSuite ./tests/integration/schedule_get_by_project_test.go -v
```

## Swagger Documentation

The endpoint is now properly documented in the Swagger/OpenAPI specification:

- **Path**: `/api/schedule/project/{project_id}`
- **Method**: GET
- **Parameter**: `project_id` (path parameter, integer)
- **Security**: Bearer token required
- **Response**: Array of schedule objects

### **Access Swagger UI**
```
http://localhost:8080/swagger/index.html
```

## Summary

The "get schedules by project ID" endpoint is now **fully functional** and properly documented. The issues were primarily related to:

1. **Documentation mismatches** - Fixed swagger annotations
2. **Route documentation** - Corrected parameter types and paths
3. **Testing gaps** - Added comprehensive integration tests

The underlying business logic, repository implementation, and use case were already working correctly. The endpoint now works as intended and is properly documented for API consumers.

## Next Steps

1. **Test the endpoint** using the Swagger UI or API client
2. **Verify authentication** works with JWT tokens
3. **Check edge cases** with your actual data
4. **Update any client applications** to use the correct endpoint path

The endpoint is production-ready and follows the established patterns in the codebase.
