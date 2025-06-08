package usecase

import (
	"github.com/snpiyasooriya/construction_design_api/dto"
)

type ScheduleCreateUseCase interface {
	Execute(scheduleCreateInputDTO *dto.ScheduleCreateInputDTO) error
}
