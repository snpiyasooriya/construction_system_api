package usecases

import (
	"github.com/snpiyasooriya/construction_design_api/internal/dto"
)

type ScheduleGetByProjectUseCase interface {
	Execute(projectID uint) ([]dto.ScheduleGetByProjectOutputDTO, error)
}
