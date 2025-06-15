package usecase

import (
	"github.com/snpiyasooriya/construction_design_api/dto"
)

type ScheduleGetByIDUseCase interface {
	Execute(id uint) (*dto.ScheduleGetByIDOutputDTO, error)
}
