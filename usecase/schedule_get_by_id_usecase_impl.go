package usecase

import (
	"fmt"

	"github.com/snpiyasooriya/construction_design_api/dto"
	"github.com/snpiyasooriya/construction_design_api/interfaces/repository"
)

type ScheduleGetByIDUseCaseImpl struct {
	scheduleRepository repository.ScheduleRepository
}

func NewScheduleGetByIDUseCaseImpl(scheduleRepository repository.ScheduleRepository) *ScheduleGetByIDUseCaseImpl {
	return &ScheduleGetByIDUseCaseImpl{
		scheduleRepository: scheduleRepository,
	}
}

func (s *ScheduleGetByIDUseCaseImpl) Execute(id uint) (*dto.ScheduleGetByIDOutputDTO, error) {
	schedule, err := s.scheduleRepository.GetByID(id)
	if err != nil {
		return nil, err
	}

	// Create output DTO
	return &dto.ScheduleGetByIDOutputDTO{
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
	}, nil
}
