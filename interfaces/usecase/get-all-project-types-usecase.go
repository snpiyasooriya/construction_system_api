package usecase

import (
	"github.com/snpiyasooriya/construction_design_api/dto"
)

type GetAllProjectTypesUseCaseInterface interface {
	Execute() ([]dto.ProjectTypeGetDto, error)
}
