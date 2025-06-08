package usecase

import (
	"fmt"
	"github.com/snpiyasooriya/construction_design_api/dto"
	"github.com/snpiyasooriya/construction_design_api/interfaces/repository"
)

type ScheduleGetByProjectUseCaseImpl struct {
	scheduleRepository repository.ScheduleRepository
}

func NewScheduleGetByProjectUseCaseImpl(scheduleRepository repository.ScheduleRepository) *ScheduleGetByProjectUseCaseImpl {
	return &ScheduleGetByProjectUseCaseImpl{
		scheduleRepository: scheduleRepository,
	}
}
func (s *ScheduleGetByProjectUseCaseImpl) Execute(projectID uint) ([]dto.ScheduleGetByProjectOutputDTO, error) {
	schedules, err := s.scheduleRepository.GetByProjectID(projectID)
	if err != nil {
		return nil, err
	}
	scheduleGetByProjectOutputDTOs := make([]dto.ScheduleGetByProjectOutputDTO, len(schedules))
	for i, schedule := range schedules {
		scheduleGetByProjectOutputDTOs[i] = dto.ScheduleGetByProjectOutputDTO{
			ID:           schedule.ID,
			ScheduleID:   schedule.ScheduleID,
			ProjectID:    schedule.ProjectID,
			Description:  schedule.Description,
			RequiredDate: schedule.RequiredDate,
			SchedularID:  schedule.SchedularID,
			Status:       schedule.Status,
			Note:         schedule.Note,
			Schedular:    fmt.Sprintf("%s %s", schedule.Schedular.FirstName, schedule.Schedular.LastName),
			ReviewerID: func() uint {
				if schedule.ReviewerID == nil {
					return 0
				} else {
					return *schedule.ReviewerID
				}
			}(),
			Reviewer: func() string {
				if schedule.ReviewerID == nil {
					return ""
				} else {
					return fmt.Sprintf("%s %s", schedule.Reviewer.FirstName, schedule.Reviewer.LastName)
				}
			}(),
			CreatedAt: schedule.CreatedAt,
			UpdatedAt: schedule.UpdatedAt,
		}
	}
	return scheduleGetByProjectOutputDTOs, nil
}
