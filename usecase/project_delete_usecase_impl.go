package usecase

import "github.com/snpiyasooriya/construction_design_api/interfaces/repository"

type projectDeleteUseCaseImpl struct {
	projectRepository repository.ProjectRepository
}

func NewProjectDeleteUseCase(projectRepository repository.ProjectRepository) ProjectDeleteUseCase {
	return &projectDeleteUseCaseImpl{
		projectRepository: projectRepository,
	}
}

func (uc *projectDeleteUseCaseImpl) Execute(id uint) error {
	return uc.projectRepository.DeleteByID(id)
}
