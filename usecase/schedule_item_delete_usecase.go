package usecase

import (
	"github.com/snpiyasooriya/construction_design_api/dto"
)

type ScheduleItemDeleteUseCase interface {
	Execute(id uint) (*dto.ScheduleItemDeleteOutputDTO, error)
}
