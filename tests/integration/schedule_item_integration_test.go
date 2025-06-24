package integration

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/snpiyasooriya/construction_design_api/dto"
	"github.com/snpiyasooriya/construction_design_api/http/controllers"
	"github.com/snpiyasooriya/construction_design_api/models"
	"github.com/snpiyasooriya/construction_design_api/repository"
	"github.com/snpiyasooriya/construction_design_api/usecase"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"gorm.io/datatypes"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type ScheduleItemIntegrationTestSuite struct {
	suite.Suite
	db         *gorm.DB
	router     *gin.Engine
	controller *controllers.ScheduleItemController
}

func (suite *ScheduleItemIntegrationTestSuite) SetupSuite() {
	gin.SetMode(gin.TestMode)

	// Setup in-memory SQLite database
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	suite.Require().NoError(err)

	// Auto-migrate the schema
	err = db.AutoMigrate(&models.ScheduleItem{}, &models.Schedule{}, &models.Shape{}, &models.Project{}, &models.User{})
	suite.Require().NoError(err)

	suite.db = db

	// Setup dependencies
	scheduleItemRepo := repository.NewGORMScheduleItemRepository(db)
	scheduleRepo := repository.NewGORMScheduleRepository(db)
	shapeRepo := repository.NewGORMShapeRepository(db)

	// Create use cases
	scheduleItemCreateUseCase := usecase.NewScheduleItemCreateUseCaseImpl(scheduleItemRepo, scheduleRepo, shapeRepo)
	scheduleItemGetByIDUseCase := usecase.NewScheduleItemGetByIDUseCaseImpl(scheduleItemRepo)
	scheduleItemGetByScheduleUseCase := usecase.NewScheduleItemGetByScheduleUseCaseImpl(scheduleItemRepo, scheduleRepo)
	scheduleItemUpdateUseCase := usecase.NewScheduleItemUpdateUseCaseImpl(scheduleItemRepo, shapeRepo)
	scheduleItemDeleteUseCase := usecase.NewScheduleItemDeleteUseCaseImpl(scheduleItemRepo)

	// Create controller
	scheduleItemController := controllers.NewScheduleItemController(
		scheduleItemCreateUseCase,
		scheduleItemGetByIDUseCase,
		scheduleItemGetByScheduleUseCase,
		scheduleItemUpdateUseCase,
		scheduleItemDeleteUseCase,
	)

	suite.controller = scheduleItemController

	// Setup router
	router := gin.New()
	router.POST("/schedule-items", scheduleItemController.CreateScheduleItem)
	router.GET("/schedule-items/:id", scheduleItemController.GetScheduleItemByID)
	router.GET("/schedules/:scheduleId/items", scheduleItemController.GetScheduleItemsByScheduleID)
	router.PUT("/schedule-items/:id", scheduleItemController.UpdateScheduleItem)
	router.DELETE("/schedule-items/:id", scheduleItemController.DeleteScheduleItem)

	suite.router = router
}

func (suite *ScheduleItemIntegrationTestSuite) TearDownTest() {
	// Clean up data after each test
	suite.db.Exec("DELETE FROM schedule_items")
	suite.db.Exec("DELETE FROM schedules")
	suite.db.Exec("DELETE FROM shapes")
}

func (suite *ScheduleItemIntegrationTestSuite) TestCreateScheduleItem_Success() {
	// Setup test data
	shape := models.Shape{
		Name:       "Rectangle",
		Path:       "/shapes/rectangle",
		Dimensions: datatypes.JSON(`{"width": 100, "height": 50}`),
	}
	suite.db.Create(&shape)

	schedule := models.Schedule{
		Description: "Test Schedule",
		ScheduleID:  "1/0001",
	}
	suite.db.Create(&schedule)

	// Test data
	input := dto.ScheduleItemCreateInputDTO{
		Name:            "Test Item",
		ShapeID:         shape.ID,
		ScheduleID:      schedule.ID,
		ShapeDimensions: datatypes.JSON(`{"width": 10, "height": 20}`),
	}

	// Convert to JSON
	jsonData, err := json.Marshal(input)
	suite.Require().NoError(err)

	// Create request
	req, err := http.NewRequest("POST", "/schedule-items", bytes.NewBuffer(jsonData))
	suite.Require().NoError(err)
	req.Header.Set("Content-Type", "application/json")

	// Execute request
	w := httptest.NewRecorder()
	suite.router.ServeHTTP(w, req)

	// Assert response
	assert.Equal(suite.T(), http.StatusCreated, w.Code)

	var response dto.ScheduleItemCreateOutputDTO
	err = json.Unmarshal(w.Body.Bytes(), &response)
	suite.Require().NoError(err)

	assert.Equal(suite.T(), input.Name, response.Name)
	assert.Equal(suite.T(), input.ShapeID, response.ShapeID)
	assert.Equal(suite.T(), input.ScheduleID, response.ScheduleID)
	assert.NotZero(suite.T(), response.ID)
}

func (suite *ScheduleItemIntegrationTestSuite) TestGetScheduleItemByID_Success() {
	// Setup test data
	shape := models.Shape{
		Name:       "Rectangle",
		Path:       "/shapes/rectangle",
		Dimensions: datatypes.JSON(`{"width": 100, "height": 50}`),
	}
	suite.db.Create(&shape)

	schedule := models.Schedule{
		Description: "Test Schedule",
		ScheduleID:  "1/0001",
	}
	suite.db.Create(&schedule)

	scheduleItem := models.ScheduleItem{
		Name:            "Test Item",
		ShapeID:         shape.ID,
		ScheduleID:      schedule.ID,
		ShapeDimensions: datatypes.JSON(`{"width": 10, "height": 20}`),
	}
	suite.db.Create(&scheduleItem)

	// Create request
	req, err := http.NewRequest("GET", fmt.Sprintf("/schedule-items/%d", scheduleItem.ID), nil)
	suite.Require().NoError(err)

	// Execute request
	w := httptest.NewRecorder()
	suite.router.ServeHTTP(w, req)

	// Assert response
	assert.Equal(suite.T(), http.StatusOK, w.Code)

	var response dto.ScheduleItemGetByIDOutputDTO
	err = json.Unmarshal(w.Body.Bytes(), &response)
	suite.Require().NoError(err)

	assert.Equal(suite.T(), scheduleItem.ID, response.ID)
	assert.Equal(suite.T(), scheduleItem.Name, response.Name)
	assert.Equal(suite.T(), scheduleItem.ShapeID, response.ShapeID)
	assert.Equal(suite.T(), scheduleItem.ScheduleID, response.ScheduleID)
}

func (suite *ScheduleItemIntegrationTestSuite) TestGetScheduleItemByID_NotFound() {
	// Create request for non-existent item
	req, err := http.NewRequest("GET", "/schedule-items/999", nil)
	suite.Require().NoError(err)

	// Execute request
	w := httptest.NewRecorder()
	suite.router.ServeHTTP(w, req)

	// Assert response
	assert.Equal(suite.T(), http.StatusInternalServerError, w.Code) // GORM returns error, not 404
}

func (suite *ScheduleItemIntegrationTestSuite) TestGetScheduleItemsByScheduleID_Success() {
	// Setup test data
	shape := models.Shape{
		Name:       "Rectangle",
		Path:       "/shapes/rectangle",
		Dimensions: datatypes.JSON(`{"width": 100, "height": 50}`),
	}
	suite.db.Create(&shape)

	schedule := models.Schedule{
		Description: "Test Schedule",
		ScheduleID:  "1/0001",
	}
	suite.db.Create(&schedule)

	scheduleItem1 := models.ScheduleItem{
		Name:            "Test Item 1",
		ShapeID:         shape.ID,
		ScheduleID:      schedule.ID,
		ShapeDimensions: datatypes.JSON(`{"width": 10, "height": 20}`),
	}
	suite.db.Create(&scheduleItem1)

	scheduleItem2 := models.ScheduleItem{
		Name:            "Test Item 2",
		ShapeID:         shape.ID,
		ScheduleID:      schedule.ID,
		ShapeDimensions: datatypes.JSON(`{"width": 15, "height": 25}`),
	}
	suite.db.Create(&scheduleItem2)

	// Create request
	req, err := http.NewRequest("GET", fmt.Sprintf("/schedules/%d/items", schedule.ID), nil)
	suite.Require().NoError(err)

	// Execute request
	w := httptest.NewRecorder()
	suite.router.ServeHTTP(w, req)

	// Assert response
	assert.Equal(suite.T(), http.StatusOK, w.Code)

	var response []dto.ScheduleItemGetByScheduleOutputDTO
	err = json.Unmarshal(w.Body.Bytes(), &response)
	suite.Require().NoError(err)

	assert.Len(suite.T(), response, 2)
	assert.Equal(suite.T(), scheduleItem1.Name, response[0].Name)
	assert.Equal(suite.T(), scheduleItem2.Name, response[1].Name)
}

func (suite *ScheduleItemIntegrationTestSuite) TestDeleteScheduleItem_Success() {
	// Setup test data
	shape := models.Shape{
		Name:       "Rectangle",
		Path:       "/shapes/rectangle",
		Dimensions: datatypes.JSON(`{"width": 100, "height": 50}`),
	}
	suite.db.Create(&shape)

	schedule := models.Schedule{
		Description: "Test Schedule",
		ScheduleID:  "1/0001",
	}
	suite.db.Create(&schedule)

	scheduleItem := models.ScheduleItem{
		Name:            "Test Item",
		ShapeID:         shape.ID,
		ScheduleID:      schedule.ID,
		ShapeDimensions: datatypes.JSON(`{"width": 10, "height": 20}`),
	}
	suite.db.Create(&scheduleItem)

	// Create request
	req, err := http.NewRequest("DELETE", fmt.Sprintf("/schedule-items/%d", scheduleItem.ID), nil)
	suite.Require().NoError(err)

	// Execute request
	w := httptest.NewRecorder()
	suite.router.ServeHTTP(w, req)

	// Assert response
	assert.Equal(suite.T(), http.StatusOK, w.Code)

	var response dto.ScheduleItemDeleteOutputDTO
	err = json.Unmarshal(w.Body.Bytes(), &response)
	suite.Require().NoError(err)

	assert.Equal(suite.T(), scheduleItem.ID, response.ID)
	assert.Equal(suite.T(), "Schedule item deleted successfully", response.Message)

	// Verify deletion
	var count int64
	suite.db.Model(&models.ScheduleItem{}).Where("id = ?", scheduleItem.ID).Count(&count)
	assert.Equal(suite.T(), int64(0), count)
}

func TestScheduleItemIntegrationTestSuite(t *testing.T) {
	suite.Run(t, new(ScheduleItemIntegrationTestSuite))
}
