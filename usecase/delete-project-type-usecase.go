package usecase

import "github.com/snpiyasooriya/construction_design_api/interfaces/repository"

type DeleteProjectTypeUseCase struct {
	projectTypeRepo repository.ProjectTypeRepositoryInterface
}

func NewDeleteProjectTypeUseCase(projectTypeRepo repository.ProjectTypeRepositoryInterface) *DeleteProjectTypeUseCase {
	return &DeleteProjectTypeUseCase{
		projectTypeRepo: projectTypeRepo,
	}
}

func (d *DeleteProjectTypeUseCase) Execute(id uint) error {
	return d.projectTypeRepo.DeleteById(id)
}
