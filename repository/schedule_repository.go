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

func (g *GORMScheduleRepository) Create(schedule *dto.ScheduleCreateInputDTO) error {
	createdSchedule := models.Schedule{
		Description:  schedule.Description,
		ProjectID:    schedule.ProjectID,
		SchedularID:  schedule.SchedularID,
		ScheduleID:   schedule.ScheduleID,
		RequiredDate: schedule.RequiredDate,
		Note:         schedule.Note,
		Status:       schedule.Status,
	}
	if err := g.db.Create(&createdSchedule).Error; err != nil {
		return err
	}
	schedule.ID = createdSchedule.ID
	schedule.CreatedAt = createdSchedule.CreatedAt
	schedule.UpdatedAt = createdSchedule.UpdatedAt
	return nil
}

func (g *GORMScheduleRepository) UpdateByID(schedule models.Schedule) (*models.Schedule, error) {
	// Start a transaction
	tx := g.db.Begin()

	// Update the schedule
	if err := tx.Save(&schedule).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	// Commit the transaction
	if err := tx.Commit().Error; err != nil {
		return nil, err
	}

	// Fetch the updated schedule with all relations
	var updatedSchedule models.Schedule
	if err := g.db.Where("id = ?", schedule.ID).Preload("Schedular").Preload("Reviewer").Preload("Project").First(&updatedSchedule).Error; err != nil {
		return nil, err
	}

	return &updatedSchedule, nil
}

func (g *GORMScheduleRepository) GetByID(id uint) (*models.Schedule, error) {
	var schedule models.Schedule
	if err := g.db.Where("id = ?", id).Preload("Schedular").Preload("Reviewer").Preload("Project").First(&schedule).Error; err != nil {
		return nil, err
	}
	return &schedule, nil
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
	if err := g.db.Where("project_id = ?", projectID).Preload("Schedular").Preload("Reviewer").Preload("Project").Find(&schedules).Error; err != nil {
		return nil, err
	}
	return schedules, nil
}
