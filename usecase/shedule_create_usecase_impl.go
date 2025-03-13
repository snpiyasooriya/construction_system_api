package usecase

import (
	"github.com/snpiyasooriya/construction_design_api/dto"
	"github.com/snpiyasooriya/construction_design_api/interfaces/repository"
)

type ScheduleCreateUseCaseImpl struct {
	scheduleRepo repository.ScheduleRepository
}

func NewScheduleCreateUseCaseImpl(scheduleRepo repository.ScheduleRepository) *ScheduleCreateUseCaseImpl {
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
