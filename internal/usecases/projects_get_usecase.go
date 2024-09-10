package usecases

import "github.com/snpiyasooriya/construction_design_api/internal/dto"

type ProjectsGetUseCase interface {
	Execute() (*dto.ProjectsGetDTO, error)
}
