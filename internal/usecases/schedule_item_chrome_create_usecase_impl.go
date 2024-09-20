package usecases

import (
	"github.com/snpiyasooriya/construction_design_api/internal/domain/interfaces"
	"github.com/snpiyasooriya/construction_design_api/internal/dto"
)

type ScheduleItemCromeCreateUseCaseImpl struct {
	scheduleItemCromeRepo interfaces.ScheduleItemCrome
}

func NewScheduleItemCromeCreateUseCaseImpl(scheduleItemCromeRepo interfaces.ScheduleItemCrome) *ScheduleItemCromeCreateUseCaseImpl {
	return &ScheduleItemCromeCreateUseCaseImpl{
		scheduleItemCromeRepo: scheduleItemCromeRepo,
	}
}

func (s ScheduleItemCromeCreateUseCaseImpl) Execute(scheduleCreateInputDTO dto.ScheduleItemCromeCreateInputDTO) (*dto.ScheduleItemCromeCreateOutputDTO, error) {
	//TODO implement me
	panic("implement me")
}
