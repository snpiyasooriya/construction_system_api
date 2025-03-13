package usecase

import (
	"github.com/snpiyasooriya/construction_design_api/dto"
	"github.com/snpiyasooriya/construction_design_api/interfaces/repository"
)

type ProjectTypeCreateUseCase struct {
	projectTypeRepository repository.ProjectTypeRepositoryInterface
}

func NewProjectTypeCreateUseCase(userRepo repository.ProjectTypeRepositoryInterface) *ProjectTypeCreateUseCase {
	return &ProjectTypeCreateUseCase{
		projectTypeRepository: userRepo,
	}
}

func (p *ProjectTypeCreateUseCase) Execute(input *dto.ProjectTypeCreateDTO) error {
	err := p.projectTypeRepository.Create(input)
	if err != nil {
		return err
	}
	return nil
}
