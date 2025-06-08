package usecase

import (
	"github.com/snpiyasooriya/construction_design_api/dto"
	"github.com/snpiyasooriya/construction_design_api/interfaces/repository"
)

type GetAllProjectTypesUseCase struct {
	projectTypeRepository repository.ProjectTypeRepositoryInterface
}

func NewGetAllProjectTypesUseCase(projectTypeRepository repository.ProjectTypeRepositoryInterface) *GetAllProjectTypesUseCase {
	return &GetAllProjectTypesUseCase{
		projectTypeRepository: projectTypeRepository,
	}
}

func (p *GetAllProjectTypesUseCase) Execute() ([]dto.ProjectTypeGetDto, error) {
	projectTypes, err := p.projectTypeRepository.Find()
	if err != nil {
		return nil, err
	}
	return projectTypes, nil
}
