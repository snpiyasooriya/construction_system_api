package services

import (
	"github.com/snpiyasooriya/construction_design_api/internal/dto"
	"github.com/snpiyasooriya/construction_design_api/internal/usecases"
)

type ProjectService struct {
	projectCreateUseCase  usecases.ProjectCreateUseCase
	scheduleCreateUseCase usecases.ScheduleCreateUseCase
	//scheduleItemCromeCreateUseCase usecases.ScheduleItemCromeCreateUseCase
}

func NewProjectCreateService(projectCreateUseCase usecases.ProjectCreateUseCase, sheduleCreateUseCase usecases.ScheduleCreateUseCase) *ProjectService {
	return &ProjectService{
		projectCreateUseCase:  projectCreateUseCase,
		scheduleCreateUseCase: sheduleCreateUseCase,
	}
}

func (ps *ProjectService) CreateProject(project dto.ProjectCreateInputDTO) error {
	createdProject, err := ps.projectCreateUseCase.Execute(project)
	if err != nil {
		return err
	}
	scheduleCreateDTO := dto.ScheduleCreateInputDTO{
		ProjectID: createdProject.ID,
	}
	_, err = ps.scheduleCreateUseCase.Execute(scheduleCreateDTO)
	if err != nil {
		return err
	}
	return nil
}
