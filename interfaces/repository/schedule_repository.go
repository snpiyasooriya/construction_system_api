package repository

import (
	"github.com/snpiyasooriya/construction_design_api/dto"
	"github.com/snpiyasooriya/construction_design_api/models"
)

type ScheduleRepository interface {
	Create(schedule dto.ScheduleCreateInputDTO) (*dto.ScheduleCreateOutputDTO, error)
	UpdateByID(schedule models.Schedule) (*models.Schedule, error)
	GetByID(id uint) (*models.Schedule, error)
	Get() ([]models.Schedule, error)
	DeleteByID(id uint) error
	GetCountByProjectID(projectID uint) (int, error)
	GetByProjectID(projectID uint) ([]models.Schedule, error)
}
