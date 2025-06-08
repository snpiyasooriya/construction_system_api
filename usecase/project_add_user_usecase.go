package usecase

import (
	"github.com/snpiyasooriya/construction_design_api/dto"
	"github.com/snpiyasooriya/construction_design_api/interfaces"
)

type ProjectAddUserUseCase interface {
	Execute(input dto.ProjectAddUserDTO) (*dto.ProjectAddUserOutputDTO, error)
}

type ProjectAddUserUseCaseImpl struct {
	projectRepo interfaces.ProjectRepository
}

func NewProjectAddUserUseCase(projectRepo interfaces.ProjectRepository) ProjectAddUserUseCase {
	return &ProjectAddUserUseCaseImpl{
		projectRepo: projectRepo,
	}
}

func (p *ProjectAddUserUseCaseImpl) Execute(input dto.ProjectAddUserDTO) (*dto.ProjectAddUserOutputDTO, error) {
	return p.projectRepo.AddUser(input)
}
