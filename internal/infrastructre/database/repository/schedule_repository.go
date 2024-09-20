package repository

import (
	"github.com/snpiyasooriya/construction_design_api/internal/dto"
	"github.com/snpiyasooriya/construction_design_api/internal/infrastructre/database/models"
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
