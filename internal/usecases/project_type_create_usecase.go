package usecases

import "github.com/snpiyasooriya/construction_design_api/internal/dto"

type ProjectTypeCreateUseCase interface {
	Execute(dto dto.ProjectTypeCreateInputDTO) (*dto.ProjectTypeCreateOutputDTO, error)
}
