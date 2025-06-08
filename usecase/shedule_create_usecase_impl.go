package usecase

import (
	"errors"
	"fmt"
	"github.com/snpiyasooriya/construction_design_api/dto"
	"github.com/snpiyasooriya/construction_design_api/interfaces/repository"
)

type ScheduleCreateUseCaseImpl struct {
	scheduleRepo repository.ScheduleRepository
	projectRepo  repository.ProjectRepository
}

func NewScheduleCreateUseCaseImpl(scheduleRepo repository.ScheduleRepository, projectRepo repository.ProjectRepository) *ScheduleCreateUseCaseImpl {
	return &ScheduleCreateUseCaseImpl{
		scheduleRepo: scheduleRepo,
		projectRepo:  projectRepo,
	}

}

func (s *ScheduleCreateUseCaseImpl) Execute(scheduleCreateInputDTO *dto.ScheduleCreateInputDTO) error {
	project, err := s.projectRepo.GetByID(scheduleCreateInputDTO.ProjectID)
	if err != nil {
		return err
	}
	if project == nil {
		return errors.New("project not found")
	}
	scheduleCreateInputDTO.ProjectID = project.ID

	schedules, err := s.scheduleRepo.GetByProjectID(project.ID)
	if err != nil {
		return err
	}
	scheduleID := fmt.Sprintf("%d/%04d", project.ID, len(schedules)+1)
	fmt.Println(scheduleID)
	scheduleCreateInputDTO.ScheduleID = scheduleID

	err = s.scheduleRepo.Create(scheduleCreateInputDTO)
	if err != nil {
		return err
	}
	return nil
}
