package usecases

import (
	"github.com/snpiyasooriya/construction_design_api/internal/domain/interfaces"
	"github.com/snpiyasooriya/construction_design_api/internal/dto"
)

type ScheduleGetByProjectUseCaseImpl struct {
	scheduleRepository interfaces.ScheduleRepository
}

func NewScheduleGetByProjectUseCaseImpl(scheduleRepository interfaces.ScheduleRepository) *ScheduleGetByProjectUseCaseImpl {
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
			ID:          schedule.ID,
			Name:        schedule.Name,
			Description: schedule.Description,
			ProjectID:   schedule.ProjectID,
			CreatedAt:   schedule.CreatedAt,
			UpdatedAt:   schedule.UpdatedAt,
		}
	}
	return scheduleGetByProjectOutputDTOs, nil
}
