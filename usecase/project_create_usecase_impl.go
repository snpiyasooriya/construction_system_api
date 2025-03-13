package usecase

import (
	"github.com/snpiyasooriya/construction_design_api/dto"
	"github.com/snpiyasooriya/construction_design_api/interfaces/repository"
)

type ProjectCreateUseCaseImpl struct {
	projectRepo repository.ProjectRepository
}

func NewProjectCreateUseCaseImpl(projectRepo repository.ProjectRepository) *ProjectCreateUseCaseImpl {
	return &ProjectCreateUseCaseImpl{projectRepo: projectRepo}
}

func (pci *ProjectCreateUseCaseImpl) Execute(input dto.ProjectCreateInputDTO) (*dto.ProjectCreateOutputDTO, error) {
	create, err := pci.projectRepo.Create(input)
	if err != nil {
		return nil, err
	}
	return create, err
}
