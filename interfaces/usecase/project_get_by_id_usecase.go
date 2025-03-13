package usecase

import "github.com/snpiyasooriya/construction_design_api/dto"

type ProjectGetByIDUseCase interface {
	Execute(id uint) (*dto.ProjectGetDTO, error)
}
