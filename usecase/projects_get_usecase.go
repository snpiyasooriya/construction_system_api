package usecase

import (
	"github.com/snpiyasooriya/construction_design_api/dto"
)

type ProjectsGetUseCase interface {
	Execute() (*dto.ProjectsGetDTO, error)
}
