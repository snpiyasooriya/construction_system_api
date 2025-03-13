package services

import (
	dto2 "github.com/snpiyasooriya/construction_design_api/dto"
	usecases2 "github.com/snpiyasooriya/construction_design_api/usecase"
)

type ProjectService struct {
	projectCreateUseCase  usecases2.ProjectCreateUseCase
	scheduleCreateUseCase usecases2.ScheduleCreateUseCase
	//scheduleItemCromeCreateUseCase usecase.ScheduleItemCromeCreateUseCase
}

func NewProjectCreateService(projectCreateUseCase usecases2.ProjectCreateUseCase, sheduleCreateUseCase usecases2.ScheduleCreateUseCase) *ProjectService {
	return &ProjectService{
		projectCreateUseCase:  projectCreateUseCase,
		scheduleCreateUseCase: sheduleCreateUseCase,
	}
}

func (ps *ProjectService) CreateProject(project dto2.ProjectCreateInputDTO) error {
	createdProject, err := ps.projectCreateUseCase.Execute(project)
	if err != nil {
		return err
	}
	scheduleCreateDTO := dto2.ScheduleCreateInputDTO{
		ProjectID: createdProject.ID,
	}
	_, err = ps.scheduleCreateUseCase.Execute(scheduleCreateDTO)
	if err != nil {
		return err
	}
	return nil
}
