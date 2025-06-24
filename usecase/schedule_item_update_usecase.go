package usecase

import (
	"github.com/snpiyasooriya/construction_design_api/dto"
)

type ScheduleItemUpdateUseCase interface {
	Execute(scheduleItemUpdateInputDTO dto.ScheduleItemUpdateInputDTO) (*dto.ScheduleItemUpdateOutputDTO, error)
}
