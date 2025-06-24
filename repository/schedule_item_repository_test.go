package repository

import (
	"testing"

	"github.com/snpiyasooriya/construction_design_api/dto"
	"github.com/snpiyasooriya/construction_design_api/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"gorm.io/datatypes"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type ScheduleItemRepositoryTestSuite struct {
	suite.Suite
	db   *gorm.DB
	repo *GORMScheduleItemRepository
}

func (suite *ScheduleItemRepositoryTestSuite) SetupSuite() {
	// Setup in-memory SQLite database
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	suite.Require().NoError(err)

	// Auto-migrate the schema
	err = db.AutoMigrate(&models.ScheduleItem{}, &models.Schedule{}, &models.Shape{}, &models.Project{}, &models.User{})
	suite.Require().NoError(err)

	suite.db = db
	suite.repo = NewGORMScheduleItemRepository(db)
}

func (suite *ScheduleItemRepositoryTestSuite) TearDownTest() {
	// Clean up data after each test
	suite.db.Exec("DELETE FROM schedule_items")
	suite.db.Exec("DELETE FROM schedules")
	suite.db.Exec("DELETE FROM shapes")
}

func (suite *ScheduleItemRepositoryTestSuite) TestCreate_Success() {
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

	// Execute
	result, err := suite.repo.Create(input)

	// Assert
	assert.NoError(suite.T(), err)
	assert.NotNil(suite.T(), result)
	assert.Equal(suite.T(), input.Name, result.Name)
	assert.Equal(suite.T(), input.ShapeID, result.ShapeID)
	assert.Equal(suite.T(), input.ScheduleID, result.ScheduleID)
	assert.NotZero(suite.T(), result.ID)
	assert.NotZero(suite.T(), result.CreatedAt)
	assert.NotZero(suite.T(), result.UpdatedAt)
}

func (suite *ScheduleItemRepositoryTestSuite) TestGetByID_Success() {
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

	// Execute
	result, err := suite.repo.GetByID(scheduleItem.ID)

	// Assert
	assert.NoError(suite.T(), err)
	assert.NotNil(suite.T(), result)
	assert.Equal(suite.T(), scheduleItem.ID, result.ID)
	assert.Equal(suite.T(), scheduleItem.Name, result.Name)
	assert.Equal(suite.T(), scheduleItem.ShapeID, result.ShapeID)
	assert.Equal(suite.T(), scheduleItem.ScheduleID, result.ScheduleID)
}

func (suite *ScheduleItemRepositoryTestSuite) TestGetByID_NotFound() {
	// Execute
	result, err := suite.repo.GetByID(999)

	// Assert
	assert.Error(suite.T(), err)
	assert.Nil(suite.T(), result)
}

func (suite *ScheduleItemRepositoryTestSuite) TestGetByScheduleID_Success() {
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

	// Execute
	result, err := suite.repo.GetByScheduleID(schedule.ID)

	// Assert
	assert.NoError(suite.T(), err)
	assert.NotNil(suite.T(), result)
	assert.Len(suite.T(), result, 2)
	assert.Equal(suite.T(), scheduleItem1.Name, result[0].Name)
	assert.Equal(suite.T(), scheduleItem2.Name, result[1].Name)
}

func (suite *ScheduleItemRepositoryTestSuite) TestUpdateByID_Success() {
	// Setup test data
	shape1 := models.Shape{
		Name:       "Rectangle",
		Path:       "/shapes/rectangle",
		Dimensions: datatypes.JSON(`{"width": 100, "height": 50}`),
	}
	suite.db.Create(&shape1)

	shape2 := models.Shape{
		Name:       "Circle",
		Path:       "/shapes/circle",
		Dimensions: datatypes.JSON(`{"radius": 25}`),
	}
	suite.db.Create(&shape2)

	schedule := models.Schedule{
		Description: "Test Schedule",
		ScheduleID:  "1/0001",
	}
	suite.db.Create(&schedule)

	scheduleItem := models.ScheduleItem{
		Name:            "Test Item",
		ShapeID:         shape1.ID,
		ScheduleID:      schedule.ID,
		ShapeDimensions: datatypes.JSON(`{"width": 10, "height": 20}`),
	}
	suite.db.Create(&scheduleItem)

	// Test data
	input := dto.ScheduleItemUpdateInputDTO{
		ID:              scheduleItem.ID,
		Name:            "Updated Item",
		ShapeID:         shape2.ID,
		ShapeDimensions: datatypes.JSON(`{"radius": 15}`),
	}

	// Execute
	result, err := suite.repo.UpdateByID(input)

	// Assert
	assert.NoError(suite.T(), err)
	assert.NotNil(suite.T(), result)
	assert.Equal(suite.T(), input.Name, result.Name)
	assert.Equal(suite.T(), input.ShapeID, result.ShapeID)
	assert.Equal(suite.T(), scheduleItem.ScheduleID, result.ScheduleID)
	assert.True(suite.T(), result.UpdatedAt.After(scheduleItem.UpdatedAt))
}

func (suite *ScheduleItemRepositoryTestSuite) TestDeleteByID_Success() {
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

	// Execute
	err := suite.repo.DeleteByID(scheduleItem.ID)

	// Assert
	assert.NoError(suite.T(), err)

	// Verify deletion
	var count int64
	suite.db.Model(&models.ScheduleItem{}).Where("id = ?", scheduleItem.ID).Count(&count)
	assert.Equal(suite.T(), int64(0), count)
}

func TestScheduleItemRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(ScheduleItemRepositoryTestSuite))
}
