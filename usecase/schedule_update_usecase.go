package usecase

import (
    "github.com/snpiyasooriya/construction_design_api/dto"
)

type ScheduleUpdateUseCase interface {
    Execute(input dto.ScheduleUpdateDTO) (*dto.ScheduleUpdateOutputDTO, error)
}