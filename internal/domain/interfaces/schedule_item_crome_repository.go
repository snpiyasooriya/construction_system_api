package interfaces

import (
	"github.com/snpiyasooriya/construction_design_api/internal/dto"
	"github.com/snpiyasooriya/construction_design_api/internal/infrastructre/database/models"
)

type ScheduleItemCrome interface {
	Create(schedule dto.ScheduleItemCromeCreateInputDTO) (*dto.ScheduleItemCromeCreateOutputDTO, error)
	UpdateByID(schedule models.ScheduleItemCrome) (*models.ScheduleItemCrome, error)
	GetByID(id uint) (*models.ScheduleItemCrome, error)
	Get() ([]models.ScheduleItemCrome, error)
	DeleteByID(id uint) error
}
