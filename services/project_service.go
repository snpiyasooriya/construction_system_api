package services

import (
	dto2 "github.com/snpiyasooriya/construction_design_api/dto"
	usecases2 "github.com/snpiyasooriya/construction_design_api/usecase"
)

type ProjectService struct {
	projectCreateUseCase  usecases2.ProjectCreateUseCase
	projectUpdateUseCase  usecases2.ProjectUpdateUseCase
	projectDeleteUseCase  usecases2.ProjectDeleteUseCase
	scheduleCreateUseCase usecases2.ScheduleCreateUseCase
	projectAddUserUseCase usecases2.ProjectAddUserUseCase
}

func NewProjectCreateService(
	projectCreateUseCase usecases2.ProjectCreateUseCase,
	projectUpdateUseCase usecases2.ProjectUpdateUseCase,
	projectDeleteUseCase usecases2.ProjectDeleteUseCase,
	scheduleCreateUseCase usecases2.ScheduleCreateUseCase,
	projectAddUserUseCase usecases2.ProjectAddUserUseCase,
) *ProjectService {
	return &ProjectService{
		projectCreateUseCase:  projectCreateUseCase,
		projectUpdateUseCase:  projectUpdateUseCase,
		projectDeleteUseCase:  projectDeleteUseCase,
		scheduleCreateUseCase: scheduleCreateUseCase,
		projectAddUserUseCase: projectAddUserUseCase,
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

func (ps *ProjectService) UpdateProject(project dto2.ProjectUpdateDTO) error {
	return ps.projectUpdateUseCase.Execute(project)
}

func (ps *ProjectService) DeleteProject(id uint) error {
	return ps.projectDeleteUseCase.Execute(id)
}

func (ps *ProjectService) AddUserToProject(input dto2.ProjectAddUserDTO) (*dto2.ProjectAddUserOutputDTO, error) {
	return ps.projectAddUserUseCase.Execute(input)
}
