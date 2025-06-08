package integration

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/snpiyasooriya/construction_design_api/dto"
	"github.com/snpiyasooriya/construction_design_api/http/controllers"
	"github.com/snpiyasooriya/construction_design_api/http/routes"
	"github.com/snpiyasooriya/construction_design_api/models"
	"github.com/snpiyasooriya/construction_design_api/repository"
	"github.com/snpiyasooriya/construction_design_api/usecase"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type ScheduleIntegrationTestSuite struct {
	suite.Suite
	db     *gorm.DB
	router *gin.Engine
}

func (suite *ScheduleIntegrationTestSuite) SetupSuite() {
	gin.SetMode(gin.TestMode)

	// Setup in-memory SQLite database
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	suite.Require().NoError(err)

	// Auto-migrate the schema
	err = db.AutoMigrate(&models.Schedule{}, &models.Project{}, &models.User{})
	suite.Require().NoError(err)

	suite.db = db

	// Setup dependencies
	scheduleRepo := repository.NewGORMScheduleRepository(db)
	scheduleCreateUseCase := usecase.NewScheduleCreateUseCaseImpl(scheduleRepo)
	scheduleGetByProjectUseCase := usecase.NewScheduleGetByProjectUseCaseImpl(scheduleRepo)

	// Create controller
	scheduleController := controllers.NewScheduleController(scheduleGetByProjectUseCase, scheduleCreateUseCase)

	// Setup router with minimal routes for testing
	router := gin.New()
	apiRoutes := router.Group("/api")
	scheduleRoutes := apiRoutes.Group("/schedule")
	{
		scheduleRoutes.POST("/", scheduleController.CreateSchedule)
		scheduleRoutes.GET("/ByProject/", scheduleController.GetSchedulesByProjectID)
	}

	suite.router = router
}

func (suite *ScheduleIntegrationTestSuite) SetupTest() {
	// Clean up data before each test
	suite.db.Exec("DELETE FROM schedules")
	suite.db.Exec("DELETE FROM projects")
}

func (suite *ScheduleIntegrationTestSuite) TestCreateSchedule_FullFlow_Success() {
	// Arrange
	input := dto.ScheduleCreateInputDTO{
		Name:        "Integration Test Schedule",
		Description: "Integration Test Description",
		ProjectID:   1,
	}

	requestBody, _ := json.Marshal(input)
	req, _ := http.NewRequest("POST", "/api/schedule/", bytes.NewBuffer(requestBody))
	req.Header.Set("Content-Type", "application/json")

	// Act
	w := httptest.NewRecorder()
	suite.router.ServeHTTP(w, req)

	// Assert HTTP Response
	assert.Equal(suite.T(), http.StatusCreated, w.Code)

	var response dto.ScheduleCreateOutputDTO
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(suite.T(), err)
	assert.NotZero(suite.T(), response.ID)
	assert.Equal(suite.T(), "Integration Test Schedule", response.Name)
	assert.Equal(suite.T(), "Integration Test Description", response.Description)
	assert.Equal(suite.T(), uint(1), response.ProjectID)
	assert.NotZero(suite.T(), response.CreatedAt)
	assert.NotZero(suite.T(), response.UpdatedAt)

	// Assert Database State
	var savedSchedule models.Schedule
	err = suite.db.First(&savedSchedule, response.ID).Error
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), "Integration Test Schedule", savedSchedule.Name)
	assert.Equal(suite.T(), "Integration Test Description", savedSchedule.Description)
	assert.Equal(suite.T(), uint(1), savedSchedule.ProjectID)
}

func (suite *ScheduleIntegrationTestSuite) TestCreateSchedule_ValidationError() {
	// Arrange - Missing required name field
	input := dto.ScheduleCreateInputDTO{
		Description: "Integration Test Description",
		ProjectID:   1,
	}

	requestBody, _ := json.Marshal(input)
	req, _ := http.NewRequest("POST", "/api/schedule/", bytes.NewBuffer(requestBody))
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

	// Assert no data was saved to database
	var count int64
	suite.db.Model(&models.Schedule{}).Count(&count)
	assert.Equal(suite.T(), int64(0), count)
}

func (suite *ScheduleIntegrationTestSuite) TestCreateMultipleSchedules_SameProject() {
	// Arrange
	input1 := dto.ScheduleCreateInputDTO{
		Name:        "Schedule 1",
		Description: "Description 1",
		ProjectID:   1,
	}
	input2 := dto.ScheduleCreateInputDTO{
		Name:        "Schedule 2",
		Description: "Description 2",
		ProjectID:   1,
	}

	// Act - Create first schedule
	requestBody1, _ := json.Marshal(input1)
	req1, _ := http.NewRequest("POST", "/api/schedule/", bytes.NewBuffer(requestBody1))
	req1.Header.Set("Content-Type", "application/json")

	w1 := httptest.NewRecorder()
	suite.router.ServeHTTP(w1, req1)

	// Act - Create second schedule
	requestBody2, _ := json.Marshal(input2)
	req2, _ := http.NewRequest("POST", "/api/schedule/", bytes.NewBuffer(requestBody2))
	req2.Header.Set("Content-Type", "application/json")

	w2 := httptest.NewRecorder()
	suite.router.ServeHTTP(w2, req2)

	// Assert both requests succeeded
	assert.Equal(suite.T(), http.StatusCreated, w1.Code)
	assert.Equal(suite.T(), http.StatusCreated, w2.Code)

	// Assert database has both schedules
	var count int64
	suite.db.Model(&models.Schedule{}).Where("project_id = ?", 1).Count(&count)
	assert.Equal(suite.T(), int64(2), count)
}

func (suite *ScheduleIntegrationTestSuite) TestCreateAndRetrieveSchedules() {
	// Arrange - Create schedules for different projects
	schedules := []dto.ScheduleCreateInputDTO{
		{Name: "Schedule 1", Description: "Desc 1", ProjectID: 1},
		{Name: "Schedule 2", Description: "Desc 2", ProjectID: 1},
		{Name: "Schedule 3", Description: "Desc 3", ProjectID: 2},
	}

	// Act - Create schedules
	for _, schedule := range schedules {
		requestBody, _ := json.Marshal(schedule)
		req, _ := http.NewRequest("POST", "/api/schedule/", bytes.NewBuffer(requestBody))
		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()
		suite.router.ServeHTTP(w, req)
		assert.Equal(suite.T(), http.StatusCreated, w.Code)
	}

	// Act - Retrieve schedules for project 1
	req, _ := http.NewRequest("GET", "/api/schedule/ByProject/?project_id=1", nil)
	w := httptest.NewRecorder()
	suite.router.ServeHTTP(w, req)

	// Assert
	assert.Equal(suite.T(), http.StatusOK, w.Code)

	var response []dto.ScheduleGetByProjectOutputDTO
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(suite.T(), err)
	assert.Len(suite.T(), response, 2)
	assert.Equal(suite.T(), "Schedule 1", response[0].Name)
	assert.Equal(suite.T(), "Schedule 2", response[1].Name)
}

func (suite *ScheduleIntegrationTestSuite) TestCreateSchedule_InvalidJSON() {
	// Arrange
	invalidJSON := `{"name": "Test", "description": "Test", "project_id": "invalid"}`
	req, _ := http.NewRequest("POST", "/api/schedule/", bytes.NewBufferString(invalidJSON))
	req.Header.Set("Content-Type", "application/json")

	// Act
	w := httptest.NewRecorder()
	suite.router.ServeHTTP(w, req)

	// Assert
	assert.Equal(suite.T(), http.StatusBadRequest, w.Code)

	// Assert no data was saved to database
	var count int64
	suite.db.Model(&models.Schedule{}).Count(&count)
	assert.Equal(suite.T(), int64(0), count)
}

func TestScheduleIntegrationTestSuite(t *testing.T) {
	suite.Run(t, new(ScheduleIntegrationTestSuite))
}
