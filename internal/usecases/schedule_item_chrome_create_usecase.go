package usecases

import "github.com/snpiyasooriya/construction_design_api/internal/dto"

type ScheduleItemCromeCreateUseCase interface {
	Execute(scheduleCreateInputDTO dto.ScheduleItemCromeCreateInputDTO) (*dto.ScheduleItemCromeCreateOutputDTO, error)
}
