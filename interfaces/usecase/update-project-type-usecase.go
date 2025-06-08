package usecase

import "github.com/snpiyasooriya/construction_design_api/dto"

type UpdateProjectTypeUseCaseInterface interface {
	Execute(input *dto.ProjectTypeUpdateDTO) error
}
