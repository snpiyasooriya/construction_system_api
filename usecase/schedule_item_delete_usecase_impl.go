package usecase

import (
	"errors"
	"time"

	"github.com/snpiyasooriya/construction_design_api/dto"
	"github.com/snpiyasooriya/construction_design_api/interfaces/repository"
)

type ScheduleItemDeleteUseCaseImpl struct {
	scheduleItemRepo repository.ScheduleItemRepository
}

func NewScheduleItemDeleteUseCaseImpl(scheduleItemRepo repository.ScheduleItemRepository) *ScheduleItemDeleteUseCaseImpl {
	return &ScheduleItemDeleteUseCaseImpl{
		scheduleItemRepo: scheduleItemRepo,
	}
}

func (s *ScheduleItemDeleteUseCaseImpl) Execute(id uint) (*dto.ScheduleItemDeleteOutputDTO, error) {
	// Check if schedule item exists
	existingItem, err := s.scheduleItemRepo.GetByID(id)
	if err != nil {
		return nil, errors.New("schedule item not found")
	}
	if existingItem == nil {
		return nil, errors.New("schedule item not found")
	}

	// Delete the schedule item
	err = s.scheduleItemRepo.DeleteByID(id)
	if err != nil {
		return nil, err
	}

	// Return success response
	output := &dto.ScheduleItemDeleteOutputDTO{
		ID:        id,
		DeletedAt: time.Now(),
		Message:   "Schedule item deleted successfully",
	}

	return output, nil
}
