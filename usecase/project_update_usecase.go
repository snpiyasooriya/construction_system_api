package usecase

import "github.com/snpiyasooriya/construction_design_api/dto"

type ProjectUpdateUseCase interface {
	Execute(project dto.ProjectUpdateDTO) error
}
