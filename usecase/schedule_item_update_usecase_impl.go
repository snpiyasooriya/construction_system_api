package usecase

import (
	"errors"

	"github.com/snpiyasooriya/construction_design_api/dto"
	"github.com/snpiyasooriya/construction_design_api/interfaces/repository"
)

type ScheduleItemUpdateUseCaseImpl struct {
	scheduleItemRepo repository.ScheduleItemRepository
	shapeRepo        repository.ShapeRepositoryInterFace
}

func NewScheduleItemUpdateUseCaseImpl(
	scheduleItemRepo repository.ScheduleItemRepository,
	shapeRepo repository.ShapeRepositoryInterFace,
) *ScheduleItemUpdateUseCaseImpl {
	return &ScheduleItemUpdateUseCaseImpl{
		scheduleItemRepo: scheduleItemRepo,
		shapeRepo:        shapeRepo,
	}
}

func (s *ScheduleItemUpdateUseCaseImpl) Execute(scheduleItemUpdateInputDTO dto.ScheduleItemUpdateInputDTO) (*dto.ScheduleItemUpdateOutputDTO, error) {
	// Check if schedule item exists
	existingItem, err := s.scheduleItemRepo.GetByID(scheduleItemUpdateInputDTO.ID)
	if err != nil {
		return nil, errors.New("schedule item not found")
	}
	if existingItem == nil {
		return nil, errors.New("schedule item not found")
	}

	// Validate that the shape exists
	shape, err := s.shapeRepo.GetByID(scheduleItemUpdateInputDTO.ShapeID)
	if err != nil {
		return nil, errors.New("shape not found")
	}
	if shape == nil {
		return nil, errors.New("shape not found")
	}

	// Update the schedule item
	scheduleItemOutput, err := s.scheduleItemRepo.UpdateByID(scheduleItemUpdateInputDTO)
	if err != nil {
		return nil, err
	}

	return scheduleItemOutput, nil
}
