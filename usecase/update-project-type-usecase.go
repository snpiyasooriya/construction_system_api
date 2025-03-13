package usecase

import (
	"github.com/snpiyasooriya/construction_design_api/dto"
	"github.com/snpiyasooriya/construction_design_api/interfaces/repository"
)

type ProjectTypeUpdateUseCase struct {
	projectTypeRepository repository.ProjectTypeRepositoryInterface
}

func NewProjectTypeUpdateUseCase(projectTypeRepo repository.ProjectTypeRepositoryInterface) *ProjectTypeUpdateUseCase {
	return &ProjectTypeUpdateUseCase{
		projectTypeRepository: projectTypeRepo,
	}
}

func (p *ProjectTypeUpdateUseCase) Execute(input *dto.ProjectTypeUpdateDTO) error {
	err := p.projectTypeRepository.Update(input)
	if err != nil {
		return err
	}
	return nil
}
