package usecases

import (
	"github.com/snpiyasooriya/construction_design_api/internal/domain/interfaces"
	"github.com/snpiyasooriya/construction_design_api/internal/dto"
)

type ProjectCreateUseCaseImpl struct {
	projectRepo interfaces.ProjectRepository
}

func NewProjectCreateUseCaseImpl(projectRepo interfaces.ProjectRepository) *ProjectCreateUseCaseImpl {
	return &ProjectCreateUseCaseImpl{projectRepo: projectRepo}
}

func (pci *ProjectCreateUseCaseImpl) Execute(input dto.ProjectCreateInputDTO) (*dto.ProjectCreateOutputDTO, error) {
	create, err := pci.projectRepo.Create(input)
	if err != nil {
		return nil, err
	}
	return create, err
}
