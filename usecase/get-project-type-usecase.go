package usecase

import (
	"github.com/snpiyasooriya/construction_design_api/dto"
	"github.com/snpiyasooriya/construction_design_api/interfaces/repository"
)

type GetProjectTypeUseCase struct {
	projectTypeRepo repository.ProjectTypeRepositoryInterface
}

func NewGetProjectTypeUseCase(projectTypeRepo repository.ProjectTypeRepositoryInterface) *GetProjectTypeUseCase {
	return &GetProjectTypeUseCase{
		projectTypeRepo: projectTypeRepo,
	}
}

func (g *GetProjectTypeUseCase) Execute(id uint) (dto.ProjectTypeGetDto, error) {
	projectTypeDto, err := g.projectTypeRepo.FindById(id)
	if err != nil {
		return dto.ProjectTypeGetDto{}, err
	}
	return projectTypeDto, nil
}
