package usecase

import "github.com/snpiyasooriya/construction_design_api/dto"

type UserUpdateUseCase interface {
	Execute(id uint, input dto.UserUpdateDTO) (*dto.UserGetDTO, error)
}
