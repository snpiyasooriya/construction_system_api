# 🧪 Testing Guide - Schedule Create API

This guide provides comprehensive testing for the Schedule Create API endpoint using Docker.

## 📋 Test Overview

We have created a complete test suite that covers:

1. **Unit Tests** - Individual component testing
2. **Integration Tests** - Full flow testing with database
3. **API Tests** - Postman-like endpoint testing
4. **Coverage Reports** - Code coverage analysis

## 🚀 Quick Start

### Run All Tests
```bash
make test
```

### Run Specific Test Types
```bash
# Unit tests only
make test-unit

# Integration tests only
make test-integration

# API tests only (Postman-like)
make test-api

# Tests with coverage report
make test-coverage
```

## 🎯 Schedule Create API Tests

### Test Cases Covered

#### ✅ Success Cases
- **Valid Schedule Creation**: Complete valid data
- **Multiple Schedules**: Creating multiple schedules for same project

#### ❌ Validation Error Cases
- **Missing Name**: Required field validation
- **Missing Description**: Required field validation
- **Missing Project ID**: Required field validation
- **Empty Name**: Empty string validation
- **Empty Description**: Empty string validation
- **Zero Project ID**: Invalid project ID validation

#### 🔧 Error Handling Cases
- **Invalid JSON**: Malformed request body
- **Empty Body**: No request body
- **Database Errors**: Connection failures

## 📊 Test Structure

### Unit Tests
```
repository/schedule_repository_test.go
├── TestCreate_Success
├── TestCreate_EmptyName
├── TestCreate_EmptyDescription
├── TestCreate_ZeroProjectID
├── TestGetByProjectID_Success
├── TestGetByProjectID_NoResults
├── TestGetCountByProjectID_Success
└── TestGetCountByProjectID_NoResults

usecase/schedule_create_usecase_test.go
├── TestExecute_Success
├── TestExecute_RepositoryError
├── TestExecute_EmptyName
├── TestExecute_ZeroProjectID
└── TestExecute_LargeProjectID

http/controllers/schedule_controller_test.go
├── TestCreateSchedule_Success
├── TestCreateSchedule_InvalidJSON
├── TestCreateSchedule_MissingRequiredFields
├── TestCreateSchedule_UseCaseError
├── TestCreateSchedule_EmptyBody
├── TestCreateSchedule_MalformedJSON
└── TestCreateSchedule_ZeroProjectID
```

### Integration Tests
```
tests/integration/schedule_integration_test.go
├── TestCreateSchedule_FullFlow_Success
├── TestCreateSchedule_ValidationError
├── TestCreateMultipleSchedules_SameProject
├── TestCreateAndRetrieveSchedules
└── TestCreateSchedule_InvalidJSON
```

### API Tests (Postman-like)
```
tests/api/schedule_api_test.go
├── Create Schedule - Success
├── Create Schedule - Missing Name
├── Create Schedule - Missing Description
├── Create Schedule - Missing Project ID
├── Create Schedule - Empty Name
├── Create Schedule - Empty Description
├── Create Schedule - Zero Project ID
├── Create Schedule - Invalid JSON
└── Create Schedule - Empty Body
```

## 🐳 Docker Test Environment

### Test Services
- **postgres-test**: Isolated test database
- **unit-tests**: Unit test runner
- **integration-tests**: Integration test runner
- **api-tests**: API test runner (Postman-like)
- **test-api**: API instance for testing

### Test Configuration
- Database: `construction_system_test`
- API Port: `8081` (to avoid conflicts)
- Test data: Isolated and cleaned between tests

## 📈 Expected Results

### Successful Test Run Output
```
🧪 Construction System API - Test Suite
=======================================

🚀 Running Unit Tests...
✅ Unit Tests completed successfully!

🚀 Running Integration Tests...
✅ Integration Tests completed successfully!

🚀 Running API Tests...
[Test 1/9] Create Schedule - Success
📝 Should successfully create a schedule with valid data
✅ PASSED

[Test 2/9] Create Schedule - Missing Name
📝 Should return validation error when name is missing
✅ PASSED

... (all tests)

📊 Test Results:
✅ Passed: 9
❌ Failed: 0
📈 Success Rate: 100.0%

🎉 All tests passed! The Schedule Create API is working correctly.
```

## 🔧 Manual Testing Commands

### Health Check
```bash
make api-health
```

### Create Schedule
```bash
make schedule-create-test
```

### Get Schedules
```bash
make schedule-get-test
```

### Custom API Test
```bash
curl -X POST http://localhost:8080/api/schedule/ \
  -H "Content-Type: application/json" \
  -d '{
    "name": "My Test Schedule",
    "description": "Testing the API",
    "project_id": 1
  }'
```

## 🐛 Troubleshooting

### Common Issues

1. **Port Conflicts**
   ```bash
   # Stop existing containers
   make clean
   ```

2. **Database Connection Issues**
   ```bash
   # Reset database
   make db-reset
   ```

3. **Test Failures**
   ```bash
   # Run with verbose output
   ./scripts/test.sh all
   ```

### Debug Mode
```bash
# Run tests with detailed output
docker-compose -f docker-compose.test.yml run --rm unit-tests go test -v ./...
```

## 📝 Test Data Examples

### Valid Request
```json
{
  "name": "Project Alpha Schedule",
  "description": "Main construction schedule for Project Alpha",
  "project_id": 1
}
```

### Expected Response
```json
{
  "id": 1,
  "name": "Project Alpha Schedule",
  "description": "Main construction schedule for Project Alpha",
  "project_id": 1,
  "created_at": "2024-01-15T10:30:00Z",
  "updated_at": "2024-01-15T10:30:00Z"
}
```

### Validation Error Response
```json
{
  "validationErrors": [
    {
      "field": "Name",
      "value": "required"
    }
  ]
}
```

## 🎯 Coverage Goals

- **Unit Tests**: >90% code coverage
- **Integration Tests**: Full API flow coverage
- **API Tests**: All endpoint scenarios covered

## 🚀 Running Tests in CI/CD

The test suite is designed to run in CI/CD pipelines:

```yaml
# Example GitHub Actions
- name: Run Tests
  run: |
    make test
    make test-coverage
```

This comprehensive test suite ensures the Schedule Create API endpoint is robust, reliable, and handles all edge cases properly.
