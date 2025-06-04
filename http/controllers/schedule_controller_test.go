package controllers

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/snpiyasooriya/construction_design_api/dto"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

// MockScheduleCreateUseCase is a mock implementation of ScheduleCreateUseCase
type MockScheduleCreateUseCase struct {
	mock.Mock
}

func (m *MockScheduleCreateUseCase) Execute(input dto.ScheduleCreateInputDTO) (*dto.ScheduleCreateOutputDTO, error) {
	args := m.Called(input)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*dto.ScheduleCreateOutputDTO), args.Error(1)
}

// MockScheduleGetByProjectUseCase is a mock implementation of ScheduleGetByProjectUseCase
type MockScheduleGetByProjectUseCase struct {
	mock.Mock
}

func (m *MockScheduleGetByProjectUseCase) Execute(projectID uint) ([]dto.ScheduleGetByProjectOutputDTO, error) {
	args := m.Called(projectID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]dto.ScheduleGetByProjectOutputDTO), args.Error(1)
}

type ScheduleControllerTestSuite struct {
	suite.Suite
	mockCreateUseCase      *MockScheduleCreateUseCase
	mockGetByProjectUseCase *MockScheduleGetByProjectUseCase
	controller             *ScheduleController
	router                 *gin.Engine
}

func (suite *ScheduleControllerTestSuite) SetupTest() {
	gin.SetMode(gin.TestMode)
	
	suite.mockCreateUseCase = new(MockScheduleCreateUseCase)
	suite.mockGetByProjectUseCase = new(MockScheduleGetByProjectUseCase)
	suite.controller = NewScheduleController(suite.mockGetByProjectUseCase, suite.mockCreateUseCase)
	
	suite.router = gin.New()
	suite.router.POST("/schedule", suite.controller.CreateSchedule)
}

func (suite *ScheduleControllerTestSuite) TestCreateSchedule_Success() {
	// Arrange
	input := dto.ScheduleCreateInputDTO{
		Name:        "Test Schedule",
		Description: "Test Description",
		ProjectID:   1,
	}

	expectedOutput := &dto.ScheduleCreateOutputDTO{
		ID:          1,
		Name:        "Test Schedule",
		Description: "Test Description",
		ProjectID:   1,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	suite.mockCreateUseCase.On("Execute", input).Return(expectedOutput, nil)

	requestBody, _ := json.Marshal(input)
	req, _ := http.NewRequest("POST", "/schedule", bytes.NewBuffer(requestBody))
	req.Header.Set("Content-Type", "application/json")

	// Act
	w := httptest.NewRecorder()
	suite.router.ServeHTTP(w, req)

	// Assert
	assert.Equal(suite.T(), http.StatusCreated, w.Code)

	var response dto.ScheduleCreateOutputDTO
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), expectedOutput.ID, response.ID)
	assert.Equal(suite.T(), expectedOutput.Name, response.Name)
	assert.Equal(suite.T(), expectedOutput.Description, response.Description)
	assert.Equal(suite.T(), expectedOutput.ProjectID, response.ProjectID)

	suite.mockCreateUseCase.AssertExpectations(suite.T())
}

func (suite *ScheduleControllerTestSuite) TestCreateSchedule_InvalidJSON() {
	// Arrange
	invalidJSON := `{"name": "Test", "description": "Test", "project_id": "invalid"}`
	req, _ := http.NewRequest("POST", "/schedule", bytes.NewBufferString(invalidJSON))
	req.Header.Set("Content-Type", "application/json")

	// Act
	w := httptest.NewRecorder()
	suite.router.ServeHTTP(w, req)

	// Assert
	assert.Equal(suite.T(), http.StatusBadRequest, w.Code)

	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(suite.T(), err)
	assert.Contains(suite.T(), response, "error")
}

func (suite *ScheduleControllerTestSuite) TestCreateSchedule_MissingRequiredFields() {
	// Arrange
	input := dto.ScheduleCreateInputDTO{
		// Missing required fields
		Description: "Test Description",
		ProjectID:   1,
	}

	requestBody, _ := json.Marshal(input)
	req, _ := http.NewRequest("POST", "/schedule", bytes.NewBuffer(requestBody))
	req.Header.Set("Content-Type", "application/json")

	// Act
	w := httptest.NewRecorder()
	suite.router.ServeHTTP(w, req)

	// Assert
	assert.Equal(suite.T(), http.StatusBadRequest, w.Code)

	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(suite.T(), err)
	assert.Contains(suite.T(), response, "validationErrors")
}

func (suite *ScheduleControllerTestSuite) TestCreateSchedule_UseCaseError() {
	// Arrange
	input := dto.ScheduleCreateInputDTO{
		Name:        "Test Schedule",
		Description: "Test Description",
		ProjectID:   1,
	}

	expectedError := errors.New("database connection failed")
	suite.mockCreateUseCase.On("Execute", input).Return(nil, expectedError)

	requestBody, _ := json.Marshal(input)
	req, _ := http.NewRequest("POST", "/schedule", bytes.NewBuffer(requestBody))
	req.Header.Set("Content-Type", "application/json")

	// Act
	w := httptest.NewRecorder()
	suite.router.ServeHTTP(w, req)

	// Assert
	assert.Equal(suite.T(), http.StatusInternalServerError, w.Code)

	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(suite.T(), err)
	assert.Contains(suite.T(), response, "error")
	assert.Contains(suite.T(), response["error"].(string), "Failed to create schedule")

	suite.mockCreateUseCase.AssertExpectations(suite.T())
}

func (suite *ScheduleControllerTestSuite) TestCreateSchedule_EmptyBody() {
	// Arrange
	req, _ := http.NewRequest("POST", "/schedule", bytes.NewBuffer([]byte{}))
	req.Header.Set("Content-Type", "application/json")

	// Act
	w := httptest.NewRecorder()
	suite.router.ServeHTTP(w, req)

	// Assert
	assert.Equal(suite.T(), http.StatusBadRequest, w.Code)
}

func (suite *ScheduleControllerTestSuite) TestCreateSchedule_MalformedJSON() {
	// Arrange
	malformedJSON := `{"name": "Test", "description": "Test"`
	req, _ := http.NewRequest("POST", "/schedule", bytes.NewBufferString(malformedJSON))
	req.Header.Set("Content-Type", "application/json")

	// Act
	w := httptest.NewRecorder()
	suite.router.ServeHTTP(w, req)

	// Assert
	assert.Equal(suite.T(), http.StatusBadRequest, w.Code)
}

func (suite *ScheduleControllerTestSuite) TestCreateSchedule_ZeroProjectID() {
	// Arrange
	input := dto.ScheduleCreateInputDTO{
		Name:        "Test Schedule",
		Description: "Test Description",
		ProjectID:   0,
	}

	requestBody, _ := json.Marshal(input)
	req, _ := http.NewRequest("POST", "/schedule", bytes.NewBuffer(requestBody))
	req.Header.Set("Content-Type", "application/json")

	// Act
	w := httptest.NewRecorder()
	suite.router.ServeHTTP(w, req)

	// Assert
	assert.Equal(suite.T(), http.StatusBadRequest, w.Code)

	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(suite.T(), err)
	assert.Contains(suite.T(), response, "validationErrors")
}

func TestScheduleControllerTestSuite(t *testing.T) {
	suite.Run(t, new(ScheduleControllerTestSuite))
}
