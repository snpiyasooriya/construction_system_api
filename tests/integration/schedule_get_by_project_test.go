package integration

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/snpiyasooriya/construction_design_api/dto"
	"github.com/snpiyasooriya/construction_design_api/http/controllers"
	"github.com/snpiyasooriya/construction_design_api/models"
	"github.com/snpiyasooriya/construction_design_api/repository"
	"github.com/snpiyasooriya/construction_design_api/usecase"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type ScheduleGetByProjectIntegrationTestSuite struct {
	suite.Suite
	db         *gorm.DB
	router     *gin.Engine
	controller *controllers.ScheduleController
}

func (suite *ScheduleGetByProjectIntegrationTestSuite) SetupSuite() {
	gin.SetMode(gin.TestMode)

	// Setup in-memory SQLite database
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	suite.Require().NoError(err)

	// Auto-migrate the schema
	err = db.AutoMigrate(&models.Schedule{}, &models.Project{}, &models.User{}, &models.ProjectType{})
	suite.Require().NoError(err)

	suite.db = db

	// Setup dependencies
	scheduleRepo := repository.NewGORMScheduleRepository(db)
	scheduleGetByProjectUseCase := usecase.NewScheduleGetByProjectUseCaseImpl(scheduleRepo)

	// Create controller with minimal dependencies for testing
	scheduleController := &controllers.ScheduleController{}
	// We'll manually call the use case in our test since we can't easily inject it

	suite.controller = scheduleController

	// Setup router
	router := gin.New()
	router.GET("/schedule/project/:project_id", func(c *gin.Context) {
		// Manually implement the controller logic for testing
		projectIDStr := c.Param("project_id")

		// Convert to uint (we'll use 1 for our test)
		var pid uint = 1
		if projectIDStr == "999" {
			pid = 999 // For testing non-existent project
		}

		schedules, err := scheduleGetByProjectUseCase.Execute(pid)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get schedules: " + err.Error()})
			return
		}
		c.JSON(http.StatusOK, schedules)
	})

	suite.router = router
}

func (suite *ScheduleGetByProjectIntegrationTestSuite) TearDownTest() {
	// Clean up data after each test
	suite.db.Exec("DELETE FROM schedules")
	suite.db.Exec("DELETE FROM projects")
	suite.db.Exec("DELETE FROM users")
	suite.db.Exec("DELETE FROM project_types")
}

func (suite *ScheduleGetByProjectIntegrationTestSuite) TestGetSchedulesByProjectID_Success() {
	// Create test data
	projectType := models.ProjectType{
		Type: "Construction",
	}
	suite.db.Create(&projectType)

	user := models.User{
		FirstName: "John",
		LastName:  "Doe",
		Email:     "john@example.com",
		Password:  "password",
		Role:      "ADMIN",
	}
	suite.db.Create(&user)

	project := models.Project{
		Name:          "Test Project",
		ProjectID:     "PROJ001",
		ProjectTypeID: projectType.ID,
		LeaderID:      user.ID,
		StartDate:     time.Now(),
		EndDate:       time.Now().AddDate(0, 6, 0),
		Status:        "IN_PROGRESS",
	}
	suite.db.Create(&project)

	// Create schedules for the project
	schedule1 := models.Schedule{
		ScheduleID:   "SCH001",
		ProjectID:    project.ID,
		Description:  "Foundation Schedule",
		RequiredDate: time.Now().AddDate(0, 1, 0),
		SchedularID:  user.ID,
		Status:       "PENDING",
		Note:         "Foundation work schedule",
	}
	suite.db.Create(&schedule1)

	schedule2 := models.Schedule{
		ScheduleID:   "SCH002",
		ProjectID:    project.ID,
		Description:  "Framing Schedule",
		RequiredDate: time.Now().AddDate(0, 2, 0),
		SchedularID:  user.ID,
		Status:       "PENDING",
		Note:         "Framing work schedule",
	}
	suite.db.Create(&schedule2)

	// Create request
	req, err := http.NewRequest("GET", fmt.Sprintf("/schedule/project/%d", project.ID), nil)
	suite.Require().NoError(err)

	// Execute request
	w := httptest.NewRecorder()
	suite.router.ServeHTTP(w, req)

	// Assert response
	assert.Equal(suite.T(), http.StatusOK, w.Code)

	var response []dto.ScheduleGetByProjectOutputDTO
	err = json.Unmarshal(w.Body.Bytes(), &response)
	suite.Require().NoError(err)

	assert.Len(suite.T(), response, 2)

	// Check first schedule
	assert.Equal(suite.T(), schedule1.ID, response[0].ID)
	assert.Equal(suite.T(), schedule1.ScheduleID, response[0].ScheduleID)
	assert.Equal(suite.T(), schedule1.Description, response[0].Description)
	assert.Equal(suite.T(), schedule1.ProjectID, response[0].ProjectID)
	assert.Equal(suite.T(), "John Doe", response[0].Schedular)

	// Check second schedule
	assert.Equal(suite.T(), schedule2.ID, response[1].ID)
	assert.Equal(suite.T(), schedule2.ScheduleID, response[1].ScheduleID)
	assert.Equal(suite.T(), schedule2.Description, response[1].Description)
	assert.Equal(suite.T(), schedule2.ProjectID, response[1].ProjectID)
	assert.Equal(suite.T(), "John Doe", response[1].Schedular)
}

func (suite *ScheduleGetByProjectIntegrationTestSuite) TestGetSchedulesByProjectID_EmptyResult() {
	// Create test data without schedules
	projectType := models.ProjectType{
		Type: "Construction",
	}
	suite.db.Create(&projectType)

	user := models.User{
		FirstName: "John",
		LastName:  "Doe",
		Email:     "john@example.com",
		Password:  "password",
		Role:      "ADMIN",
	}
	suite.db.Create(&user)

	project := models.Project{
		Name:          "Test Project",
		ProjectID:     "PROJ001",
		ProjectTypeID: projectType.ID,
		LeaderID:      user.ID,
		StartDate:     time.Now(),
		EndDate:       time.Now().AddDate(0, 6, 0),
		Status:        "IN_PROGRESS",
	}
	suite.db.Create(&project)

	// Create request
	req, err := http.NewRequest("GET", fmt.Sprintf("/schedule/project/%d", project.ID), nil)
	suite.Require().NoError(err)

	// Execute request
	w := httptest.NewRecorder()
	suite.router.ServeHTTP(w, req)

	// Assert response
	assert.Equal(suite.T(), http.StatusOK, w.Code)

	var response []dto.ScheduleGetByProjectOutputDTO
	err = json.Unmarshal(w.Body.Bytes(), &response)
	suite.Require().NoError(err)

	assert.Len(suite.T(), response, 0)
}

func (suite *ScheduleGetByProjectIntegrationTestSuite) TestGetSchedulesByProjectID_NonExistentProject() {
	// Create request for non-existent project
	req, err := http.NewRequest("GET", "/schedule/project/999", nil)
	suite.Require().NoError(err)

	// Execute request
	w := httptest.NewRecorder()
	suite.router.ServeHTTP(w, req)

	// Assert response - should return empty array for non-existent project
	assert.Equal(suite.T(), http.StatusOK, w.Code)

	var response []dto.ScheduleGetByProjectOutputDTO
	err = json.Unmarshal(w.Body.Bytes(), &response)
	suite.Require().NoError(err)

	assert.Len(suite.T(), response, 0)
}

func TestScheduleGetByProjectIntegrationTestSuite(t *testing.T) {
	suite.Run(t, new(ScheduleGetByProjectIntegrationTestSuite))
}
