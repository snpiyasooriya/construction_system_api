package usecase

import (
	"github.com/snpiyasooriya/construction_design_api/dto"
)

type ScheduleGetByProjectUseCase interface {
	Execute(projectID uint) ([]dto.ScheduleGetByProjectOutputDTO, error)
}
