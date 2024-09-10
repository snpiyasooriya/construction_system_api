package usecases

import (
	"github.com/snpiyasooriya/construction_design_api/internal/domain/interfaces"
	"github.com/snpiyasooriya/construction_design_api/internal/dto"
)

type ProjectsGetUseCaseImpl struct {
	projectRepository interfaces.ProjectRepository
}

func NewProjectsGetUseCaseImpl(repository interfaces.ProjectRepository) *ProjectsGetUseCaseImpl {
	return &ProjectsGetUseCaseImpl{
		projectRepository: repository,
	}
}

func (p *ProjectsGetUseCaseImpl) Execute() (*dto.ProjectsGetDTO, error) {
	projects, err := p.projectRepository.Get()
	if err != nil {
		return nil, err
	}
	return projects, nil
}
