package repository

import (
	"github.com/snpiyasooriya/construction_design_api/dto"
	"github.com/snpiyasooriya/construction_design_api/models"
	"gorm.io/gorm"
)

type GORMScheduleItemRepository struct {
	db *gorm.DB
}

func NewGORMScheduleItemRepository(db *gorm.DB) *GORMScheduleItemRepository {
	return &GORMScheduleItemRepository{
		db: db,
	}
}

func (g *GORMScheduleItemRepository) Create(scheduleItem dto.ScheduleItemCreateInputDTO) (*dto.ScheduleItemCreateOutputDTO, error) {
	createdScheduleItem := models.ScheduleItem{
		Name:            scheduleItem.Name,
		ShapeID:         scheduleItem.ShapeID,
		ScheduleID:      scheduleItem.ScheduleID,
		ShapeDimensions: scheduleItem.ShapeDimensions,
	}

	if err := g.db.Create(&createdScheduleItem).Error; err != nil {
		return nil, err
	}

	output := &dto.ScheduleItemCreateOutputDTO{
		ID:              createdScheduleItem.ID,
		Name:            createdScheduleItem.Name,
		ShapeID:         createdScheduleItem.ShapeID,
		ScheduleID:      createdScheduleItem.ScheduleID,
		ShapeDimensions: createdScheduleItem.ShapeDimensions,
		CreatedAt:       createdScheduleItem.CreatedAt,
		UpdatedAt:       createdScheduleItem.UpdatedAt,
	}

	return output, nil
}

func (g *GORMScheduleItemRepository) UpdateByID(scheduleItem dto.ScheduleItemUpdateInputDTO) (*dto.ScheduleItemUpdateOutputDTO, error) {
	var existingItem models.ScheduleItem
	if err := g.db.First(&existingItem, scheduleItem.ID).Error; err != nil {
		return nil, err
	}

	// Update fields
	existingItem.Name = scheduleItem.Name
	existingItem.ShapeID = scheduleItem.ShapeID
	existingItem.ShapeDimensions = scheduleItem.ShapeDimensions

	if err := g.db.Save(&existingItem).Error; err != nil {
		return nil, err
	}

	output := &dto.ScheduleItemUpdateOutputDTO{
		ID:              existingItem.ID,
		Name:            existingItem.Name,
		ShapeID:         existingItem.ShapeID,
		ScheduleID:      existingItem.ScheduleID,
		ShapeDimensions: existingItem.ShapeDimensions,
		UpdatedAt:       existingItem.UpdatedAt,
	}

	return output, nil
}

func (g *GORMScheduleItemRepository) GetByID(id uint) (*models.ScheduleItem, error) {
	var scheduleItem models.ScheduleItem
	if err := g.db.Where("id = ?", id).Preload("Shape").Preload("Schedule").First(&scheduleItem).Error; err != nil {
		return nil, err
	}
	return &scheduleItem, nil
}

func (g *GORMScheduleItemRepository) GetByScheduleID(scheduleID uint) ([]models.ScheduleItem, error) {
	var scheduleItems []models.ScheduleItem
	if err := g.db.Where("schedule_id = ?", scheduleID).Preload("Shape").Find(&scheduleItems).Error; err != nil {
		return nil, err
	}
	return scheduleItems, nil
}

func (g *GORMScheduleItemRepository) DeleteByID(id uint) error {
	if err := g.db.Delete(&models.ScheduleItem{}, id).Error; err != nil {
		return err
	}
	return nil
}
