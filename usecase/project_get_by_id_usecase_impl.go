package usecase

import (
	"github.com/snpiyasooriya/construction_design_api/dto"
	"github.com/snpiyasooriya/construction_design_api/interfaces/repository"
)

type projectGetByIDUseCaseImpl struct {
	projectRepository repository.ProjectRepository
}

func NewProjectGetByIDUseCase(projectRepository repository.ProjectRepository) ProjectGetByIDUseCase {
	return &projectGetByIDUseCaseImpl{
		projectRepository: projectRepository,
	}
}

func (uc *projectGetByIDUseCaseImpl) Execute(id uint) (*dto.ProjectGetDTO, error) {
	project, err := uc.projectRepository.GetByID(id)
	if err != nil {
		return nil, err
	}

	// Get user IDs
	userIDs := make([]uint, len(project.Users))
	for i, user := range project.Users {
		userIDs[i] = user.ID
	}

	return &dto.ProjectGetDTO{
		ID:            project.ID,
		Name:          project.Name,
		ProjectID:     project.ProjectID,
		ProjectType:   project.ProjectType.Type,
		LeaderID:      project.LeaderID,
		StartDate:     project.StartDate,
		EndDate:       project.EndDate,
		Note:          project.Note,
		Status:        project.Status,
		CreatedAt:     project.CreatedAt,
		UpdatedAt:     project.UpdatedAt,
		ScheduleCount: len(project.Schedules),
		UserIDs:       userIDs,
	}, nil
}
