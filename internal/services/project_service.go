package services

import (
	"github.com/snpiyasooriya/construction_design_api/internal/dto"
	"github.com/snpiyasooriya/construction_design_api/internal/usecases"
)

type ProjectService struct {
	projectCreateUseCase usecases.ProjectCreateUseCase
	//scheduleCreateUseCase          usecases.ScheduleCreateUseCase
	//scheduleItemCromeCreateUseCase usecases.ScheduleItemCromeCreateUseCase
}

func NewProjectCreateService(projectCreateUseCase usecases.ProjectCreateUseCase) *ProjectService {
	return &ProjectService{
		projectCreateUseCase: projectCreateUseCase,
	}
}

func (ps *ProjectService) CreateProject(project dto.ProjectCreateInputDTO) error {
	_, err := ps.projectCreateUseCase.Execute(project)
	if err != nil {
		return err
	}
	return nil
}
