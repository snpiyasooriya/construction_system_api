package usecase

import (
	"errors"
	"fmt"

	"github.com/snpiyasooriya/construction_design_api/constants"
	"github.com/snpiyasooriya/construction_design_api/dto"
	"github.com/snpiyasooriya/construction_design_api/interfaces/repository"
)

type ScheduleUpdateUseCaseImpl struct {
	scheduleRepository repository.ScheduleRepository
}

func NewScheduleUpdateUseCaseImpl(scheduleRepository repository.ScheduleRepository) *ScheduleUpdateUseCaseImpl {
	return &ScheduleUpdateUseCaseImpl{
		scheduleRepository: scheduleRepository,
	}
}

func (s *ScheduleUpdateUseCaseImpl) Execute(input dto.ScheduleUpdateDTO) (*dto.ScheduleUpdateOutputDTO, error) {
	// Check if schedule exists
	existingSchedule, err := s.scheduleRepository.GetByID(input.ID)
	if err != nil {
		return nil, errors.New("schedule not found")
	}

	// Update schedule fields
	existingSchedule.Description = input.Description
	existingSchedule.RequiredDate = input.RequiredDate
	existingSchedule.Status = constants.ScheduleStatus(input.Status)
	existingSchedule.Note = input.Note

	// Update reviewer if provided
	if input.ReviewerID > 0 {
		existingSchedule.ReviewerID = &input.ReviewerID
	}

	// Save updated schedule
	updatedSchedule, err := s.scheduleRepository.UpdateByID(*existingSchedule)
	if err != nil {
		return nil, err
	}

	// Create output DTO
	return &dto.ScheduleUpdateOutputDTO{
		ID:           updatedSchedule.ID,
		ScheduleID:   updatedSchedule.ScheduleID,
		ProjectID:    updatedSchedule.ProjectID,
		Description:  updatedSchedule.Description,
		RequiredDate: updatedSchedule.RequiredDate,
		SchedularID:  updatedSchedule.SchedularID,
		Status:       string(updatedSchedule.Status),
		Note:         updatedSchedule.Note,
		Schedular:    fmt.Sprintf("%s %s", updatedSchedule.Schedular.FirstName, updatedSchedule.Schedular.LastName),
		ReviewerID: func() uint {
			if updatedSchedule.ReviewerID == nil {
				return 0
			} else {
				return *updatedSchedule.ReviewerID
			}
		}(),
		Reviewer: func() string {
			if updatedSchedule.ReviewerID == nil {
				return ""
			} else {
				return fmt.Sprintf("%s %s", updatedSchedule.Reviewer.FirstName, updatedSchedule.Reviewer.LastName)
			}
		}(),
		UpdatedAt: updatedSchedule.UpdatedAt,
	}, nil
}
