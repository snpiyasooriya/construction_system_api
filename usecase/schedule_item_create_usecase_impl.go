package usecase

import (
	"errors"

	"github.com/snpiyasooriya/construction_design_api/dto"
	"github.com/snpiyasooriya/construction_design_api/interfaces/repository"
	"github.com/snpiyasooriya/construction_design_api/pkg/utils"
)

type ScheduleItemCreateUseCaseImpl struct {
	scheduleItemRepo repository.ScheduleItemRepository
	scheduleRepo     repository.ScheduleRepository
	shapeRepo        repository.ShapeRepositoryInterFace
}

func NewScheduleItemCreateUseCaseImpl(
	scheduleItemRepo repository.ScheduleItemRepository,
	scheduleRepo repository.ScheduleRepository,
	shapeRepo repository.ShapeRepositoryInterFace,
) *ScheduleItemCreateUseCaseImpl {
	return &ScheduleItemCreateUseCaseImpl{
		scheduleItemRepo: scheduleItemRepo,
		scheduleRepo:     scheduleRepo,
		shapeRepo:        shapeRepo,
	}
}

func (s *ScheduleItemCreateUseCaseImpl) Execute(scheduleItemCreateInputDTO dto.ScheduleItemCreateInputDTO) (*dto.ScheduleItemCreateOutputDTO, error) {
	// Validate that the schedule exists
	schedule, err := s.scheduleRepo.GetByID(scheduleItemCreateInputDTO.ScheduleID)
	if err != nil {
		return nil, errors.New("schedule not found")
	}
	if schedule == nil {
		return nil, errors.New("schedule not found")
	}

	// Validate that the shape exists
	shape, err := s.shapeRepo.GetByID(scheduleItemCreateInputDTO.ShapeID)
	if err != nil {
		return nil, errors.New("shape not found")
	}
	if shape == nil {
		return nil, errors.New("shape not found")
	}

	scheduleItemCreateInputDTO.ShapeDimensions = utils.InterfaceToDataTypeJson(shape.Labels)

	// Create the schedule item
	scheduleItemOutput, err := s.scheduleItemRepo.Create(scheduleItemCreateInputDTO)
	if err != nil {
		return nil, err
	}

	return scheduleItemOutput, nil
}
