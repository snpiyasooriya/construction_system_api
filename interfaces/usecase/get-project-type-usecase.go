package usecase

import (
	"github.com/snpiyasooriya/construction_design_api/dto"
)

type GetProjectTypeUseCaseInterface interface {
	Execute(id uint) (dto.ProjectTypeGetDto, error)
}
