package usecase

import (
	"errors"

	"github.com/snpiyasooriya/construction_design_api/dto"
	"github.com/snpiyasooriya/construction_design_api/interfaces/repository"
)

type ScheduleItemGetByIDUseCaseImpl struct {
	scheduleItemRepo repository.ScheduleItemRepository
}

func NewScheduleItemGetByIDUseCaseImpl(scheduleItemRepo repository.ScheduleItemRepository) *ScheduleItemGetByIDUseCaseImpl {
	return &ScheduleItemGetByIDUseCaseImpl{
		scheduleItemRepo: scheduleItemRepo,
	}
}

func (s *ScheduleItemGetByIDUseCaseImpl) Execute(id uint) (*dto.ScheduleItemGetByIDOutputDTO, error) {
	scheduleItem, err := s.scheduleItemRepo.GetByID(id)
	if err != nil {
		return nil, err
	}
	if scheduleItem == nil {
		return nil, errors.New("schedule item not found")
	}

	// Map model to DTO
	output := &dto.ScheduleItemGetByIDOutputDTO{
		ID:              scheduleItem.ID,
		Name:            scheduleItem.Name,
		ShapeID:         scheduleItem.ShapeID,
		ShapeName:       scheduleItem.Shape.Name,
		ScheduleID:      scheduleItem.ScheduleID,
		ScheduleName:    scheduleItem.Schedule.Description,
		ShapeDimensions: scheduleItem.ShapeDimensions,
		CreatedAt:       scheduleItem.CreatedAt,
		UpdatedAt:       scheduleItem.UpdatedAt,
	}

	return output, nil
}
