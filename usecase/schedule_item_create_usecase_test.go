package usecase

import (
	"errors"
	"testing"
	"time"

	"github.com/snpiyasooriya/construction_design_api/dto"
	"github.com/snpiyasooriya/construction_design_api/entities"
	"github.com/snpiyasooriya/construction_design_api/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"gorm.io/datatypes"
)

// MockScheduleItemRepository is a mock implementation of ScheduleItemRepository
type MockScheduleItemRepository struct {
	mock.Mock
}

func (m *MockScheduleItemRepository) Create(scheduleItem dto.ScheduleItemCreateInputDTO) (*dto.ScheduleItemCreateOutputDTO, error) {
	args := m.Called(scheduleItem)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*dto.ScheduleItemCreateOutputDTO), args.Error(1)
}

func (m *MockScheduleItemRepository) UpdateByID(scheduleItem dto.ScheduleItemUpdateInputDTO) (*dto.ScheduleItemUpdateOutputDTO, error) {
	args := m.Called(scheduleItem)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*dto.ScheduleItemUpdateOutputDTO), args.Error(1)
}

func (m *MockScheduleItemRepository) GetByID(id uint) (*models.ScheduleItem, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.ScheduleItem), args.Error(1)
}

func (m *MockScheduleItemRepository) GetByScheduleID(scheduleID uint) ([]models.ScheduleItem, error) {
	args := m.Called(scheduleID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]models.ScheduleItem), args.Error(1)
}

func (m *MockScheduleItemRepository) DeleteByID(id uint) error {
	args := m.Called(id)
	return args.Error(0)
}

// MockShapeRepository is a mock implementation of ShapeRepositoryInterFace
type MockShapeRepository struct {
	mock.Mock
}

func (m *MockShapeRepository) Create(shape entities.Shape) (*entities.Shape, error) {
	args := m.Called(shape)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entities.Shape), args.Error(1)
}

func (m *MockShapeRepository) GetByID(id uint) (*entities.Shape, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entities.Shape), args.Error(1)
}

func (m *MockShapeRepository) Get() ([]entities.Shape, error) {
	args := m.Called()
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]entities.Shape), args.Error(1)
}

func (m *MockShapeRepository) DeleteByID(id uint) error {
	args := m.Called(id)
	return args.Error(0)
}

// MockScheduleRepositoryForScheduleItem is a mock implementation of ScheduleRepository for schedule item tests
type MockScheduleRepositoryForScheduleItem struct {
	mock.Mock
}

func (m *MockScheduleRepositoryForScheduleItem) Create(schedule *dto.ScheduleCreateInputDTO) error {
	args := m.Called(schedule)
	return args.Error(0)
}

func (m *MockScheduleRepositoryForScheduleItem) UpdateByID(schedule models.Schedule) (*models.Schedule, error) {
	args := m.Called(schedule)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.Schedule), args.Error(1)
}

func (m *MockScheduleRepositoryForScheduleItem) GetByID(id uint) (*models.Schedule, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.Schedule), args.Error(1)
}

func (m *MockScheduleRepositoryForScheduleItem) Get() ([]models.Schedule, error) {
	args := m.Called()
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]models.Schedule), args.Error(1)
}

func (m *MockScheduleRepositoryForScheduleItem) DeleteByID(id uint) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockScheduleRepositoryForScheduleItem) GetCountByProjectID(projectID uint) (int, error) {
	args := m.Called(projectID)
	return args.Int(0), args.Error(1)
}

func (m *MockScheduleRepositoryForScheduleItem) GetByProjectID(projectID uint) ([]models.Schedule, error) {
	args := m.Called(projectID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]models.Schedule), args.Error(1)
}

type ScheduleItemCreateUseCaseTestSuite struct {
	suite.Suite
	mockScheduleItemRepo *MockScheduleItemRepository
	mockScheduleRepo     *MockScheduleRepositoryForScheduleItem
	mockShapeRepo        *MockShapeRepository
	useCase              *ScheduleItemCreateUseCaseImpl
}

func (suite *ScheduleItemCreateUseCaseTestSuite) SetupTest() {
	suite.mockScheduleItemRepo = new(MockScheduleItemRepository)
	suite.mockScheduleRepo = new(MockScheduleRepositoryForScheduleItem)
	suite.mockShapeRepo = new(MockShapeRepository)
	suite.useCase = NewScheduleItemCreateUseCaseImpl(
		suite.mockScheduleItemRepo,
		suite.mockScheduleRepo,
		suite.mockShapeRepo,
	)
}

func (suite *ScheduleItemCreateUseCaseTestSuite) TestExecute_Success() {
	// Arrange
	input := dto.ScheduleItemCreateInputDTO{
		Name:            "Test Item",
		ShapeID:         1,
		ScheduleID:      1,
		ShapeDimensions: datatypes.JSON(`{"width": 10, "height": 20}`),
	}

	schedule := &models.Schedule{}
	schedule.ID = 1

	shape := &entities.Shape{
		ID:   1,
		Name: "Rectangle",
	}

	expectedOutput := &dto.ScheduleItemCreateOutputDTO{
		ID:              1,
		Name:            "Test Item",
		ShapeID:         1,
		ScheduleID:      1,
		ShapeDimensions: datatypes.JSON(`{"width": 10, "height": 20}`),
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
	}

	suite.mockScheduleRepo.On("GetByID", uint(1)).Return(schedule, nil)
	suite.mockShapeRepo.On("GetByID", uint(1)).Return(shape, nil)
	suite.mockScheduleItemRepo.On("Create", input).Return(expectedOutput, nil)

	// Act
	result, err := suite.useCase.Execute(input)

	// Assert
	assert.NoError(suite.T(), err)
	assert.NotNil(suite.T(), result)
	assert.Equal(suite.T(), expectedOutput.ID, result.ID)
	assert.Equal(suite.T(), expectedOutput.Name, result.Name)
	suite.mockScheduleRepo.AssertExpectations(suite.T())
	suite.mockShapeRepo.AssertExpectations(suite.T())
	suite.mockScheduleItemRepo.AssertExpectations(suite.T())
}

func (suite *ScheduleItemCreateUseCaseTestSuite) TestExecute_ScheduleNotFound() {
	// Arrange
	input := dto.ScheduleItemCreateInputDTO{
		Name:            "Test Item",
		ShapeID:         1,
		ScheduleID:      999,
		ShapeDimensions: datatypes.JSON(`{"width": 10, "height": 20}`),
	}

	suite.mockScheduleRepo.On("GetByID", uint(999)).Return(nil, errors.New("schedule not found"))

	// Act
	result, err := suite.useCase.Execute(input)

	// Assert
	assert.Error(suite.T(), err)
	assert.Nil(suite.T(), result)
	assert.Contains(suite.T(), err.Error(), "schedule not found")
	suite.mockScheduleRepo.AssertExpectations(suite.T())
}

func (suite *ScheduleItemCreateUseCaseTestSuite) TestExecute_ShapeNotFound() {
	// Arrange
	input := dto.ScheduleItemCreateInputDTO{
		Name:            "Test Item",
		ShapeID:         999,
		ScheduleID:      1,
		ShapeDimensions: datatypes.JSON(`{"width": 10, "height": 20}`),
	}

	schedule := &models.Schedule{}
	schedule.ID = 1

	suite.mockScheduleRepo.On("GetByID", uint(1)).Return(schedule, nil)
	suite.mockShapeRepo.On("GetByID", uint(999)).Return(nil, errors.New("shape not found"))

	// Act
	result, err := suite.useCase.Execute(input)

	// Assert
	assert.Error(suite.T(), err)
	assert.Nil(suite.T(), result)
	assert.Contains(suite.T(), err.Error(), "shape not found")
	suite.mockScheduleRepo.AssertExpectations(suite.T())
	suite.mockShapeRepo.AssertExpectations(suite.T())
}

func (suite *ScheduleItemCreateUseCaseTestSuite) TestExecute_RepositoryError() {
	// Arrange
	input := dto.ScheduleItemCreateInputDTO{
		Name:            "Test Item",
		ShapeID:         1,
		ScheduleID:      1,
		ShapeDimensions: datatypes.JSON(`{"width": 10, "height": 20}`),
	}

	schedule := &models.Schedule{}
	schedule.ID = 1

	shape := &entities.Shape{
		ID:   1,
		Name: "Rectangle",
	}

	suite.mockScheduleRepo.On("GetByID", uint(1)).Return(schedule, nil)
	suite.mockShapeRepo.On("GetByID", uint(1)).Return(shape, nil)
	suite.mockScheduleItemRepo.On("Create", input).Return(nil, errors.New("database error"))

	// Act
	result, err := suite.useCase.Execute(input)

	// Assert
	assert.Error(suite.T(), err)
	assert.Nil(suite.T(), result)
	assert.Contains(suite.T(), err.Error(), "database error")
	suite.mockScheduleRepo.AssertExpectations(suite.T())
	suite.mockShapeRepo.AssertExpectations(suite.T())
	suite.mockScheduleItemRepo.AssertExpectations(suite.T())
}

func TestScheduleItemCreateUseCaseTestSuite(t *testing.T) {
	suite.Run(t, new(ScheduleItemCreateUseCaseTestSuite))
}
