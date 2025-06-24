package usecase

import (
	"errors"

	"github.com/snpiyasooriya/construction_design_api/dto"
	"github.com/snpiyasooriya/construction_design_api/interfaces/repository"
)

type ScheduleItemGetByScheduleUseCaseImpl struct {
	scheduleItemRepo repository.ScheduleItemRepository
	scheduleRepo     repository.ScheduleRepository
}

func NewScheduleItemGetByScheduleUseCaseImpl(
	scheduleItemRepo repository.ScheduleItemRepository,
	scheduleRepo repository.ScheduleRepository,
) *ScheduleItemGetByScheduleUseCaseImpl {
	return &ScheduleItemGetByScheduleUseCaseImpl{
		scheduleItemRepo: scheduleItemRepo,
		scheduleRepo:     scheduleRepo,
	}
}

func (s *ScheduleItemGetByScheduleUseCaseImpl) Execute(scheduleID uint) ([]dto.ScheduleItemGetByScheduleOutputDTO, error) {
	// Validate that the schedule exists
	schedule, err := s.scheduleRepo.GetByID(scheduleID)
	if err != nil {
		return nil, errors.New("schedule not found")
	}
	if schedule == nil {
		return nil, errors.New("schedule not found")
	}

	scheduleItems, err := s.scheduleItemRepo.GetByScheduleID(scheduleID)
	if err != nil {
		return nil, err
	}

	// Map models to DTOs
	var output []dto.ScheduleItemGetByScheduleOutputDTO
	for _, item := range scheduleItems {
		outputItem := dto.ScheduleItemGetByScheduleOutputDTO{
			ID:              item.ID,
			Name:            item.Name,
			ShapeID:         item.ShapeID,
			ShapeName:       item.Shape.Name,
			ShapeDimensions: item.ShapeDimensions,
			CreatedAt:       item.CreatedAt,
			UpdatedAt:       item.UpdatedAt,
		}
		output = append(output, outputItem)
	}

	return output, nil
}
