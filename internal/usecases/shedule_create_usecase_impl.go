package usecases

import (
	"github.com/snpiyasooriya/construction_design_api/internal/domain/interfaces"
	"github.com/snpiyasooriya/construction_design_api/internal/dto"
)

type ScheduleCreateUseCaseImpl struct {
	scheduleRepo interfaces.ScheduleRepository
}

func NewScheduleCreateUseCaseImpl(scheduleRepo interfaces.ScheduleRepository) *ScheduleCreateUseCaseImpl {
	return &ScheduleCreateUseCaseImpl{
		scheduleRepo: scheduleRepo,
	}

}

func (s *ScheduleCreateUseCaseImpl) Execute(scheduleCreateInputDTO dto.ScheduleCreateInputDTO) (*dto.ScheduleCreateOutputDTO, error) {
	createdSchedule, err := s.scheduleRepo.Create(scheduleCreateInputDTO)
	if err != nil {
		return nil, err
	}
	return createdSchedule, nil
}
