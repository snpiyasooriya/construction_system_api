package usecase

import (
	"github.com/snpiyasooriya/construction_design_api/dto"
)

type ScheduleItemCreateUseCase interface {
	Execute(scheduleCreateInputDTO dto.ScheduleItemCreateInputDTO) (*dto.ScheduleItemCreateOutputDTO, error)
}
