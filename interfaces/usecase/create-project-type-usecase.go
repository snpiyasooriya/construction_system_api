package usecase

import (
	"github.com/snpiyasooriya/construction_design_api/dto"
)

type CreateProjectTypeUseCaseInterface interface {
	Execute(dto *dto.ProjectTypeCreateDTO) error
}
