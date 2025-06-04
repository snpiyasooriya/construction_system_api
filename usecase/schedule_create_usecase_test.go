package usecase

import (
	"errors"
	"testing"
	"time"

	"github.com/snpiyasooriya/construction_design_api/dto"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

// MockScheduleRepository is a mock implementation of ScheduleRepository
type MockScheduleRepository struct {
	mock.Mock
}

func (m *MockScheduleRepository) Create(schedule dto.ScheduleCreateInputDTO) (*dto.ScheduleCreateOutputDTO, error) {
	args := m.Called(schedule)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*dto.ScheduleCreateOutputDTO), args.Error(1)
}

func (m *MockScheduleRepository) UpdateByID(schedule interface{}) (interface{}, error) {
	args := m.Called(schedule)
	return args.Get(0), args.Error(1)
}

func (m *MockScheduleRepository) GetByID(id uint) (interface{}, error) {
	args := m.Called(id)
	return args.Get(0), args.Error(1)
}

func (m *MockScheduleRepository) Get() (interface{}, error) {
	args := m.Called()
	return args.Get(0), args.Error(1)
}

func (m *MockScheduleRepository) DeleteByID(id uint) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockScheduleRepository) GetCountByProjectID(projectID uint) (int, error) {
	args := m.Called(projectID)
	return args.Int(0), args.Error(1)
}

func (m *MockScheduleRepository) GetByProjectID(projectID uint) (interface{}, error) {
	args := m.Called(projectID)
	return args.Get(0), args.Error(1)
}

type ScheduleCreateUseCaseTestSuite struct {
	suite.Suite
	mockRepo *MockScheduleRepository
	useCase  *ScheduleCreateUseCaseImpl
}

func (suite *ScheduleCreateUseCaseTestSuite) SetupTest() {
	suite.mockRepo = new(MockScheduleRepository)
	suite.useCase = NewScheduleCreateUseCaseImpl(suite.mockRepo)
}

func (suite *ScheduleCreateUseCaseTestSuite) TestExecute_Success() {
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

	suite.mockRepo.On("Create", input).Return(expectedOutput, nil)

	// Act
	result, err := suite.useCase.Execute(input)

	// Assert
	assert.NoError(suite.T(), err)
	assert.NotNil(suite.T(), result)
	assert.Equal(suite.T(), expectedOutput.ID, result.ID)
	assert.Equal(suite.T(), expectedOutput.Name, result.Name)
	assert.Equal(suite.T(), expectedOutput.Description, result.Description)
	assert.Equal(suite.T(), expectedOutput.ProjectID, result.ProjectID)

	suite.mockRepo.AssertExpectations(suite.T())
}

func (suite *ScheduleCreateUseCaseTestSuite) TestExecute_RepositoryError() {
	// Arrange
	input := dto.ScheduleCreateInputDTO{
		Name:        "Test Schedule",
		Description: "Test Description",
		ProjectID:   1,
	}

	expectedError := errors.New("database connection failed")
	suite.mockRepo.On("Create", input).Return(nil, expectedError)

	// Act
	result, err := suite.useCase.Execute(input)

	// Assert
	assert.Error(suite.T(), err)
	assert.Nil(suite.T(), result)
	assert.Equal(suite.T(), expectedError, err)

	suite.mockRepo.AssertExpectations(suite.T())
}

func (suite *ScheduleCreateUseCaseTestSuite) TestExecute_EmptyName() {
	// Arrange
	input := dto.ScheduleCreateInputDTO{
		Name:        "",
		Description: "Test Description",
		ProjectID:   1,
	}

	expectedOutput := &dto.ScheduleCreateOutputDTO{
		ID:          1,
		Name:        "",
		Description: "Test Description",
		ProjectID:   1,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	suite.mockRepo.On("Create", input).Return(expectedOutput, nil)

	// Act
	result, err := suite.useCase.Execute(input)

	// Assert
	assert.NoError(suite.T(), err)
	assert.NotNil(suite.T(), result)
	assert.Equal(suite.T(), "", result.Name)

	suite.mockRepo.AssertExpectations(suite.T())
}

func (suite *ScheduleCreateUseCaseTestSuite) TestExecute_ZeroProjectID() {
	// Arrange
	input := dto.ScheduleCreateInputDTO{
		Name:        "Test Schedule",
		Description: "Test Description",
		ProjectID:   0,
	}

	expectedOutput := &dto.ScheduleCreateOutputDTO{
		ID:          1,
		Name:        "Test Schedule",
		Description: "Test Description",
		ProjectID:   0,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	suite.mockRepo.On("Create", input).Return(expectedOutput, nil)

	// Act
	result, err := suite.useCase.Execute(input)

	// Assert
	assert.NoError(suite.T(), err)
	assert.NotNil(suite.T(), result)
	assert.Equal(suite.T(), uint(0), result.ProjectID)

	suite.mockRepo.AssertExpectations(suite.T())
}

func (suite *ScheduleCreateUseCaseTestSuite) TestExecute_LargeProjectID() {
	// Arrange
	input := dto.ScheduleCreateInputDTO{
		Name:        "Test Schedule",
		Description: "Test Description",
		ProjectID:   999999,
	}

	expectedOutput := &dto.ScheduleCreateOutputDTO{
		ID:          1,
		Name:        "Test Schedule",
		Description: "Test Description",
		ProjectID:   999999,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	suite.mockRepo.On("Create", input).Return(expectedOutput, nil)

	// Act
	result, err := suite.useCase.Execute(input)

	// Assert
	assert.NoError(suite.T(), err)
	assert.NotNil(suite.T(), result)
	assert.Equal(suite.T(), uint(999999), result.ProjectID)

	suite.mockRepo.AssertExpectations(suite.T())
}

func TestScheduleCreateUseCaseTestSuite(t *testing.T) {
	suite.Run(t, new(ScheduleCreateUseCaseTestSuite))
}
