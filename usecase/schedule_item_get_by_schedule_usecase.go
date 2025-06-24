package usecase

import (
	"github.com/snpiyasooriya/construction_design_api/dto"
)

type ScheduleItemGetByScheduleUseCase interface {
	Execute(scheduleID uint) ([]dto.ScheduleItemGetByScheduleOutputDTO, error)
}
