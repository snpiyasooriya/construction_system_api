package usecase

import (
	"github.com/snpiyasooriya/construction_design_api/dto"
	"github.com/snpiyasooriya/construction_design_api/interfaces/repository"
)

type ProjectsGetUseCaseImpl struct {
	projectRepository  repository.ProjectRepository
	scheduleRepository repository.ScheduleRepository // Add this line

}

func NewProjectsGetUseCaseImpl(repository repository.ProjectRepository, scheduleRepository repository.ScheduleRepository) *ProjectsGetUseCaseImpl {
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
