package usecase

import (
	"github.com/snpiyasooriya/construction_design_api/dto"
	"github.com/snpiyasooriya/construction_design_api/interfaces/repository"
)

type ScheduleItemCromeCreateUseCaseImpl struct {
	scheduleItemCromeRepo repository.ScheduleItemCrome
}

func NewScheduleItemCromeCreateUseCaseImpl(scheduleItemCromeRepo repository.ScheduleItemCrome) *ScheduleItemCromeCreateUseCaseImpl {
	return &ScheduleItemCromeCreateUseCaseImpl{
		scheduleItemCromeRepo: scheduleItemCromeRepo,
	}
}

func (s ScheduleItemCromeCreateUseCaseImpl) Execute(scheduleCreateInputDTO dto.ScheduleItemCromeCreateInputDTO) (*dto.ScheduleItemCromeCreateOutputDTO, error) {
	//TODO implement me
	panic("implement me")
}
