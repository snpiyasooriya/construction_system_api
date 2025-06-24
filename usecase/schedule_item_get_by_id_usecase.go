package usecase

import (
	"github.com/snpiyasooriya/construction_design_api/dto"
)

type ScheduleItemGetByIDUseCase interface {
	Execute(id uint) (*dto.ScheduleItemGetByIDOutputDTO, error)
}
