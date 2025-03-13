package usecase

import (
	"github.com/snpiyasooriya/construction_design_api/dto"
)

type ScheduleItemCromeCreateUseCase interface {
	Execute(scheduleCreateInputDTO dto.ScheduleItemCromeCreateInputDTO) (*dto.ScheduleItemCromeCreateOutputDTO, error)
}
