package repository

import (
	"github.com/snpiyasooriya/construction_design_api/dto"
	"github.com/snpiyasooriya/construction_design_api/models"
)

type ScheduleItemRepository interface {
	Create(scheduleItem dto.ScheduleItemCreateInputDTO) (*dto.ScheduleItemCreateOutputDTO, error)
	UpdateByID(scheduleItem dto.ScheduleItemUpdateInputDTO) (*dto.ScheduleItemUpdateOutputDTO, error)
	GetByID(id uint) (*models.ScheduleItem, error)
	GetByScheduleID(scheduleID uint) ([]models.ScheduleItem, error)
	DeleteByID(id uint) error
}
