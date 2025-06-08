package repository

import (
	"testing"
	"time"

	"github.com/snpiyasooriya/construction_design_api/dto"
	"github.com/snpiyasooriya/construction_design_api/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type ScheduleRepositoryTestSuite struct {
	suite.Suite
	db   *gorm.DB
	repo *GORMScheduleRepository
}

func (suite *ScheduleRepositoryTestSuite) SetupSuite() {
	// Use in-memory SQLite for testing
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	suite.Require().NoError(err)

	// Auto-migrate the schema
	err = db.AutoMigrate(&models.Schedule{}, &models.Project{})
	suite.Require().NoError(err)

	suite.db = db
	suite.repo = NewGORMScheduleRepository(db)
}

func (suite *ScheduleRepositoryTestSuite) SetupTest() {
	// Clean up data before each test
	suite.db.Exec("DELETE FROM schedules")
	suite.db.Exec("DELETE FROM projects")
}

func (suite *ScheduleRepositoryTestSuite) TestCreate_Success() {
	// Arrange
	input := dto.ScheduleCreateInputDTO{
		Name:        "Test Schedule",
		Description: "Test Description",
		ProjectID:   1,
	}

	// Act
	result, err := suite.repo.Create(input)

	// Assert
	assert.NoError(suite.T(), err)
	assert.NotNil(suite.T(), result)
	assert.Equal(suite.T(), "Test Schedule", result.Name)
	assert.Equal(suite.T(), "Test Description", result.Description)
	assert.Equal(suite.T(), uint(1), result.ProjectID)
	assert.NotZero(suite.T(), result.ID)
	assert.NotZero(suite.T(), result.CreatedAt)
	assert.NotZero(suite.T(), result.UpdatedAt)

	// Verify data was actually saved to database
	var savedSchedule models.Schedule
	err = suite.db.First(&savedSchedule, result.ID).Error
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), "Test Schedule", savedSchedule.Name)
	assert.Equal(suite.T(), "Test Description", savedSchedule.Description)
	assert.Equal(suite.T(), uint(1), savedSchedule.ProjectID)
}

func (suite *ScheduleRepositoryTestSuite) TestCreate_EmptyName() {
	// Arrange
	input := dto.ScheduleCreateInputDTO{
		Name:        "",
		Description: "Test Description",
		ProjectID:   1,
	}

	// Act
	result, err := suite.repo.Create(input)

	// Assert
	assert.NoError(suite.T(), err) // GORM allows empty strings
	assert.NotNil(suite.T(), result)
	assert.Equal(suite.T(), "", result.Name)
}

func (suite *ScheduleRepositoryTestSuite) TestCreate_EmptyDescription() {
	// Arrange
	input := dto.ScheduleCreateInputDTO{
		Name:        "Test Schedule",
		Description: "",
		ProjectID:   1,
	}

	// Act
	result, err := suite.repo.Create(input)

	// Assert
	assert.NoError(suite.T(), err)
	assert.NotNil(suite.T(), result)
	assert.Equal(suite.T(), "", result.Description)
}

func (suite *ScheduleRepositoryTestSuite) TestCreate_ZeroProjectID() {
	// Arrange
	input := dto.ScheduleCreateInputDTO{
		Name:        "Test Schedule",
		Description: "Test Description",
		ProjectID:   0,
	}

	// Act
	result, err := suite.repo.Create(input)

	// Assert
	assert.NoError(suite.T(), err)
	assert.NotNil(suite.T(), result)
	assert.Equal(suite.T(), uint(0), result.ProjectID)
}

func (suite *ScheduleRepositoryTestSuite) TestGetByProjectID_Success() {
	// Arrange - Create test schedules
	schedule1 := models.Schedule{
		Name:        "Schedule 1",
		Description: "Description 1",
		ProjectID:   1,
	}
	schedule2 := models.Schedule{
		Name:        "Schedule 2",
		Description: "Description 2",
		ProjectID:   1,
	}
	schedule3 := models.Schedule{
		Name:        "Schedule 3",
		Description: "Description 3",
		ProjectID:   2,
	}

	suite.db.Create(&schedule1)
	suite.db.Create(&schedule2)
	suite.db.Create(&schedule3)

	// Act
	result, err := suite.repo.GetByProjectID(1)

	// Assert
	assert.NoError(suite.T(), err)
	assert.Len(suite.T(), result, 2)
	assert.Equal(suite.T(), "Schedule 1", result[0].Name)
	assert.Equal(suite.T(), "Schedule 2", result[1].Name)
}

func (suite *ScheduleRepositoryTestSuite) TestGetByProjectID_NoResults() {
	// Act
	result, err := suite.repo.GetByProjectID(999)

	// Assert
	assert.NoError(suite.T(), err)
	assert.Len(suite.T(), result, 0)
}

func (suite *ScheduleRepositoryTestSuite) TestGetCountByProjectID_Success() {
	// Arrange - Create test schedules
	schedule1 := models.Schedule{ProjectID: 1}
	schedule2 := models.Schedule{ProjectID: 1}
	schedule3 := models.Schedule{ProjectID: 2}

	suite.db.Create(&schedule1)
	suite.db.Create(&schedule2)
	suite.db.Create(&schedule3)

	// Act
	count, err := suite.repo.GetCountByProjectID(1)

	// Assert
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), 2, count)
}

func (suite *ScheduleRepositoryTestSuite) TestGetCountByProjectID_NoResults() {
	// Act
	count, err := suite.repo.GetCountByProjectID(999)

	// Assert
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), 0, count)
}

func TestScheduleRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(ScheduleRepositoryTestSuite))
}
