package usecases

import "github.com/snpiyasooriya/construction_design_api/internal/dto"

type ScheduleCreateUseCase interface {
	Execute(scheduleCreateInputDTO dto.ScheduleCreateInputDTO) (*dto.ScheduleCreateOutputDTO, error)
}
