package usecases

import (
	"github.com/snpiyasooriya/construction_design_api/internal/domain/interfaces"
	"github.com/snpiyasooriya/construction_design_api/internal/dto"
)

type ProjectsGetUseCaseImpl struct {
	projectRepository  interfaces.ProjectRepository
	scheduleRepository interfaces.ScheduleRepository // Add this line

}

func NewProjectsGetUseCaseImpl(repository interfaces.ProjectRepository, scheduleRepository interfaces.ScheduleRepository) *ProjectsGetUseCaseImpl {
	return &ProjectsGetUseCaseImpl{
		projectRepository:  repository,
		scheduleRepository: scheduleRepository,
	}
}

func (p *ProjectsGetUseCaseImpl) Execute() (*dto.ProjectsGetDTO, error) {
	projects, err := p.projectRepository.Get()
	if err != nil {
		return nil, err
	}
	// Fetch schedule counts for each project
	for i, project := range projects.Projects {
		count, err := p.scheduleRepository.GetCountByProjectID(project.ID)
		if err != nil {
			return nil, err
		}
		projects.Projects[i].ScheduleCount = count
	}
	return projects, nil
}
