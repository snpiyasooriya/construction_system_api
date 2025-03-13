package usecase

import (
	"github.com/snpiyasooriya/construction_design_api/dto"
)

type ProjectCreateUseCase interface {
	Execute(input dto.ProjectCreateInputDTO) (*dto.ProjectCreateOutputDTO, error)
}
