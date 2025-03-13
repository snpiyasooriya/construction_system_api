package repository

import (
	"github.com/snpiyasooriya/construction_design_api/dto"
	"github.com/snpiyasooriya/construction_design_api/models"
	"gorm.io/gorm"
)

type GORMScheduleRepository struct {
	db *gorm.DB
}

func NewGORMScheduleRepository(db *gorm.DB) *GORMScheduleRepository {
	return &GORMScheduleRepository{
		db: db,
	}
}

func (g *GORMScheduleRepository) Create(schedule dto.ScheduleCreateInputDTO) (*dto.ScheduleCreateOutputDTO, error) {
	createdSchedule := models.Schedule{
		ProjectID: schedule.ProjectID,
	}
	if err := g.db.Create(&createdSchedule).Error; err != nil {
		return nil, err
	}
	scheduleCreateOutputDTO := dto.ScheduleCreateOutputDTO{
		ID:          createdSchedule.ID,
		Name:        createdSchedule.Name,
		Description: createdSchedule.Description,
		ProjectID:   createdSchedule.ProjectID,
		CreatedAt:   createdSchedule.CreatedAt,
		UpdatedAt:   createdSchedule.UpdatedAt,
	}

	return &scheduleCreateOutputDTO, nil
}

func (g *GORMScheduleRepository) UpdateByID(schedule models.Schedule) (*models.Schedule, error) {
	//TODO implement me
	panic("implement me")
}

func (g *GORMScheduleRepository) GetByID(id uint) (*models.Schedule, error) {
	//TODO implement me
	panic("implement me")
}

func (g *GORMScheduleRepository) Get() ([]models.Schedule, error) {
	//TODO implement me
	panic("implement me")
}

func (g *GORMScheduleRepository) DeleteByID(id uint) error {
	//TODO implement me
	panic("implement me")
}

func (g *GORMScheduleRepository) GetCountByProjectID(projectID uint) (int, error) {
	var count int64
	if err := g.db.Model(&models.Schedule{}).Where("project_id = ?", projectID).Count(&count).Error; err != nil {
		return 0, err
	}
	return int(count), nil
}

// GetByProjectID retrieves all schedules associated with the specified project ID.
// It takes a project ID as input and returns a slice of Schedule models along with any error that occurred.
func (g *GORMScheduleRepository) GetByProjectID(projectID uint) ([]models.Schedule, error) {
	var schedules []models.Schedule
	if err := g.db.Where("project_id = ?", projectID).Find(&schedules).Error; err != nil {
		return nil, err
	}
	return schedules, nil
}
